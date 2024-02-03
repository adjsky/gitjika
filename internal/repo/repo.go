package repo

import "github.com/go-git/go-git/v5"

type Repo struct {
	raw *git.Repository
}

func New(path string) (Repo, error) {
	repo, err := git.PlainOpen(path)

	if err != nil {
		return Repo{}, err
	}

	return Repo{
		raw: repo,
	}, nil
}
