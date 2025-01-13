package main

import (
	"fmt"
	"log/slog"
	"os"
	"path"
	"path/filepath"

	"github.com/phsym/console-slog"
	"github.com/thombashi/eoe"
	"github.com/thombashi/gh-content/pkg/content"
)

func newLogger(level slog.Level) *slog.Logger {
	logger := slog.New(
		console.NewHandler(os.Stderr, &console.HandlerOptions{
			Level: level,
		}),
	)

	return logger
}

func subMain(repoID, filePath, outputFilePath string) error {
	var err error

	body, err := content.FetchContent(repoID, filePath)
	if err != nil {
		return fmt.Errorf("failed to fetch the content: %w", err)
	}

	if outputFilePath == "" {
		fmt.Println(body.String())
		return nil
	}

	fi, err := os.Stat(outputFilePath)
	if err == nil && !fi.IsDir() {
		return fmt.Errorf("the output file already exists: path=%s, error=%w", outputFilePath, err)
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("failed to get the file info: %w", err)
	}

	if err == nil && fi.IsDir() {
		err = os.MkdirAll(outputFilePath, 0755)
		if err != nil {
			return fmt.Errorf("failed to create the directory: %w", err)
		}

		fileName := path.Base(filePath)
		outputFilePath = path.Join(outputFilePath, fileName)
	} else {
		// if the outputFilePath is a file, create the parent directory
		dir := filepath.Dir(outputFilePath)
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			return fmt.Errorf("failed to create the directory: %w", err)
		}
	}

	return os.WriteFile(outputFilePath, body.Bytes(), 0744)
}

func main() {
	flags, args, err := setFlags()
	eoe.ExitOnError(err, eoe.NewParams().WithMessage("failed to set flags"))

	var logLevel slog.Level
	err = logLevel.UnmarshalText([]byte(flags.LogLevelStr))
	eoe.ExitOnError(err, eoe.NewParams().WithMessage("failed to get a slog level"))

	logger := newLogger(logLevel)
	eoeParams := eoe.NewParams().WithLogger(logger)

	err = subMain(flags.RepoID, args[0], flags.OutputFilePath)
	eoe.ExitOnError(err, eoeParams.WithMessage("failed to execute subMain"))
}
