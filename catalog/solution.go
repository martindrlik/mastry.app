package catalog

type Problem struct {
	Description string
	Solutions   []Solution
}

type Solution struct {
	Description string
	IsCorrect   bool
}
