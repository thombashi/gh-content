package main

import (
	"fmt"

	"github.com/cli/go-gh/v2/pkg/repository"
	"github.com/spf13/pflag"
)

type Flags struct {
	RepoID         string
	FilePath       string
	OutputFilePath string

	LogLevelStr string
}

func toRepoID(repo repository.Repository) string {
	return fmt.Sprintf("%s/%s", repo.Owner, repo.Name)
}

func setFlags() (*Flags, []string, error) {
	var flags Flags

	pflag.StringVarP(
		&flags.RepoID,
		"repo",
		"R",
		"",
		"GitHub repository ID. If not specified, use the current repository.",
	)
	pflag.StringVarP(
		&flags.OutputFilePath,
		"output",
		"o",
		"",
		"output file path. If not specified, output to stdout.",
	)

	pflag.StringVar(
		&flags.LogLevelStr,
		"log-level",
		"info",
		"log level (debug, info, warn, error)",
	)

	pflag.Parse()

	if flags.RepoID == "" {
		repo, err := repository.Current()
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get the current repository: %w", err)
		}

		flags.RepoID = toRepoID(repo)
	}

	args := pflag.Args()
	if len(args) == 0 {
		return nil, nil, fmt.Errorf("require a file path in the repository")
	}

	return &flags, args, nil
}
