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

	ret := domainToApiLink(link)

	render.Respond(w, r, ret)
}

func (l HttpServer) GetLinkById(w http.ResponseWriter, r *http.Request, linkId string) {
	link, err := l.app.GetLinkByShortned(linkId)
	if err != nil {
		http.NotFound(w, r)
	}

	ret := domainToApiLink(link)

	render.Respond(w, r, ret)
}

func (l HttpServer) GetHealth(w http.ResponseWriter, r *http.Request) {
	render.Respond(w, r, struct {
		Status string `json:"status"`
	}{Status: "ok"})
}

func (l HttpServer) GetLinks(w http.ResponseWriter, r *http.Request) {

	link, err := l.app.GetAllLinks()
	if err != nil {
		BadRequest("bad-request", err, w, r)
		return
	}
	ret := domainToApiLinkList(link)
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

func domainToApiLink(link *domain.Link) *links.LinkObject {
	var version links.LinkObjectApiVersion = "v1"
	var ret *links.LinkObject = &links.LinkObject{
		ApiVersion: &version,
		Metadata:   nil,
		Spec: links.LinkObjectSpec{
			Shortned: &link.Shortned,
			Url:      link.Url,
		},
	}
	return ret
}

func domainToApiLinkList(l []*domain.Link) *links.LinkObjectList {
	var version string = "v1"

	var ret *links.LinkObjectList = &links.LinkObjectList{
		ApiVersion: (*links.LinkObjectListApiVersion)(&version),
		Metadata:   nil,
		Items:      make([]links.LinkObject, 0),
	}
	for _, crtLink := range l {
		ret.Items = append(ret.Items, *domainToApiLink(crtLink))
	}

	return ret
}
