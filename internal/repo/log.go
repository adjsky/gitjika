package repo

import (
	"fmt"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type LogStatement struct {
	CommitHash string
	Message    string
	Author     string
	Date       string
}

func Log(path string, commitHash string, size uint8) ([]LogStatement, error) {
	repo, err := git.PlainOpen(path)

	if err != nil {
		return nil, err
	}

	lgs := make([]LogStatement, 0, size)

	cIter, err := repo.Log(&git.LogOptions{
		From: plumbing.NewHash(commitHash),
	})

	if err != nil {
		return nil, err
	}

	for i := uint8(0); i < size; i++ {
		commit, err := cIter.Next()

		if err != nil {
			break
		}

		lgs = append(lgs, LogStatement{
			Message:    strings.TrimSpace(commit.Message),
			Author:     fmt.Sprintf("%s <%s>", commit.Author.Name, commit.Author.Email),
			Date:       commit.Author.When.UTC().Format("Mon Jan 2 15:04:05 2006"),
			CommitHash: commit.Hash.String(),
		})
	}

	return lgs, nil
}
