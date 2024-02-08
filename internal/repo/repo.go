package repo

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/config"
)

var (
	ErrRepositoryNotFound           = errors.New("repository not found")
	ErrFailedToReadRepositoryConfig = errors.New("failed to read repository config")
	ErrFailedToReadRefs             = errors.New("failed to read references")
)

type Repo struct {
	path   string
	raw    *git.Repository
	config *config.Config
}

func New(path string) (Repo, error) {
	repo, err := git.PlainOpen(path)

	if err != nil {
		return Repo{}, ErrRepositoryNotFound
	}

	config, err := repo.Config()

	if err != nil {
		return Repo{}, ErrFailedToReadRepositoryConfig
	}

	return Repo{
		path:   path,
		config: config,
		raw:    repo,
	}, nil
}

func (repo Repo) Name() string {
	return filepath.Base(repo.path)
}

func (repo Repo) Description() string {
	description, err := os.ReadFile(fmt.Sprintf("%s/description", repo.path))

	if err != nil {
		return ""
	}

	return strings.TrimSpace(string(description))
}

func (repo Repo) Author() string {
	return repo.config.Raw.Section("gitjika").Option("author")
}

func (repo Repo) Age() time.Time {
	agefile, err := os.ReadFile(fmt.Sprintf("%s/last-modified", repo.path))

	if err != nil {
		return time.Time{}
	}

	age, _ := time.Parse("2006-01-02 15:04:05 -0700", strings.TrimSpace(string(agefile)))

	return age.UTC()
}
