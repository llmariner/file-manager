package server

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/llmariner/common/pkg/id"
	v1 "github.com/llmariner/file-manager/api/v1"
	"github.com/llmariner/file-manager/server/internal/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

const (
	purposeFineTune   = "fine-tune"
	purposeAssistants = "assistants"
)

// CreateFile creates a file.
func (s *S) CreateFile(
	w http.ResponseWriter,
	req *http.Request,
	pathParams map[string]string,
) {
	status, userInfo, err := s.reqIntercepter.InterceptHTTPRequest(req)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	if err := req.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	purpose := req.FormValue("purpose")
	if purpose == "" {
		http.Error(w, "purpose is required", http.StatusBadRequest)
		return
	}
	if err := validatePurpose(purpose); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Read the file content.
	file, header, err := req.FormFile("file")
	if err != nil {
		if err == http.ErrMissingFile {
			http.Error(w, "file is required", http.StatusBadRequest)
			return
		}
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer func() {
		_ = file.Close()
	}()

	s.log.Info("Uploading the file to S3")
	fileID, err := id.GenerateID("file-", 24)
	if err != nil {
		http.Error(w, fmt.Sprintf("generate file id: %s", err.Error()), http.StatusInternalServerError)
		return
	}

	path := s.filePath(fileID)
	if err := s.s3Client.Upload(req.Context(), file, path); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.log.Info("Uploaded the file", "header(bytes)", header.Size)

	f, err := s.store.CreateFile(store.FileSpec{
		FileID:         fileID,
		TenantID:       userInfo.TenantID,
		OrganizationID: userInfo.OrganizationID,
		ProjectID:      userInfo.ProjectID,

		Purpose:  purpose,
		Filename: header.Filename,
		Bytes:    header.Size,

		ObjectStorePath: path,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fj := toFileJSON(f)
	b, err := json.Marshal(fj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// GetFileContent gets a file content.
func (s *S) GetFileContent(
	w http.ResponseWriter,
	req *http.Request,
	pathParams map[string]string,
) {
	http.Error(w, "Not implemented", http.StatusInternalServerError)
}

// ListFiles lists files.
func (s *S) ListFiles(
	ctx context.Context,
	req *v1.ListFilesRequest,
) (*v1.ListFilesResponse, error) {
	userInfo, err := s.extractUserInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var fs []*store.File
	if p := req.Purpose; p != "" {
		if err := validatePurpose(p); err != nil {
			return nil, err
		}
		fs, err = s.store.ListFilesByProjectIDAndPurpose(userInfo.ProjectID, p)
	} else {
		fs, err = s.store.ListFilesByProjectID(userInfo.ProjectID)
	}
	if err != nil {
		return nil, status.Errorf(codes.Internal, "list files: %s", err)
	}

	var fileProtos []*v1.File
	for _, f := range fs {
		fileProtos = append(fileProtos, toFileProto(f))
	}
	return &v1.ListFilesResponse{
		Object: "list",
		Data:   fileProtos,
	}, nil
}

// GetFile gets a file.
func (s *S) GetFile(
	ctx context.Context,
	req *v1.GetFileRequest,
) (*v1.File, error) {
	userInfo, err := s.extractUserInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	f, err := s.store.GetFile(req.Id, userInfo.ProjectID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "file %q not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "get file: %s", err)
	}
	return toFileProto(f), nil
}

// DeleteFile deletes a file.
func (s *S) DeleteFile(
	ctx context.Context,
	req *v1.DeleteFileRequest,
) (*v1.DeleteFileResponse, error) {
	userInfo, err := s.extractUserInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	if err := s.store.DeleteFile(req.Id, userInfo.ProjectID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "file %q not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "delete file: %s", err)
	}
	return &v1.DeleteFileResponse{
		Id:      req.Id,
		Object:  "file",
		Deleted: true,
	}, nil
}

// GetFilePath gets a file path.
func (s *WS) GetFilePath(
	ctx context.Context,
	req *v1.GetFilePathRequest,
) (*v1.GetFilePathResponse, error) {
	clusterInfo, err := s.extractClusterInfoFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	f, err := s.store.GetFileByFileIDAndTenantID(req.Id, clusterInfo.TenantID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "file %q not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "get file: %s", err)
	}
	return &v1.GetFilePathResponse{
		Path:     f.ObjectStorePath,
		Filename: f.Filename,
	}, nil
}

// GetFilePath gets a file path.
func (s *IS) GetFilePath(
	ctx context.Context,
	req *v1.GetFilePathRequest,
) (*v1.GetFilePathResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	f, err := s.store.GetFileByFileID(req.Id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(codes.NotFound, "file %q not found", req.Id)
		}
		return nil, status.Errorf(codes.Internal, "get file: %s", err)
	}
	return &v1.GetFilePathResponse{
		Path:     f.ObjectStorePath,
		Filename: f.Filename,
	}, nil
}

func (s *S) filePath(key string) string {
	return fmt.Sprintf("%s/%s", s.pathPrefix, key)
}

func validatePurpose(p string) error {
	switch p {
	case purposeFineTune, purposeAssistants:
		return nil
	default:
		return status.Errorf(codes.InvalidArgument, "invalid purpose")
	}
}

func toFileProto(f *store.File) *v1.File {
	return &v1.File{
		Id:        f.FileID,
		Bytes:     f.Bytes,
		CreatedAt: f.CreatedAt.UTC().Unix(),
		Filename:  f.Filename,
		Object:    "file",
		Purpose:   f.Purpose,
	}
}

type fileJSON struct {
	ID        string `json:"id"`
	Bytes     int64  `json:"bytes"`
	CreatedAt int64  `json:"created_at"`
	Filename  string `json:"filename"`
	Object    string `json:"object"`
	Purpose   string `json:"purpose"`
}

func toFileJSON(f *store.File) *fileJSON {
	return &fileJSON{
		ID:        f.FileID,
		Bytes:     f.Bytes,
		CreatedAt: f.CreatedAt.UTC().Unix(),
		Filename:  f.Filename,
		Object:    "file",
		Purpose:   f.Purpose,
	}
}
