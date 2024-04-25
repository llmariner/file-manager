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
		fileID   = "m0"
		tenantID = "tid0"
	)

	k := FileKey{
		FileID:   fileID,
		TenantID: tenantID,
	}
	_, err := st.GetFile(k)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))

	_, err = st.CreateFile(FileSpec{
		Key:      k,
		Filename: "filename0",
		Purpose:  "purpose0",
	})
	assert.NoError(t, err)

	gotM, err := st.GetFile(k)
	assert.NoError(t, err)
	assert.Equal(t, fileID, gotM.FileID)
	assert.Equal(t, tenantID, gotM.TenantID)

	gotMs, err := st.ListFilesByTenantID(tenantID)
	assert.NoError(t, err)
	assert.Len(t, gotMs, 1)

	k1 := FileKey{
		FileID:   "m1",
		TenantID: "tid1",
	}
	_, err = st.CreateFile(FileSpec{
		Key:      k1,
		Filename: "filename1",
		Purpose:  "purpose1",
	})
	assert.NoError(t, err)

	gotMs, err = st.ListFilesByTenantID(tenantID)
	assert.NoError(t, err)
	assert.Len(t, gotMs, 1)

	err = st.DeleteFile(k)
	assert.NoError(t, err)

	gotMs, err = st.ListFilesByTenantID(tenantID)
	assert.NoError(t, err)
	assert.Len(t, gotMs, 0)

	err = st.DeleteFile(k)
	assert.True(t, errors.Is(err, gorm.ErrRecordNotFound))
}
