package repo_test

import (
	"path"
	"testing"

	"github.com/adjsky/gitjika/internal/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLog(t *testing.T) {
	basicBareRepo, err := repo.New(path.Join("fixtures", "repos", "basic_bare"))

	assert.NoError(t, err)

	lgs, err := basicBareRepo.Log("HEAD", 20)

	require.NoError(t, err)
	assert.Equal(t, lgs, []repo.LogStatement{
		{
			CommitHash:   "c4f56d7ad71e1f930b9a3fed3cf1f430905ad566",
			Message:      "change readme 2",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         "Sat Feb 3 20:09:10 2024",
			LinesDeleted: 1,
			LinesAdded:   7,
			References:   []string{"master"},
		},
		{
			CommitHash:   "722306ed5764ec0cb6cc841c5e569f298b79e63b",
			Message:      "add license",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         "Fri Feb 2 15:07:26 2024",
			LinesDeleted: 0,
			LinesAdded:   0,
			References:   nil,
		},
		{
			CommitHash:   "f243af9ca7948209b97c9c7956ad26d4c1237829",
			Message:      "change readme",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         "Fri Feb 2 15:07:11 2024",
			LinesDeleted: 0,
			LinesAdded:   1,
			References:   nil,
		},
		{
			CommitHash:   "bef512e4ff5027ed895b90d19506b05bf6faab65",
			Message:      "add readme",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         "Fri Feb 2 15:05:47 2024",
			LinesDeleted: 0,
			LinesAdded:   0,
			References:   nil,
		},
	})
}

func TestLogBranch(t *testing.T) {
	basicBareRepo, err := repo.New(path.Join("fixtures", "repos", "basic_bare"))

	assert.NoError(t, err)

	lgs, err := basicBareRepo.LogRef("refs/heads/test-branch", 20)

	require.NoError(t, err)
	assert.Equal(t, lgs, []repo.LogStatement{
		{
			CommitHash:   "45c3f2fc3726fc48bf9e22757090b420e2d52976",
			Message:      "add some code",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         "Sat Feb 3 20:18:12 2024",
			LinesDeleted: 0,
			LinesAdded:   2,
			References:   []string{"test-branch"},
		},
		{
			CommitHash:   "c4f56d7ad71e1f930b9a3fed3cf1f430905ad566",
			Message:      "change readme 2",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         "Sat Feb 3 20:09:10 2024",
			LinesDeleted: 1,
			LinesAdded:   7,
			References:   []string{"master"},
		},
		{
			CommitHash:   "722306ed5764ec0cb6cc841c5e569f298b79e63b",
			Message:      "add license",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         "Fri Feb 2 15:07:26 2024",
			LinesDeleted: 0,
			LinesAdded:   0,
			References:   nil,
		},
		{
			CommitHash:   "f243af9ca7948209b97c9c7956ad26d4c1237829",
			Message:      "change readme",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         "Fri Feb 2 15:07:11 2024",
			LinesDeleted: 0,
			LinesAdded:   1,
			References:   nil,
		},
		{
			CommitHash:   "bef512e4ff5027ed895b90d19506b05bf6faab65",
			Message:      "add readme",
			Author:       "adjsky <igorlfmartins@mail.ru>",
			Date:         "Fri Feb 2 15:05:47 2024",
			LinesDeleted: 0,
			LinesAdded:   0,
			References:   nil,
		},
	})
}
