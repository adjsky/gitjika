package repo_test

import (
	"path"
	"testing"

	"github.com/adjsky/gitjika/internal/repo"
	"gotest.tools/v3/assert"
)

func TestBranches(t *testing.T) {
	basicBareRepo, err := repo.New(path.Join("fixtures", "repos", "basic_bare"))

	assert.NilError(t, err)

	_, err = basicBareRepo.Branches()

	assert.NilError(t, err)
}
