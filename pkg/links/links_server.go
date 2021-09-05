package links

import (
	"context"
	"fmt"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/links/adapters"
	"github.com/AndreiBarbuOz/lnkshrtn/pkg/links/domain"
	links "github.com/AndreiBarbuOz/lnkshrtn/pkg/links/ports"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"net/http"
	"time"
)

type HttpServer struct {
	repo domain.Repository
}

func (l HttpServer) CreateLink(w http.ResponseWriter, r *http.Request) {
	postLink := links.LinkObjectSpec{}
	err := render.Decode(r, &postLink)
	if err != nil {
		BadRequest("bad-request", err, w, r)
		return
	}

	newLink, err := domain.NewLink(postLink.Url, *postLink.Shortned, time.Now())
	if err != nil {
		BadRequest("bad-request", err, w, r)
		return
	}
	link, err := l.repo.CreateLink(newLink)
	if err != nil {
		BadRequest("bad-request", err, w, r)
		return
	}

	ret, err := domainToApiLink(link)
	if err != nil {
		BadRequest("bad-request", err, w, r)
		return
	}

	render.Respond(w, r, ret)
}

func (l HttpServer) GetLinkById(w http.ResponseWriter, r *http.Request, linkId string) {
	ret, err := l.repo.GetLink(linkId)
	if err != nil {
		http.NotFound(w, r)
	}

	render.Respond(w, r, ret)
}

func (l HttpServer) GetHealth(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, struct {
		Status string `json:"status"`
	}{Status: "ok"})
}

func (l HttpServer) GetLinks(w http.ResponseWriter, r *http.Request) {

	ret := l.repo.GetLinks()
	render.Respond(w, r, ret)
}

var _ links.ServerInterface = (*HttpServer)(nil)

func (l *HttpServer) Run(ctx context.Context, port int) {
	router := chi.NewRouter()
	router.Use(middleware.AllowContentType("application/json"))
	handler := links.HandlerFromMux(*l, router)

	fmt.Printf("Server execution\n")

	s := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: handler,
	}
	s.ListenAndServe()
}

func NewServer(ctx context.Context) *HttpServer {
	var repo domain.Repository = adapters.NewMemoryLinkRepository()
	return &HttpServer{repo}
}

func domainToApiLink(link *domain.Link) (*links.LinkObject, error) {
	var version links.LinkObjectApiVersion = "v1"
	var ret *links.LinkObject = &links.LinkObject{
		ApiVersion: &version,
		Metadata:   nil,
		Spec: &links.LinkObjectSpec{
			Shortned: &link.Shortned,
			Url:      link.Url,
		},
	}
	return ret, nil
}
