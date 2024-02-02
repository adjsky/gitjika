package repo

import (
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type LogStatement struct {
	Message string
	Author  string
	Date    time.Time
}

func Log(p string, commit string, nCommits uint8) ([]LogStatement, error) {
	repo, err := git.PlainOpen(p)

	if err != nil {
		return nil, err
	}

	lgs := make([]LogStatement, 0, nCommits)

	cIter, err := repo.Log(&git.LogOptions{
		From: plumbing.NewHash(commit),
	})

	if err != nil {
		return nil, err
	}

	for i := uint8(0); i < nCommits; i++ {
		commit, err := cIter.Next()

		if err != nil {
			break
		}

		lgs = append(lgs, LogStatement{
			Message: commit.Message,
			Author:  commit.Author.Name,
			Date:    commit.Author.When,
		})
	}

	return lgs, nil
}
