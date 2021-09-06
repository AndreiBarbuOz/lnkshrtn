package domain

import (
	"context"
	"fmt"
	"time"
)

type Application struct {
	repo Repository
}

func (a *Application) RequestCreateLink(url *string, shortned *string) (*Link, error) {
	newLink, err := NewLink(*url, *shortned, time.Now())
	if err != nil {
		return nil, fmt.Errorf("could not create new link: %w", err)
	}
	ret, err := a.repo.CreateLink(newLink)
	if err != nil {
		return nil, fmt.Errorf("could not add link to repository: %w", err)
	}
	return ret, nil
}

func (a *Application) GetLinkByShortned(shortned string) (*Link, error) {
	link, err := a.repo.GetLinkByShortned(shortned)
	if err != nil {
		return nil, fmt.Errorf("could not find shortned %s: %w", shortned, err)
	}

	return link, nil
}

func (a *Application) GetAllLinks() ([]*Link, error) {
	links := a.repo.GetLinks()

	return links, nil
}

func NewApplication(ctx context.Context, r Repository) Application {
	return Application{
		repo: r,
	}
}
