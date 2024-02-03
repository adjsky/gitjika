package repo_test

import (
	"path"
	"testing"
	"time"

	"github.com/adjsky/gitjika/internal/repo"
	"gotest.tools/v3/assert"
)

func TestLog(t *testing.T) {
	basicBareRepo, err := repo.New(path.Join("fixtures", "repos", "basic_bare"))

	assert.NilError(t, err)

	lgs, err := basicBareRepo.Log("HEAD", 20)

	assert.NilError(t, err)
	assert.DeepEqual(t, lgs, []repo.LogStatement{
		{
			CommitHash:   "c4f56d7ad71e1f930b9a3fed3cf1f430905ad566",
			Message:      "change readme 2",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         time.Date(2024, time.February, 3, 23, 9, 10, 0, time.FixedZone("UTC+0300", 3*60*60)),
			LinesDeleted: 1,
			LinesAdded:   7,
			References:   []string{"master"},
		},
		{
			CommitHash:   "722306ed5764ec0cb6cc841c5e569f298b79e63b",
			Message:      "add license",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         time.Date(2024, time.February, 2, 18, 7, 26, 0, time.FixedZone("UTC+0300", 3*60*60)),
			LinesDeleted: 0,
			LinesAdded:   0,
			References:   nil,
		},
		{
			CommitHash:   "f243af9ca7948209b97c9c7956ad26d4c1237829",
			Message:      "change readme",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         time.Date(2024, time.February, 2, 18, 7, 11, 0, time.FixedZone("UTC+0300", 3*60*60)),
			LinesDeleted: 0,
			LinesAdded:   1,
			References:   nil,
		},
		{
			CommitHash:   "bef512e4ff5027ed895b90d19506b05bf6faab65",
			Message:      "add readme",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         time.Date(2024, time.February, 2, 18, 5, 47, 0, time.FixedZone("", 3*60*60)),
			LinesDeleted: 0,
			LinesAdded:   0,
			References:   nil,
		},
	})
}

func TestLogBranch(t *testing.T) {
	basicBareRepo, err := repo.New(path.Join("fixtures", "repos", "basic_bare"))

	assert.NilError(t, err)

	lgs, err := basicBareRepo.LogRef("refs/heads/test-branch", 20)

	assert.NilError(t, err)
	assert.DeepEqual(t, lgs, []repo.LogStatement{
		{
			CommitHash:   "45c3f2fc3726fc48bf9e22757090b420e2d52976",
			Message:      "add some code",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         time.Date(2024, time.February, 3, 23, 18, 12, 0, time.FixedZone("UTC+0300", 3*60*60)),
			LinesDeleted: 0,
			LinesAdded:   2,
			References:   []string{"test-branch"},
		},
		{
			CommitHash:   "c4f56d7ad71e1f930b9a3fed3cf1f430905ad566",
			Message:      "change readme 2",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         time.Date(2024, time.February, 3, 23, 9, 10, 0, time.FixedZone("UTC+0300", 3*60*60)),
			LinesDeleted: 1,
			LinesAdded:   7,
			References:   []string{"master"},
		},
		{
			CommitHash:   "722306ed5764ec0cb6cc841c5e569f298b79e63b",
			Message:      "add license",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         time.Date(2024, time.February, 2, 18, 7, 26, 0, time.FixedZone("UTC+0300", 3*60*60)),
			LinesDeleted: 0,
			LinesAdded:   0,
			References:   nil,
		},
		{
			CommitHash:   "f243af9ca7948209b97c9c7956ad26d4c1237829",
			Message:      "change readme",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         time.Date(2024, time.February, 2, 18, 7, 11, 0, time.FixedZone("UTC+0300", 3*60*60)),
			LinesDeleted: 0,
			LinesAdded:   1,
			References:   nil,
		},
		{
			CommitHash:   "bef512e4ff5027ed895b90d19506b05bf6faab65",
			Message:      "add readme",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         time.Date(2024, time.February, 2, 18, 5, 47, 0, time.FixedZone("UTC+0300", 3*60*60)),
			LinesDeleted: 0,
			LinesAdded:   0,
			References:   nil,
		},
	})
}
