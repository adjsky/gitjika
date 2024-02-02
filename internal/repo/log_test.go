package repo_test

import (
	"path"
	"testing"

	"github.com/adjsky/gitjika/internal/repo"
	"github.com/stretchr/testify/assert"
)

func TestLog(t *testing.T) {
	lgs, err := repo.Log(path.Join("fixtures", "repos", "basic_bare"), "HEAD", 20)

	assert.NoError(t, err)
	assert.Equal(t, lgs, []repo.LogStatement{
		{
			CommitHash: "31d73f2d0d2fb0ad4c389b435885c50f060f80c4",
			Message:    "add code.js",
			Author:     "adjsky <igorlfmartins@mail.ru>",
			Date:       "Fri Feb 2 15:07:42 2024",
		},
		{
			CommitHash: "722306ed5764ec0cb6cc841c5e569f298b79e63b",
			Message:    "add license",
			Author:     "adjsky <igorlfmartins@mail.ru>",
			Date:       "Fri Feb 2 15:07:26 2024",
		},
		{
			CommitHash: "f243af9ca7948209b97c9c7956ad26d4c1237829",
			Message:    "change readme",
			Author:     "adjsky <igorlfmartins@mail.ru>",
			Date:       "Fri Feb 2 15:07:11 2024",
		},
		{
			CommitHash: "bef512e4ff5027ed895b90d19506b05bf6faab65",
			Message:    "add readme",
			Author:     "adjsky <igorlfmartins@mail.ru>",
			Date:       "Fri Feb 2 15:05:47 2024",
		},
	})
}
