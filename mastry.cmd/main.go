package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/martindrlik/mastry.app/catalog"
	"github.com/martindrlik/mastry.app/server"
)

var (
	addr   = flag.String("addr", ":8042", "listen and serve on TCP network address addr")
	source = flag.String("source", "c", "catalog's source file")
	seed   = flag.Int64("seed", 17, "value used to initialize the source to a deterministic state")
)

func main() {
	flag.Parse()
	rand.Seed(*seed)
	c, err := catalog.NewCatalog(
		catalog.SourceFile(*source),
		catalog.Intn(rand.Intn))
	if err != nil {
		fatal("load error: %v", err)
		os.Exit(1)
	}
	if err := server.ListenAndServe(
		server.Addr(*addr),
		server.WithCatalog(c)); err != nil {
		fatal("listen and serve on %v error: %v", *addr, err)
	}
}

func fatal(format string, a ...interface{}) {
	fmt.Printf("mastry.cmd ")
	fmt.Printf(format, a...)
	fmt.Println()
	os.Exit(1)
}
