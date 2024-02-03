package repo

import (
	"fmt"
	"time"

	"github.com/go-git/go-git/v5/plumbing"
)

type Branch struct {
	Name              string
	LastCommitMessage string
	LastCommitAuthor  string
	UpdatedAt         time.Time
}

func (repo Repo) Branches() ([]Branch, error) {
	rIter, err := repo.raw.References()

	if err != nil {
		return nil, err
	}

	branches := make([]Branch, 0)

	err = rIter.ForEach(func(ref *plumbing.Reference) error {
		if !ref.Name().IsBranch() {
			return nil
		}

		commit, err := repo.raw.CommitObject(ref.Hash())

		if err != nil {
			return err
		}

		branches = append(branches, Branch{
			Name:              ref.Name().Short(),
			LastCommitMessage: commit.Message,
			LastCommitAuthor:  fmt.Sprintf("%s <%s>", commit.Author.Name, commit.Author.Email),
			UpdatedAt:         commit.Author.When,
		})

		return nil
	})

	if err != nil {
		return nil, err
	}

	return branches, nil
}
