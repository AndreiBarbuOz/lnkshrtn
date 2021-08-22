package adapters

import (
	"fmt"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/links/domain"
	"sync"
	"time"
)

var _ domain.Repository = &MemoryLinkRepository{}

type linkModel struct {
	url      string
	shortned string
	created  time.Time
}

type MemoryLinkRepository struct {
	links map[string]linkModel
	mutex sync.RWMutex
}

func (r *MemoryLinkRepository) GetLinks() []*domain.Link {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	ret := make([]*domain.Link, 0)
	for _, value := range r.links {
		ret = append(ret, newLinkFromModel(&value))
	}
	return ret
}

func (r *MemoryLinkRepository) GetLink(shortned string) (*domain.Link, error) {
	if ret, ok := r.links[shortned]; ok {
		return newLinkFromModel(&ret), nil
	}
	return nil, fmt.Errorf("didn't find %s\n", shortned)
}

func (r *MemoryLinkRepository) CreateLink(l *domain.Link) (*domain.Link, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	if l == nil {
		return nil, fmt.Errorf("cannot create Link from nil argument")
	}

	newLink := linkModel{
		url:      l.Url,
		shortned: l.Shortned,
		created:  l.Created,
	}
	r.links[l.Shortned] = newLink

	return newLinkFromModel(&newLink), nil
}

func NewMemoryLinkRepository() *MemoryLinkRepository {
	return &MemoryLinkRepository{
		links: make(map[string]linkModel),
	}
}

func newLinkFromModel(l *linkModel) *domain.Link {
	return &domain.Link{
		Url:      l.url,
		Shortned: l.shortned,
		Created:  l.created,
	}
}
