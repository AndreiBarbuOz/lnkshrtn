package links

import (
	"context"
	"fmt"
	"net/http"
	"os"
)

type LnkshrtnServer struct {
	tmp int
}

func NewServer(ctx context.Context) *LnkshrtnServer {
	return &LnkshrtnServer{0}
}

func index(w http.ResponseWriter, r *http.Request) {
	var name, _ = os.Hostname()
	fmt.Printf("handling request\n")
	fmt.Fprintf(w, "This request was processed by host: %s\n", name)
}

func (l *LnkshrtnServer) Run(ctx context.Context, port int) {
	fmt.Printf("Server execution\n")
	mux := http.NewServeMux()
	fmt.Fprintf(os.Stdout, "Web Server started. Listening on 0.0.0.0:8080\n")
	mux.HandleFunc("/", index)
	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()

}
