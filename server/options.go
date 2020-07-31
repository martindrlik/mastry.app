package server

import (
	"github.com/martindrlik/mastry.app/catalog"
)

type option func(*server)

// Addr sets server's listen and serve TCP network address to addr.
func Addr(addr string) option {
	return func(srv *server) {
		srv.addr = addr
	}
}

// WithCatalog sets server's catalog to c.
func WithCatalog(c *catalog.Catalog) option {
	return func(srv *server) {
		srv.ctg = c
	}
}
