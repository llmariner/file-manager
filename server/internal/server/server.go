package server

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/go-logr/logr"
	"github.com/llmariner/api-usage/pkg/sender"
	v1 "github.com/llmariner/file-manager/api/v1"
	"github.com/llmariner/file-manager/server/internal/config"
	"github.com/llmariner/file-manager/server/internal/store"
	"github.com/llmariner/rbac-manager/pkg/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const (
	defaultProjectID = "default"
	defaultTenantID  = "default-tenant-id"
)

// S3Client is an interface for an S3 client.
type S3Client interface {
	Upload(ctx context.Context, r io.Reader, key string) error
}

// NoopS3Client is a no-op S3 client.
type NoopS3Client struct{}

// Upload is a no-op implementation of Upload.
func (n *NoopS3Client) Upload(ctx context.Context, r io.Reader, key string) error {
	return nil
}

type reqIntercepter interface {
	InterceptHTTPRequest(req *http.Request) (int, auth.UserInfo, error)
}

type noopReqIntercepter struct {
}

func (n noopReqIntercepter) InterceptHTTPRequest(req *http.Request) (int, auth.UserInfo, error) {
	ui := auth.UserInfo{
		OrganizationID: "default",
		ProjectID:      defaultProjectID,
		AssignedKubernetesEnvs: []auth.AssignedKubernetesEnv{
			{
				ClusterID: defaultClusterID,
				Namespace: "default",
			},
		},
		TenantID: defaultTenantID,
	}
	return http.StatusOK, ui, nil
}

// New creates a server.
func New(store *store.S, s3Client S3Client, sender sender.UsageSetter, pathPrefix string, log logr.Logger) *S {
	return &S{
		store:          store,
		s3Client:       s3Client,
		usage:          sender,
		log:            log.WithName("grpc"),
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
	usage    sender.UsageSetter
	log      logr.Logger

	pathPrefix string

	reqIntercepter reqIntercepter
}

// Run starts the gRPC server.
func (s *S) Run(ctx context.Context, port int, authConfig config.AuthConfig) error {
	s.log.Info("Starting gRPC server...", "port", port)

	var opt grpc.ServerOption
	if authConfig.Enable {
		ai, err := auth.NewInterceptor(ctx, auth.Config{
			RBACServerAddr: authConfig.RBACInternalServerAddr,
			AccessResource: "api.files",
		})
		if err != nil {
			return err
		}
		opt = grpc.ChainUnaryInterceptor(ai.Unary("/grpc.health.v1.Health/Check"), sender.Unary(s.usage))
		s.reqIntercepter = ai
	} else {
		fakeAuth := func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			return handler(fakeAuthInto(ctx), req)
		}
		opt = grpc.ChainUnaryInterceptor(fakeAuth, sender.Unary(s.usage))
	}

	grpcServer := grpc.NewServer(opt)
	v1.RegisterFilesServiceServer(grpcServer, s)
	reflection.Register(grpcServer)

	healthCheck := health.NewServer()
	healthCheck.SetServingStatus("", grpc_health_v1.HealthCheckResponse_SERVING)
	grpc_health_v1.RegisterHealthServer(grpcServer, healthCheck)

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

// fakeAuthInto sets dummy user info and token into the context.
func fakeAuthInto(ctx context.Context) context.Context {
	return auth.AppendUserInfoToContext(ctx, auth.UserInfo{
		OrganizationID: "default",
		ProjectID:      defaultProjectID,
		AssignedKubernetesEnvs: []auth.AssignedKubernetesEnv{
			{
				ClusterID: defaultClusterID,
				Namespace: "default",
			},
		},
		TenantID: defaultTenantID,
	})
}
