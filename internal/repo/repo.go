package repo

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
)

type Repo struct {
	Name        string
	Description string
	Author      string

	raw *git.Repository
}

func New(path string) (Repo, error) {
	repo, err := git.PlainOpen(path)

	if err != nil {
		return Repo{}, fmt.Errorf("failed to open repository: %w", err)
	}

	description, err := os.ReadFile(fmt.Sprintf("%s/description", path))

	if err != nil {
		return Repo{}, fmt.Errorf("failed to read repository description: %w", err)
	}

	config, err := repo.Config()

	if err != nil {
		return Repo{}, fmt.Errorf("failed to read repository config: %w", err)
	}

	return Repo{
		Name:        filepath.Base(path),
		Description: strings.TrimSpace(string(description)),
		Author:      config.Raw.Section("gitjika").Option("author"),
		raw:         repo,
	}, nil
}
