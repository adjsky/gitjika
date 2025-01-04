package repo

import (
	"fmt"
	"os"
	"path"
)

func ListAll(dirpath string) ([]Repo, error) {
	files, err := os.ReadDir(dirpath)

	if err != nil {
		return nil, fmt.Errorf("failed to read directory: %w", err)
	}

	repos := make([]Repo, 0)

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		repo, err := Open(path.Join(dirpath, file.Name()))

		if err != nil {
			continue
		}

		repos = append(repos, repo)
	}

	return repos, nil
}
