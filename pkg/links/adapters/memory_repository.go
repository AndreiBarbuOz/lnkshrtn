package adapters

import (
	"fmt"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/links/domain"
	"sync"
	"time"
)

type MemoryLinkRepository struct {
	links map[string]domain.Link
	mutex sync.RWMutex
}

func (r *MemoryLinkRepository) GetLinks() []*domain.Link {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	ret := make([]*domain.Link, 0)
	for _, value := range r.links {
		ret = append(ret, &value)
	}
	return ret
}

func (r *MemoryLinkRepository) GetLink(shortned string) (*domain.Link, error) {
	if ret, ok := r.links[shortned]; ok {
		return &ret, nil
	}
	return nil, fmt.Errorf("didn't find %s\n", shortned)
}

func (r *MemoryLinkRepository) CreateLink(url string, shortned string) (*domain.Link, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if url == "" || shortned == "" {
		return nil, fmt.Errorf("cannot create link with empty url or link")
	}

	link := domain.Link{
		Url:      url,
		Shortned: shortned,
		Created:  time.Now(),
	}
	r.links[shortned] = link
	return &link, nil
}

func NewMemoryLinkRepository() *MemoryLinkRepository {
	return &MemoryLinkRepository{
		links: make(map[string]domain.Link),
	}
}
