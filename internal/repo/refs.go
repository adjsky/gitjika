package repo

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/go-git/go-git/v5/plumbing"
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
		return result, fmt.Errorf("failed to get refs: %w", err)
	}

	err = rIter.ForEach(func(ref *plumbing.Reference) error {
		switch {
		case ref.Name().IsTag():
			tagObject, err := repo.raw.TagObject(ref.Hash())

			if err != nil {
				return err
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
				return err
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

	if err != nil {
		return result, fmt.Errorf("failed to process refs: %w", err)
	}

	slices.SortFunc(result.Branches, func(a, b Branch) int {
		return b.UpdatedAt.Compare(a.UpdatedAt)
	})

	slices.SortFunc(result.Tags, func(a, b Tag) int {
		return b.CreatedAt.Compare(a.CreatedAt)
	})

	return result, nil
}
