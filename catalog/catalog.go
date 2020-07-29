package catalog

import (
	"encoding/json"
	"fmt"
	"os"
)

// Catalog loads and stores problems with solutions.
type Catalog struct {
	sourceFile string
	problems   []Problem
	intn       func(n int) int
}

// NewCatalog creates new catalog.
func NewCatalog(opts ...option) (*Catalog, error) {
	c := new(Catalog)
	for _, opt := range opts {
		opt(c)
	}
	err := c.load()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (c *Catalog) load() error {
	f, err := os.Open(c.sourceFile)
	if err != nil {
		return fmt.Errorf("load %v: %w", c.sourceFile, err)
	}
	defer f.Close()
	dec := json.NewDecoder(f)
	err = dec.Decode(&c.problems)
	if err != nil {
		return fmt.Errorf("decode %v: %w", c.sourceFile, err)
	}
	return nil
}

// Problem returns some of stored problems.
func (c *Catalog) Problem() Problem {
	i := c.intn(len(c.problems))
	return c.problems[i]
}
