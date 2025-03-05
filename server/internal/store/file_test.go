package store

import (
	"errors"
	"fmt"
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

	gotM, err = st.GetFileByFileIDAndTenantID(fileID, tenantID)
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

func TestFilePagination(t *testing.T) {
	const (
		projectID = "pid0"
		purpose   = "purpose0"
	)

	tcs := []struct {
		name           string
		purpose        string
		cursor         uint
		limit          int
		order          string
		wantFiles      []string
		wantHasMore    bool
		wantTotalCount int64
	}{
		{
			name:           "descending order first page",
			cursor:         0,
			limit:          3,
			order:          "desc",
			wantFiles:      []string{"f4", "f3", "f2"},
			wantHasMore:    true,
			wantTotalCount: 5,
		},
		{
			name:           "ascending order first page",
			cursor:         0,
			limit:          3,
			order:          "asc",
			wantFiles:      []string{"f0", "f1", "f2"},
			wantHasMore:    true,
			wantTotalCount: 5,
		},
		{
			name:           "descending order with cursor",
			cursor:         2,
			limit:          3,
			order:          "desc",
			wantFiles:      []string{"f0"},
			wantHasMore:    false,
			wantTotalCount: 5,
		},
		{
			name:           "ascending order with cursor",
			cursor:         2,
			limit:          3,
			order:          "asc",
			wantFiles:      []string{"f2", "f3", "f4"},
			wantHasMore:    false,
			wantTotalCount: 5,
		},
		{
			name:           "with purpose filter",
			purpose:        purpose,
			cursor:         0,
			limit:          3,
			order:          "desc",
			wantFiles:      []string{"f4", "f3", "f2"},
			wantHasMore:    true,
			wantTotalCount: 5,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			st, tearDown := NewTest(t)
			defer tearDown()

			for i := 0; i < 5; i++ {
				_, err := st.CreateFile(FileSpec{
					FileID:         fmt.Sprintf("f%d", i),
					TenantID:       fmt.Sprintf("tid%d", i),
					OrganizationID: fmt.Sprintf("oid%d", i),
					ProjectID:      projectID,
					Filename:       fmt.Sprintf("filename%d", i),
					Purpose:        purpose,
				})
				assert.NoError(t, err)
			}

			var (
				files   []*File
				hasMore bool
				err     error
			)

			if tc.purpose != "" {
				files, hasMore, err = st.ListFilesByProjectIDAndPurposeWithPagination(projectID, tc.purpose, tc.cursor, tc.limit, tc.order)
			} else {
				files, hasMore, err = st.ListFilesByProjectIDWithPagination(projectID, tc.cursor, tc.limit, tc.order)
			}

			assert.NoError(t, err)
			assert.Equal(t, tc.wantHasMore, hasMore)
			assert.Len(t, files, len(tc.wantFiles))

			for i, wantFileID := range tc.wantFiles {
				assert.Equal(t, wantFileID, files[i].FileID)
			}

			count, err := st.CountFilesByProjectID(projectID)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantTotalCount, count)
		})
	}
}
