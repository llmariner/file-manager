package server

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/llmariner/file-manager/api/v1"
	"github.com/llmariner/file-manager/server/internal/store"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestFiles(t *testing.T) {
	st, tearDown := store.NewTest(t)
	defer tearDown()

	srv := New(st, &NoopS3Client{}, "pathPrefix")
	ctx := context.Background()

	const (
		fileID = "f0"
		orgID  = "o0"
	)

	_, err := st.CreateFile(store.FileSpec{
		FileID:         fileID,
		TenantID:       defaultTenantID,
		OrganizationID: orgID,
		ProjectID:      defaultProjectID,

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

func TestCreateFile(t *testing.T) {
	st, tearDown := store.NewTest(t)
	defer tearDown()

	srv := New(st, &NoopS3Client{}, "pathPrefix")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		srv.CreateFile(w, r, nil)
	})

	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	fw, err := w.CreateFormFile("file", "test-file.jsonl")
	assert.NoError(t, err)
	_, err = fw.Write([]byte("hello"))
	assert.NoError(t, err)

	fw, err = w.CreateFormField("purpose")
	assert.NoError(t, err)
	_, err = fw.Write([]byte(purposeFineTune))
	assert.NoError(t, err)

	err = w.Close()
	assert.NoError(t, err)

	req, err := http.NewRequest("POST", "v1/files", &b)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", w.FormDataContentType())

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusCreated, rr.Code)

	var fj fileJSON
	err = json.Unmarshal(rr.Body.Bytes(), &fj)
	assert.NoError(t, err)
	assert.True(t, fj.ID != "")
	assert.Equal(t, purposeFineTune, fj.Purpose)
	assert.Equal(t, "test-file.jsonl", fj.Filename)
	assert.Equal(t, int64(5), fj.Bytes)

	resp, err := srv.GetFile(context.Background(), &v1.GetFileRequest{
		Id: fj.ID,
	})
	assert.NoError(t, err)
	assert.Equal(t, fj.ID, resp.Id)
}

func TestGetFilePath(t *testing.T) {
	st, tearDown := store.NewTest(t)
	defer tearDown()

	const (
		fileID = "f0"
		orgID  = "o0"
	)

	_, err := st.CreateFile(store.FileSpec{
		FileID:         fileID,
		TenantID:       defaultTenantID,
		OrganizationID: orgID,
		ProjectID:      defaultProjectID,

		Filename:        "filename0",
		Purpose:         "purpose0",
		ObjectStorePath: "path0",
	})
	assert.NoError(t, err)

	wsrv := NewWorkerServiceServer(st)
	got, err := wsrv.GetFilePath(context.Background(), &v1.GetFilePathRequest{
		Id: fileID,
	})
	assert.NoError(t, err)
	assert.Equal(t, "path0", got.Path)
	assert.Equal(t, "filename0", got.Filename)
}
