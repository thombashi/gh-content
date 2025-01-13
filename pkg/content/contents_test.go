package content

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchContent(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	testCases := []struct {
		repo     string
		filePath string
		wantErr  bool
	}{
		{
			repo:     "cli/cli",
			filePath: "LICENSE",
			wantErr:  false,
		},
		{
			repo:     "cli/cli",
			filePath: "non_existent_file.md",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.repo+"/"+tc.filePath, func(t *testing.T) {
			t.Parallel()

			got, err := FetchContent(tc.repo, tc.filePath)
			if tc.wantErr {
				r.Error(err)
			} else {
				r.NoError(err)
				a.NotEmpty(got.Bytes())
			}
		})
	}
}

func TestFetchLastModified(t *testing.T) {
	const (
		repo            = "cli/cli"
		validFilePath   = "LICENSE"
		invalidFilePath = "invalid_file_path"
	)

	a := assert.New(t)
	r := require.New(t)

	t.Run("normal case", func(t *testing.T) {
		t.Parallel()

		lastModified, err := FetchLastModified(repo, validFilePath)
		r.NoError(err)
		a.NotNil(lastModified)
	})

	t.Run("no commits", func(t *testing.T) {
		t.Parallel()

		_, err := FetchLastModified(repo, "path/to/nonexistent/file")
		r.Error(err)
		a.ErrorContains(err, "no commits found")
	})

	t.Run("invalid repository", func(t *testing.T) {
		t.Parallel()

		invalidRepoID := "invalid/repo/id"

		_, err := FetchLastModified(invalidRepoID, validFilePath)
		r.Error(err)
		a.ErrorContainsf(err, "repository not found", "Could not resolve to a Repository with the name '%s'", invalidRepoID)
	})
}
