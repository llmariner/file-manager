package store

import (
	"gorm.io/gorm"
)

// File represents a file.
type File struct {
	gorm.Model

	FileID string `gorm:"uniqueIndex"`

	TenantID       string `gorm:"index"`
	OrganizationID string
	ProjectID      string `gorm:"index"`

	Filename string
	Purpose  string

	Bytes int64

	ObjectStorePath string
}

// FileSpec is a spec of the file
type FileSpec struct {
	FileID         string
	TenantID       string
	OrganizationID string
	ProjectID      string

	Filename string
	Purpose  string
	Bytes    int64

	ObjectStorePath string
}

// CreateFile creates a file.
func (s *S) CreateFile(spec FileSpec) (*File, error) {
	f := &File{
		FileID:         spec.FileID,
		TenantID:       spec.TenantID,
		OrganizationID: spec.OrganizationID,
		ProjectID:      spec.ProjectID,

		Filename: spec.Filename,
		Purpose:  spec.Purpose,
		Bytes:    spec.Bytes,

		ObjectStorePath: spec.ObjectStorePath,
	}
	if err := s.db.Create(f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

// GetFile returns a file by file ID and projectID
func (s *S) GetFile(fileID, projectID string) (*File, error) {
	var f File
	if err := s.db.Where("file_id = ? AND project_id = ?", fileID, projectID).Take(&f).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

// GetFileByFileIDAndTenantID returns a file by file ID and a tenant ID.
func (s *S) GetFileByFileIDAndTenantID(fileID, tenantID string) (*File, error) {
	var f File
	if err := s.db.Where("file_id = ? AND tenant_id = ?", fileID, tenantID).Take(&f).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

// GetFileByFileID returns a file by file ID.
func (s *S) GetFileByFileID(fileID string) (*File, error) {
	var f File
	if err := s.db.Where("file_id = ?", fileID).Take(&f).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

// GetFileByFileIDAndProjectID returns a file by file ID and project ID.
func (s *S) GetFileByFileIDAndProjectID(fileID, projectID string) (*File, error) {
	var f File
	if err := s.db.Where("file_id = ? AND project_id = ?", fileID, projectID).Take(&f).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

// ListFilesByProjectID lists files.
func (s *S) ListFilesByProjectID(projectID string) ([]*File, error) {
	var fs []*File
	if err := s.db.Where("project_id = ?", projectID).Order("id DESC").Find(&fs).Error; err != nil {
		return nil, err
	}
	return fs, nil
}

// ListFilesByProjectIDAndPurpose list files.
func (s *S) ListFilesByProjectIDAndPurpose(projectID, purpose string) ([]*File, error) {
	var fs []*File
	if err := s.db.Where("project_id = ? AND purpose = ?", projectID).Order("id DESC").Find(&fs).Error; err != nil {
		return nil, err
	}
	return fs, nil
}

// ListFilesByProjectIDWithPagination lists files with pagination.
func (s *S) ListFilesByProjectIDWithPagination(projectID string, afterID uint, limit int, order string) ([]*File, bool, error) {
	var fs []*File
	query := s.db.Where("project_id = ?", projectID)
	if afterID > 0 {
		if order == "asc" {
			query = query.Where("id > ?", afterID)
		} else {
			query = query.Where("id < ?", afterID)
		}
	}

	if order == "asc" {
		query = query.Order("id ASC")
	} else {
		query = query.Order("id DESC")
	}
	query = query.Limit(limit + 1)
	if err := query.Find(&fs).Error; err != nil {
		return nil, false, err
	}

	hasMore := len(fs) > limit
	if hasMore {
		fs = fs[:limit]
	}
	return fs, hasMore, nil
}

// ListFilesByProjectIDAndPurposeWithPagination lists files with pagination.
func (s *S) ListFilesByProjectIDAndPurposeWithPagination(projectID, purpose string, afterID uint, limit int, order string) ([]*File, bool, error) {
	var fs []*File
	query := s.db.Where("project_id = ? AND purpose = ?", projectID, purpose)
	if afterID > 0 {
		if order == "asc" {
			query = query.Where("id > ?", afterID)
		} else {
			query = query.Where("id < ?", afterID)
		}
	}

	if order == "asc" {
		query = query.Order("id ASC")
	} else {
		query = query.Order("id DESC")
	}
	query = query.Limit(limit + 1)
	if err := query.Find(&fs).Error; err != nil {
		return nil, false, err
	}

	hasMore := len(fs) > limit
	if hasMore {
		fs = fs[:limit]
	}
	return fs, hasMore, nil
}

// CountFilesByProjectID counts files by project ID.
func (s *S) CountFilesByProjectID(projectID string) (int64, error) {
	var count int64
	if err := s.db.Model(&File{}).Where("project_id = ?", projectID).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// DeleteFile deletes a file by file ID and project ID.
func (s *S) DeleteFile(fileID, projectID string) error {
	res := s.db.Unscoped().Where("file_id = ? AND project_id = ?", fileID, projectID).Delete(&File{})
	if err := res.Error; err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
