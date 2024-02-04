package repo_test

import (
	"path"
	"testing"

	"github.com/adjsky/gitjika/internal/repo"
	"gotest.tools/v3/assert"
)

func TestNew(t *testing.T) {
	basicBareRepo, err := repo.New(path.Join("fixtures", "repos", "basic_bare"))

	assert.NilError(t, err)
	assert.Equal(t, basicBareRepo.Name, "basic_bare")
	assert.Equal(t, basicBareRepo.Description, "Basic test repository.")
	assert.Equal(t, basicBareRepo.Author, "adjsky")
}
