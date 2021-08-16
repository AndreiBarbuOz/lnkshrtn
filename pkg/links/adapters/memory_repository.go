package adapters

import (
	"fmt"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/links/domain"
	"github.com/pkg/errors"
	"sync"
	"time"
)

type MemoryLinkRepository struct {
	links map[string]domain.Link
	mutex sync.RWMutex
}

func (r *MemoryLinkRepository) GetLinks() []*domain.Link {
	panic("implement me")
}

func (r *MemoryLinkRepository) GetLink(shortned string) (*domain.Link, error) {
	if ret, ok := r.links[shortned]; ok {
		return &ret, nil
	}
	return nil, errors.New(fmt.Sprintf("didn't find %s\n", shortned))
}

func (r *MemoryLinkRepository) CreateLink(url string, shortned string) (*domain.Link, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	link := domain.Link{
		Url:      url,
		Shortned: shortned,
		Created:  time.Now(),
	}
	r.links[url] = link
	return &link, nil
}

func NewMemoryLinkRepository() *MemoryLinkRepository {
	return &MemoryLinkRepository{
		links: make(map[string]domain.Link),
	}
}
