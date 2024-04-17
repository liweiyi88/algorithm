package graph

type Search interface {
	Marked(v int) bool // is v connected to s
	Count() int        // how many vertices are connected to s
}
