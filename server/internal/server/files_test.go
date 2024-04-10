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

	const fileID = "f0"

	_, err := st.CreateFile(store.FileSpec{
		Key: store.FileKey{
			FileID:   fileID,
			TenantID: fakeTenantID,
		},
		Filename: "filename0",
		Purpose:  "purpose0",
	})
	assert.NoError(t, err)

	getResp, err := srv.GetFile(ctx, &v1.GetFileRequest{
		Id: fileID,
	})
	assert.NoError(t, err)
	assert.Equal(t, fileID, getResp.Id)

	listResp, err := srv.ListFiles(ctx, &v1.ListFilesRequest{})
	assert.NoError(t, err)
	assert.Len(t, listResp.Data, 1)
	assert.Equal(t, fileID, listResp.Data[0].Id)

	_, err = srv.DeleteFile(ctx, &v1.DeleteFileRequest{
		Id: fileID,
	})
	assert.NoError(t, err)

	_, err = srv.GetFile(ctx, &v1.GetFileRequest{
		Id: fileID,
	})
	assert.Error(t, err)
	assert.Equal(t, codes.NotFound, status.Code(err))

	listResp, err = srv.ListFiles(ctx, &v1.ListFilesRequest{})
	assert.NoError(t, err)
	assert.Len(t, listResp.Data, 0)
}
