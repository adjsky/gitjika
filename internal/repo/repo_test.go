package repo_test

import (
	"path"
	"testing"
	"time"

	"github.com/adjsky/gitjika/internal/repo"
	"gotest.tools/v3/assert"
)

func TestOpen(t *testing.T) {
	type expectedData struct {
		name        string
		description string
		author      string
		age         time.Time
	}

	tests := []struct {
		repo     string
		expected expectedData
	}{
		{repo: "basic_bare", expected: expectedData{
			name:        "basic_bare",
			description: "Basic test repository.",
			author:      "adjsky",
			age:         time.Time{},
		}},
		{repo: "empty_bare", expected: expectedData{
			name:        "empty_bare",
			description: "Unnamed repository; edit this file 'description' to name the repository.",
			author:      "",
			age:         time.Time{},
		}},
		{repo: "agefile_bare", expected: expectedData{
			name:        "agefile_bare",
			description: "Unnamed repository; edit this file 'description' to name the repository.",
			author:      "",
			age:         time.Date(2024, time.February, 5, 20, 52, 26, 0, time.UTC),
		}},
	}

	for _, test := range tests {
		t.Run(test.repo, func(t *testing.T) {
			testRepo, err := repo.Open(path.Join("fixtures", "repos", test.repo))

			assert.NilError(t, err)

			description, err := testRepo.Description()
			assert.NilError(t, err)
			assert.Equal(t, description, test.expected.description)

			age, err := testRepo.Age()
			assert.Equal(t, age, test.expected.age)

			assert.Equal(t, testRepo.Name(), test.expected.name)
			assert.Equal(t, testRepo.Author(), test.expected.author)
		})
	}
}
