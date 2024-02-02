package repo_test

import (
	"path"
	"testing"

	"github.com/adjsky/gitjika/internal/repo"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	lgs, err := repo.Log(path.Join("fixtures", "basic"), "HEAD", 20)

	assert.NoError(t, err)
	assert.Len(t, lgs, 2)
}
