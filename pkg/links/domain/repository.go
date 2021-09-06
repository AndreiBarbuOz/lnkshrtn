package domain

import (
	"fmt"
	"time"
)

type Link struct {
	Url      string
	Shortned string
	Created  time.Time
}

func NewLink(url string, shortned string, created time.Time) (*Link, error) {
	if url == "" {
		return nil, fmt.Errorf("NewLink cannot create link with empty url")
	}
	if shortned == "" {
		return nil, fmt.Errorf("NewLink cannot create link with empty shortned")
	}
	if created.IsZero() {
		return nil, fmt.Errorf("NewLink cannot create link with nil created")
	}
	return &Link{url, shortned, created}, nil
}

type Repository interface {
	CreateLink(*Link) (*Link, error)

	GetLinks() []*Link
	GetLinkByShortned(shortned string) (*Link, error)
}
