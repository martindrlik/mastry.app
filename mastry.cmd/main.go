package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/martindrlik/mastry.app/catalog"
)

var (
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
		fmt.Printf("mastry.cmd load error: %v", err)
		os.Exit(1)
	}
	printProblem(c.Problem())
}

func printProblem(p catalog.Problem) {
	fmt.Println(p.Description)
	for i, solution := range p.Solutions {
		fmt.Printf("%d: %s\n", i, solution.Description)
	}
}
