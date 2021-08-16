package domain

import (
	"time"
)

type Link struct {
	Url      string
	Shortned string
	Created  time.Time
}

type Repository interface {
	CreateLink(url string, shortned string) (*Link, error)

	GetLinks() []*Link
	GetLink(shortned string) (*Link, error)
}
