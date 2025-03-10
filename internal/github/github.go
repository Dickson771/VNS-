package github

import (
	"fmt"
	"io"
	"net/http"
)

// FetchFile fetches the content of a file from a GitHub repository.
func FetchFile(repo, file string) ([]byte, error) {
	// Correct the URL to point to the raw content on GitHub
	url := fmt.Sprintf("https://raw.githubusercontent.com/%s/main/%s", repo, file)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch file: %s", resp.Status)
	}

	// Read the content of the file and return it
	content, err := io.ReadAll(resp.Body) // Using io.ReadAll for Go 1.16+
	if err != nil {
		return nil, err
	}

	return content, nil
}
