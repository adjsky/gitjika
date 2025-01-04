package repo_test

import (
	"path"
	"testing"
	"time"

	"github.com/adjsky/gitjika/internal/repo"
	"gotest.tools/v3/assert"
)

func TestRefs(t *testing.T) {
	refsBareRepo, err := repo.New(path.Join("fixtures", "repos", "refs_bare"))

	assert.NilError(t, err)

	refs, err := refsBareRepo.Refs()

	assert.NilError(t, err)

	assert.DeepEqual(t, refs.Branches, []repo.Branch{
		{
			Name:      "code",
			Message:   "add code",
			Author:    "adjsky <igorlfmartins@mail.ru>",
			UpdatedAt: time.Date(2024, time.February, 4, 18, 31, 15, 0, time.UTC),
		},
		{
			Name:      "readme",
			Message:   "update readme",
			Author:    "adjsky <igorlfmartins@mail.ru>",
			UpdatedAt: time.Date(2024, time.February, 4, 18, 26, 27, 0, time.UTC),
		},
		{
			Name:      "master",
			Message:   "init",
			Author:    "adjsky <igorlfmartins@mail.ru>",
			UpdatedAt: time.Date(2024, time.February, 4, 18, 25, 49, 0, time.UTC),
		},
	})

	assert.DeepEqual(t, refs.Tags, []repo.Tag{
		{
			Name:      "v1.0",
			Message:   "version 1.0",
			Author:    "adjsky <igorlfmartins@mail.ru>",
			CreatedAt: time.Date(2024, time.February, 4, 18, 29, 25, 0, time.UTC),
		},
	})
}
