package undirectedgraph

import (
	"github.com/liweiyi88/algorithm/collections"
	"github.com/liweiyi88/algorithm/undirectedgraph"
)

type BreadthFirstPaths struct {
	marked []bool
	edgeTo []int
	s      int
}

func NewBreadthFirstPaths(g undirectedgraph.Graph, s int) *BreadthFirstPaths {
	bfp := &BreadthFirstPaths{
		marked: make([]bool, g.V),
		edgeTo: make([]int, g.V),
		s:      s,
	}

	bfp.bfs(g, s)

	return bfp
}

func (b *BreadthFirstPaths) bfs(g undirectedgraph.Graph, v int) {
	queue := collections.NewQueue[int]()
	b.marked[v] = true
	queue.Enqueue(v)

	for {
		v, ok := queue.Dequeue()
		if !ok {
			break
		}

		for w := range g.Adj[v] {
			if !b.marked[w] {
				b.edgeTo[w] = v
				b.marked[w] = true
				queue.Enqueue(w)
			}
		}
	}
}

func (b *BreadthFirstPaths) HasPathTo(v int) bool {
	return b.marked[v]
}

func (b *BreadthFirstPaths) PathTo(v int) []int {
	if !b.HasPathTo(v) {
		return nil
	}

	path := collections.NewStack[int]()

	for i := v; i != b.s; i = b.edgeTo[i] {
		path.Push(i)
	}

	path.Push(b.s)

	var results []int

	for {
		item, ok := path.Pop()
		if !ok {
			break
		}

		results = append(results, item)
	}

	return results
}
