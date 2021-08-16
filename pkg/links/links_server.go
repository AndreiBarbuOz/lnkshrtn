package links

import (
	"context"
	"fmt"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/links/adapters"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/links/domain"
	links "github.com/AndreiBarbuOz/lnkshrtn/pkg/links/ports"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
)

type HttpServer struct {
	repo domain.Repository
}

func (l HttpServer) CreateLink(w http.ResponseWriter, r *http.Request) {
	postLink := links.LinkObjectSpec{}
	err := render.Decode(r, &postLink)
	if err != nil {
		BadRequest("bad-request", err, w, r)
	}
}

func (l HttpServer) GetLinkById(w http.ResponseWriter, r *http.Request, linkId string) {
	panic("implement me")
}

func (l HttpServer) GetHealth(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprint(w, `{"status": "ok"}`)
}

func (l HttpServer) GetLinks(w http.ResponseWriter, r *http.Request) {

	panic("implement me")
}

var _ links.ServerInterface = (*HttpServer)(nil)

func (l *HttpServer) Run(ctx context.Context, port int) {
	apiRouter := chi.NewRouter()
	links.HandlerFromMux(*l, apiRouter)

	fmt.Printf("Server execution\n")

	s := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: apiRouter,
	}
	s.ListenAndServe()
}

func NewServer(ctx context.Context) *HttpServer {
	var repo domain.Repository = adapters.NewMemoryLinkRepository()
	return &HttpServer{repo}
}
