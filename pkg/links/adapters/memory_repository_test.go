package adapters

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMemoryLinkRepository_CreateLink(t *testing.T) {
	repo := NewMemoryLinkRepository()
	l, err := repo.CreateLink("url", "shortned")
	assert.Nil(t, err, "Failed to create new Link")
	assert.NotNil(t, l, "Create new link returned nil object")

	assert.Equal(t, "url", l.Url, "Wrong Link object values")
}

func TestMemoryLinkRepository_CreateLinkEmpty(t *testing.T) {
	repo := NewMemoryLinkRepository()
	_, err := repo.CreateLink("", "shortned")
	assert.NotNil(t, err, "Failed to return error on empty url")
	_, err = repo.CreateLink("url", "")
	assert.NotNil(t, err, "Failed to return error on empty shortned")
}

func TestMemoryLinkRepository_GetLinksEmpty(t *testing.T) {
	repo := NewMemoryLinkRepository()
	links := repo.GetLinks()
	assert.NotNil(t, links, "GetLinks returned empty object")
	assert.Len(t, links, 0, "Links len not zero")
}

func TestMemoryLinkRepository_GetLinksOneElem(t *testing.T) {
	repo := NewMemoryLinkRepository()
	_, _ = repo.CreateLink("url", "shortned")
	links := repo.GetLinks()
	assert.NotNil(t, links, "GetLinks returned empty object")
	assert.Len(t, links, 1, "Links len not equal to 1")
	assert.Equal(t, "url", links[0].Url)
	assert.Equal(t, "shortned", links[0].Shortned)
}

func TestMemoryLinkRepository_GetLink(t *testing.T) {
	repo := NewMemoryLinkRepository()
	_, err := repo.GetLink("shortned")
	assert.NotNil(t, err, "Returned non nil error for non existing link")
	_, _ = repo.CreateLink("url", "shortned")
	l, err := repo.GetLink("shortned")

	assert.Nil(t, err, "Failed to get link")
	assert.NotNil(t, l, "GetLink returned nil link")

	_, err = repo.GetLink("non_existing")
	assert.NotNil(t, err, "Returned non nil error for non existing link")
}
