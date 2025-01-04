package repo_test

import (
	"path"
	"testing"

	"github.com/adjsky/gitjika/internal/repo"
	"gotest.tools/v3/assert"
)

func TestListAll(t *testing.T) {
	repos, err := repo.ListAll(path.Join("fixtures", "repos"))

	assert.NilError(t, err)
	assert.Equal(t, len(repos), 4)
}
