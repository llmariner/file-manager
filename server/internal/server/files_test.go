package server

import (
	"context"
	"testing"

	v1 "github.com/llm-operator/file-manager/api/v1"
	"github.com/llm-operator/file-manager/common/pkg/store"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFiles(t *testing.T) {
	st, tearDown := store.NewTest(t)
	defer tearDown()

	srv := New(st)
	ctx := context.Background()

	f, err := srv.CreateFile(ctx, &v1.CreateFileRequest{
		File:    []byte("file-content"),
		Purpose: purposeFineTune,
	})
	assert.NoError(t, err)
	assert.Equal(t, purposeFineTune, f.Purpose)
	assert.Equal(t, int64(12), f.Bytes)

	getResp, err := srv.GetFile(ctx, &v1.GetFileRequest{
		Id: f.Id,
	})
	assert.NoError(t, err)
	assert.Equal(t, f.Id, getResp.Id)

	listResp, err := srv.ListFiles(ctx, &v1.ListFilesRequest{})
	assert.NoError(t, err)
	assert.Len(t, listResp.Data, 1)
	assert.Equal(t, f.Id, listResp.Data[0].Id)

	_, err = srv.DeleteFile(ctx, &v1.DeleteFileRequest{
		Id: f.Id,
	})
	assert.NoError(t, err)

	_, err = srv.GetFile(ctx, &v1.GetFileRequest{
		Id: f.Id,
	})
	assert.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))

	listResp, err = srv.ListFiles(ctx, &v1.ListFilesRequest{})
	assert.NoError(t, err)
	assert.Len(t, listResp.Data, 0)
}
