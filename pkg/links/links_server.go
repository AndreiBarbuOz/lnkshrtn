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
)

type HttpServer struct {
	app domain.Application
}

var _ links.ServerInterface = (*HttpServer)(nil)

func (l HttpServer) CreateLink(w http.ResponseWriter, r *http.Request) {
	postLink := links.LinkObjectSpec{}
	err := render.Decode(r, &postLink)
	if err != nil {
		BadRequest("bad-request", err, w, r)
		return
	}

	link, err := l.app.RequestCreateLink(&postLink.Url, postLink.Shortned)
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
	link, err := l.app.GetLinkByShortned(linkId)
	if err != nil {
		http.NotFound(w, r)
	}

	ret, err := domainToApiLink(link)
	if err != nil {
		BadRequest("bad-request", err, w, r)
		return
	}

	render.Respond(w, r, ret)
}

func (l HttpServer) GetHealth(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, struct {
		Status string `json:"status"`
	}{Status: "ok"})
}

func (l HttpServer) GetLinks(w http.ResponseWriter, r *http.Request) {

	ret, err := l.app.GetAllLinks()
	if err != nil {
		BadRequest("bad-request", err, w, r)
		return
	}
	render.Respond(w, r, ret)
}

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
	var r domain.Repository = adapters.NewMemoryLinkRepository()
	var app domain.Application = domain.NewApplication(ctx, r)
	return &HttpServer{app}
}

func domainToApiLink(link *domain.Link) (*links.LinkObject, error) {
	var version links.LinkObjectApiVersion = "v1"
	var ret *links.LinkObject = &links.LinkObject{
		ApiVersion: &version,
		Metadata:   nil,
		Spec: links.LinkObjectSpec{
			Shortned: &link.Shortned,
			Url:      link.Url,
		},
	}
	return ret, nil
}
