package catalog

type option func(*Catalog)

// SourceFile sets Catalogs's source file to name.
func SourceFile(name string) option {
	return func(c *Catalog) {
		c.sourceFile = name
	}
}

// Intn sets Catalog's intn func, the source of some value in [0, n) range.
func Intn(fn func(n int) int) option {
	return func(c *Catalog) {
		c.intn = fn
	}
}
