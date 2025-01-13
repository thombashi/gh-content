package contents

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
	"path"
	"strings"
	"time"

	gh "github.com/cli/go-gh/v2"
)

var repoExistenceCache map[string]bool

func init() {
	repoExistenceCache = map[string]bool{}
}

// IsExistRepo checks if the specified repository exists.
func IsExistRepo(repo string) (bool, error) {
	if _, ok := repoExistenceCache[repo]; ok {
		return true, nil
	}

	_, stderr, err := gh.Exec("repo", "view", repo)
	if err != nil {
		return false, fmt.Errorf("repository not found: %s", stderr.String())
	}

	repoExistenceCache[repo] = true

	return true, nil
}

// FetchContent fetches the content of the specified file in the repository.
func FetchContent(repo, filePath string) (bytes.Buffer, error) {
	if exist, err := IsExistRepo(repo); !exist {
		return bytes.Buffer{}, err
	}

	filePath = path.Clean(strings.TrimSpace(filePath))
	endpoint := fmt.Sprintf("/repos/%s/contents/%s", repo, filePath)
	stdout, _, err := gh.Exec("api", "-H", "Accept: application/vnd.github.v3.raw", endpoint)
	if err != nil {
		return stdout, fmt.Errorf("failed to execute gh: %w", err)
	}

	return stdout, nil
}

// FetchLastModified fetches the last modified time of the specified file in the repository.
func FetchLastModified(repo, filePath string) (*time.Time, error) {
	type Commit struct {
		Commit struct {
			Committer struct {
				Date string `json:"date"`
			} `json:"committer"`
		} `json:"commit"`
	}

	if exist, err := IsExistRepo(repo); !exist {
		return nil, err
	}

	filePath = path.Clean(strings.TrimSpace(filePath))
	filePath = url.QueryEscape(filePath)

	stdout, _, err := gh.Exec(
		"api", "-H", "Accept: application/vnd.github.v3+json",
		fmt.Sprintf("/repos/%s/commits?path=%s", repo, filePath),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to execute gh: %w", err)
	}

	var commits []Commit
	if err := json.Unmarshal(stdout.Bytes(), &commits); err != nil {
		return nil, fmt.Errorf("failed to decode the response: %w", err)
	}

	if len(commits) == 0 {
		return nil, fmt.Errorf("no commits found for the %s file in the %s repository", filePath, repo)
	}

	t, err := time.Parse(time.RFC3339, commits[0].Commit.Committer.Date)
	if err != nil {
		return nil, fmt.Errorf("failed to parse the last modified time: %w", err)
	}

	return &t, nil
}
