package repo

import (
	"errors"
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/go-git/go-git/v5/plumbing"
)

var (
	ErrAnnotatedTagNotFound = errors.New("annotated tag not found")
	ErrCommitNotFound       = errors.New("commit not found")
)

type Branch struct {
	Name      string
	Message   string
	Author    string
	UpdatedAt time.Time
}

type Tag struct {
	Name      string
	Message   string
	Author    string
	CreatedAt time.Time
}

type RefsResult struct {
	Tags     []Tag
	Branches []Branch
}

func (repo Repo) Refs() (RefsResult, error) {
	result := RefsResult{}

	rIter, err := repo.raw.References()

	if err != nil {
		return result, ErrFailedToReadRefs
	}

	err = rIter.ForEach(func(ref *plumbing.Reference) error {
		switch {
		case ref.Name().IsTag():
			tagObject, err := repo.raw.TagObject(ref.Hash())

			if err != nil {
				return ErrAnnotatedTagNotFound
			}

			result.Tags = append(result.Tags, Tag{
				Name:      ref.Name().Short(),
				Message:   strings.TrimSpace(tagObject.Message),
				Author:    fmt.Sprintf("%s <%s>", tagObject.Tagger.Name, tagObject.Tagger.Email),
				CreatedAt: tagObject.Tagger.When.UTC(),
			})

		case ref.Name().IsBranch():
			commit, err := repo.raw.CommitObject(ref.Hash())

			if err != nil {
				return ErrCommitNotFound
			}

			result.Branches = append(result.Branches, Branch{
				Name:      ref.Name().Short(),
				Message:   strings.TrimSpace(commit.Message),
				Author:    fmt.Sprintf("%s <%s>", commit.Author.Name, commit.Author.Email),
				UpdatedAt: commit.Author.When.UTC(),
			})
		}

		return nil
	})

	slices.SortFunc(result.Branches, func(a, b Branch) int {
		return b.UpdatedAt.Compare(a.UpdatedAt)
	})

	slices.SortFunc(result.Tags, func(a, b Tag) int {
		return b.CreatedAt.Compare(a.CreatedAt)
	})

	return result, err
}
