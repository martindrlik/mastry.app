package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/martindrlik/mastry.app/catalog"
)

type server struct {
	addr string
	ctg  *catalog.Catalog
}

// ListenAndServe listens and serves according to options opts.
func ListenAndServe(opts ...option) error {
	srv := new(server)
	for _, opt := range opts {
		opt(srv)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/problem", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		err := enc.Encode(srv.Problem())
		if err != nil {
			log.Printf("/problem encoding error: %v", err)
			http.Error(w, "something wrong on our side", http.StatusInternalServerError)
		}
	})
	return http.ListenAndServe(srv.addr, mux)
}
