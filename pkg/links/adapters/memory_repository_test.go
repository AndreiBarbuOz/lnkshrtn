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
}
