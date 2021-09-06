package adapters

import (
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/links/domain"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMemoryLinkRepository_CreateLink(t *testing.T) {
	t.Parallel()
	repo := NewMemoryLinkRepository()
	testCases := []struct {
		Name            string
		LinkConstructor func(t *testing.T) *domain.Link
	}{
		{
			Name:            "standard_link",
			LinkConstructor: newStandardLink,
		},
	}

	for _, crtCase := range testCases {
		t.Run(crtCase.Name, func(t *testing.T) {
			newLink := crtCase.LinkConstructor(t)
			l, err := repo.CreateLink(newLink)
			require.NoError(t, err)

			assertPersistedLinkEquals(t, repo, l)
		})
	}
}

func newStandardLink(t *testing.T) *domain.Link {
	t.Helper()
	l, err := domain.NewLink("url", "shortned", time.Now().Round(time.Hour))
	require.NoError(t, err)
	return l
}

func assertPersistedLinkEquals(t *testing.T, repo domain.Repository, l *domain.Link) {
	t.Helper()
	persistedLink, err := repo.GetLinkByShortned(l.Shortned)
	require.NoError(t, err)
	assertLinkEquals(t, persistedLink, l)
}

var cmpRoundTimeOpt = cmp.Comparer(func(x, y time.Time) bool {
	return x.Truncate(time.Millisecond).Equal(y.Truncate(time.Millisecond))
})

func assertLinkEquals(t *testing.T, lnk1 *domain.Link, lnk2 *domain.Link) {
	t.Helper()
	cmpOpts := []cmp.Option{
		cmpRoundTimeOpt,
		cmp.AllowUnexported(time.Time{}),
	}
	assert.True(t, cmp.Equal(lnk1, lnk2, cmpOpts...))
}

func TestMemoryLinkRepository_CreateLinkEmpty(t *testing.T) {
	repo := NewMemoryLinkRepository()
	_, err := repo.CreateLink(nil)
	assert.NotNil(t, err, "Failed to return error on nil Link")
}

func TestMemoryLinkRepository_GetLinksEmpty(t *testing.T) {
	repo := NewMemoryLinkRepository()
	links := repo.GetLinks()
	assert.NotNil(t, links, "GetLinks returned empty object")
	assert.Len(t, links, 0, "Links len not zero")
}

func TestMemoryLinkRepository_GetLinksOneElem(t *testing.T) {
	repo := NewMemoryLinkRepository()
	newLink := newStandardLink(t)
	_, _ = repo.CreateLink(newLink)
	links := repo.GetLinks()
	assert.NotNil(t, links, "GetLinks returned empty object")
	assert.Len(t, links, 1, "Links len not equal to 1")
	assert.Equal(t, "url", links[0].Url)
	assert.Equal(t, "shortned", links[0].Shortned)
}

func TestMemoryLinkRepository_GetLink(t *testing.T) {
	repo := NewMemoryLinkRepository()
	_, err := repo.GetLinkByShortned("shortned")
	assert.NotNil(t, err, "Returned non nil error for non existing link")
	newLink := newStandardLink(t)
	_, _ = repo.CreateLink(newLink)
	l, err := repo.GetLinkByShortned("shortned")

	assert.Nil(t, err, "Failed to get link")
	assert.NotNil(t, l, "GetLinkByShortned returned nil link")

	_, err = repo.GetLinkByShortned("non_existing")
	assert.NotNil(t, err, "Returned non nil error for non existing link")
}
