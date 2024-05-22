package store

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestFile(t *testing.T) {
	st, tearDown := NewTest(t)
	defer tearDown()

	const (
		fileID         = "f0"
		tenantID       = "tid0"
		organizationID = "oid0"
		projectID      = "pid0"
	)

	_, err := st.GetFile(fileID, projectID)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))

	_, err = st.CreateFile(FileSpec{
		FileID:         fileID,
		TenantID:       tenantID,
		OrganizationID: organizationID,
		ProjectID:      projectID,

		Filename: "filename0",
		Purpose:  "purpose0",
	})
	assert.NoError(t, err)

	gotM, err := st.GetFile(fileID, projectID)
	assert.NoError(t, err)
	assert.Equal(t, fileID, gotM.FileID)
	assert.Equal(t, tenantID, gotM.TenantID)

	gotMs, err := st.ListFilesByProjectID(projectID)
	assert.NoError(t, err)
	assert.Len(t, gotMs, 1)

	_, err = st.CreateFile(FileSpec{
		FileID:         "f1",
		TenantID:       "tid1",
		OrganizationID: "oid1",
		ProjectID:      "pid1",

		Filename: "filename1",
		Purpose:  "purpose1",
	})
	assert.NoError(t, err)

	gotMs, err = st.ListFilesByProjectID(projectID)
	assert.NoError(t, err)
	assert.Len(t, gotMs, 1)

	err = st.DeleteFile(fileID, projectID)
	assert.NoError(t, err)

	gotMs, err = st.ListFilesByProjectID(projectID)
	assert.NoError(t, err)
	assert.Len(t, gotMs, 0)

	err = st.DeleteFile(fileID, projectID)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
