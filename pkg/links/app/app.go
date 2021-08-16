package app

import "context"

type Application struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
}

type Queries struct {
}

func NewApplication(ctx context.Context) Application {
	return Application{
		Commands: Commands{},
		Queries:  Queries{},
	}
}
