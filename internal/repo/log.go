package repo

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/plumbing/object"
)

type LogStatement struct {
	CommitHash   string
	Message      string
	Author       string
	Date         time.Time
	LinesDeleted int
	LinesAdded   int
	References   []string
}

func (repo Repo) Log(commitHash string, size uint8) ([]LogStatement, error) {
	rIter, err := repo.raw.References()

	if err != nil {
		return nil, err
	}

	refm := make(map[plumbing.Hash][]string)

	err = rIter.ForEach(func(ref *plumbing.Reference) error {
		hash := ref.Hash()

		refm[hash] = append(refm[hash], ref.Name().Short())

		return nil
	})

	if err != nil {
		return nil, err
	}

	lgs := make([]LogStatement, 0, size)

	cIter, err := repo.raw.Log(&git.LogOptions{
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

		totalCommitStats, err := getTotalCommitStats(commit)

		if err != nil {
			return nil, err
		}

		lgs = append(lgs, LogStatement{
			Message:      strings.TrimSpace(commit.Message),
			Author:       fmt.Sprintf("%s <%s>", commit.Author.Name, commit.Author.Email),
			Date:         commit.Author.When,
			CommitHash:   commit.Hash.String(),
			LinesDeleted: totalCommitStats.LinesDeleted,
			LinesAdded:   totalCommitStats.LinesAdded,
			References:   refm[commit.Hash],
		})
	}

	return lgs, nil
}

func (repo Repo) LogRef(name string, size uint8) ([]LogStatement, error) {
	ref, err := repo.raw.Reference(plumbing.ReferenceName(name), false)

	if err != nil {
		return nil, err
	}

	return repo.Log(ref.Hash().String(), size)
}

type totalCommitStats struct {
	LinesDeleted int
	LinesAdded   int
}

func getTotalCommitStats(commit *object.Commit) (totalCommitStats, error) {
	totalStats := totalCommitStats{}

	fileStats, err := commit.Stats()

	if err != nil {
		return totalStats, err
	}

	for _, fileStat := range fileStats {
		totalStats.LinesAdded += fileStat.Addition
		totalStats.LinesDeleted += fileStat.Deletion
	}

	return totalStats, nil
}
