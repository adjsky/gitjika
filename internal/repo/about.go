package repo

import (
	"errors"
	"fmt"

	"github.com/go-git/go-git/v5/plumbing"
)

var readmeCandidates = [...]string{"README", "README.md"}

var ErrNoReadMe = errors.New("no readme file found")

func (repo Repo) About(refName string) (string, error) {
	ref, err := repo.raw.Reference(plumbing.ReferenceName(refName), false)

	if err != nil {
		return "", ErrRefNotFound
	}

	commit, err := repo.raw.CommitObject(ref.Hash())

	if err != nil {
		return "", ErrCommitNotFound
	}

	for _, candidate := range readmeCandidates {
		file, err := commit.File(candidate)

		if err != nil {
			continue
		}

		contents, err := file.Contents()

		if err != nil {
			return "", fmt.Errorf("found readme, but failed to read it: %w", err)
		}

		return contents, nil
	}

	return "", ErrNoReadMe
}
