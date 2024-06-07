package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"

	v1 "github.com/llm-operator/file-manager/api/v1"
	"github.com/llm-operator/file-manager/server/internal/config"
	"github.com/llm-operator/file-manager/server/internal/store"
	"github.com/llm-operator/rbac-manager/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

const (
	defaultProjectID = "default"
	defaultTenantID  = "default-tenant-id"
)

// S3Client is an interface for an S3 client.
type S3Client interface {
	Upload(r io.Reader, key string) error
}

// NoopS3Client is a no-op S3 client.
type NoopS3Client struct{}

// Upload is a no-op implementation of Upload.
func (n *NoopS3Client) Upload(r io.Reader, key string) error {
	return nil
}

type reqIntercepter interface {
	InterceptHTTPRequest(req *http.Request) (int, auth.UserInfo, error)
}

type noopReqIntercepter struct {
}

func (n noopReqIntercepter) InterceptHTTPRequest(req *http.Request) (int, auth.UserInfo, error) {
	ui := auth.UserInfo{
		OrganizationID:      "default",
		ProjectID:           defaultProjectID,
		KubernetesNamespace: "default",
		TenantID:            defaultTenantID,
	}
	return http.StatusOK, ui, nil
}

// New creates a server.
func New(store *store.S, s3Client S3Client, pathPrefix string) *S {
	return &S{
		store:          store,
		s3Client:       s3Client,
		pathPrefix:     pathPrefix,
		reqIntercepter: noopReqIntercepter{},
	}
}

// S is a server.
type S struct {
	v1.UnimplementedFilesServiceServer

	srv *grpc.Server

	store    *store.S
	s3Client S3Client

	pathPrefix string

	reqIntercepter reqIntercepter
	enableAuth     bool
}

// Run starts the gRPC server.
func (s *S) Run(ctx context.Context, port int, authConfig config.AuthConfig) error {
	log.Printf("Starting server on port %d\n", port)

	var opts []grpc.ServerOption
	if authConfig.Enable {
		ai, err := auth.NewInterceptor(ctx, auth.Config{
			RBACServerAddr: authConfig.RBACInternalServerAddr,
			AccessResource: "api.files",
		})
		if err != nil {
			return err
		}
		opts = append(opts, grpc.ChainUnaryInterceptor(ai.Unary()))

		s.reqIntercepter = ai
		s.enableAuth = true
	}

	grpcServer := grpc.NewServer(opts...)
	v1.RegisterFilesServiceServer(grpcServer, s)
	reflection.Register(grpcServer)

	s.srv = grpcServer

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("listen: %s", err)
	}
	if err := grpcServer.Serve(l); err != nil {
		return fmt.Errorf("serve: %s", err)
	}
	return nil
}

// Stop stops the gRPC server.
func (s *S) Stop() {
	s.srv.Stop()
}

func (s *S) extractUserInfoFromContext(ctx context.Context) (*auth.UserInfo, error) {
	if !s.enableAuth {
		return &auth.UserInfo{
			OrganizationID:      "default",
			ProjectID:           defaultProjectID,
			KubernetesNamespace: "default",
			TenantID:            defaultTenantID,
		}, nil
	}
	userInfo, ok := auth.ExtractUserInfoFromContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "user info not found")
	}
	return userInfo, nil
}
