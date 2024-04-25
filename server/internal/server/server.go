package server

import (
	"fmt"
	"io"
	"log"
	"net"

	v1 "github.com/llm-operator/file-manager/api/v1"
	"github.com/llm-operator/file-manager/server/internal/store"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

// New creates a server.
func New(store *store.S, s3Client S3Client, pathPrefix string) *S {
	return &S{
		store:      store,
		s3Client:   s3Client,
		pathPrefix: pathPrefix,
	}
}

// S is a server.
type S struct {
	v1.UnimplementedFilesServiceServer

	srv *grpc.Server

	store    *store.S
	s3Client S3Client

	pathPrefix string
}

// Run starts the gRPC server.
func (s *S) Run(port int) error {
	log.Printf("Starting server on port %d\n", port)

	grpcServer := grpc.NewServer()
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
