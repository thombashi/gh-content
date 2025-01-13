package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	cliCliRepo      = "cli/cli"
	validFilePath   = "LICENSE"
	invalidFilePath = "invalid_file_path"
)

func TestSubMain(t *testing.T) {
	a := assert.New(t)
	r := require.New(t)

	tempDir, err := os.MkdirTemp("", "testdir")
	r.NoError(err)
	defer os.RemoveAll(tempDir)

	normalTestCases := []struct {
		name           string
		repoID         string
		filePath       string
		outputFilePath string
		expectedError  string
	}{
		{
			name:           "stdout",
			repoID:         cliCliRepo,
			filePath:       validFilePath,
			outputFilePath: "",
		},
		{
			name:           "output file",
			repoID:         cliCliRepo,
			filePath:       validFilePath,
			outputFilePath: "test.txt",
		},
		{
			name:           "output file with directory",
			repoID:         cliCliRepo,
			filePath:       validFilePath,
			outputFilePath: filepath.Join("test", "test.txt"),
		},
	}
	for _, tc := range normalTestCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.outputFilePath != "" {
				defer os.Remove(tc.outputFilePath)
			}

			err := subMain(tc.repoID, tc.filePath, tc.outputFilePath)
			r.NoError(err)

			if tc.outputFilePath != "" {
				_, err := os.Stat(tc.outputFilePath)
				a.NoError(err)
			}
		})
	}

	t.Run("output file already exists", func(t *testing.T) {
		outputFilePath := filepath.Join(tempDir, "existingFilePath")

		_, err := os.Create(outputFilePath)
		r.NoError(err)
		defer os.Remove(outputFilePath)

		err = subMain(cliCliRepo, validFilePath, outputFilePath)
		a.ErrorContains(err, "file already exists")
	})

	t.Run("invalid repository ID", func(t *testing.T) {
		t.Parallel()

		invalidRepoID := "invalid/repo/id"
		err = subMain(invalidRepoID, validFilePath, "test.md")
		a.ErrorContainsf(err, "failed to fetch the content", "Could not resolve to a Repository with the name '%s'", invalidRepoID)
	})

}
