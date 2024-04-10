package store

import (
	"gorm.io/gorm"
)

// File represents a file.
type File struct {
	gorm.Model

	FileID   string `gorm:"uniqueIndex:idx_file_file_id_tenant_id"`
	TenantID string `gorm:"uniqueIndex:idx_file_file_id_tenant_id"`

	Filename string
	Purpose  string `gorm:"index:idx_file_tenant_id_purpose"`

	Bytes int64

	// TODO(kenji): Add a file location. The actual file content is stored outside of the SQL database.
}

// FileKey represents a file key.
type FileKey struct {
	FileID   string
	TenantID string
}

// FileSpec is a spec of the file
type FileSpec struct {
	Key FileKey

	Filename string
	Purpose  string
	Bytes    int64
}

// CreateFile creates a file.
func (s *S) CreateFile(spec FileSpec) (*File, error) {
	f := &File{
		FileID:   spec.Key.FileID,
		TenantID: spec.Key.TenantID,

		Filename: spec.Filename,
		Purpose:  spec.Purpose,
		Bytes:    spec.Bytes,
	}
	if err := s.db.Create(f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

// GetFile returns a file by file ID and tenant ID.
func (s *S) GetFile(k FileKey) (*File, error) {
	var f File
	if err := s.db.Where("file_id = ? AND tenant_id = ?", k.FileID, k.TenantID).Take(&f).Error; err != nil {
		return nil, err
	}
	return &f, nil
}

// ListFilesByTenantID list files.
func (s *S) ListFilesByTenantID(tenantID string) ([]*File, error) {
	var fs []*File
	if err := s.db.Where("tenant_id = ?", tenantID).Find(&fs).Error; err != nil {
		return nil, err
	}
	return fs, nil
}

// ListFilesByTenantIDAndPurpose list files.
func (s *S) ListFilesByTenantIDAndPurpose(tenantID, purpose string) ([]*File, error) {
	var fs []*File
	if err := s.db.Where("tenant_id = ? AND purpose = ?", tenantID).Find(&fs).Error; err != nil {
		return nil, err
	}
	return fs, nil
}

// DeleteFile deletes a file by file ID and tenant ID.
func (s *S) DeleteFile(k FileKey) error {
	res := s.db.Unscoped().Where("file_id = ? AND tenant_id = ?", k.FileID, k.TenantID).Delete(&File{})
	if err := res.Error; err != nil {
		return err
	}
	if res.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}
