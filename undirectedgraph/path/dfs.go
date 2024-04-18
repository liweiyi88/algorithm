package undirectedgraph

import (
	"github.com/liweiyi88/algorithm/collections"
	"github.com/liweiyi88/algorithm/undirectedgraph"
)

type DeepFirstPaths struct {
	marked []bool
	edgeTo []int
	s      int
}

func NewDeepFirstPaths(g undirectedgraph.Graph, s int) *DeepFirstPaths {
	dfp := &DeepFirstPaths{
		marked: make([]bool, g.V),
		edgeTo: make([]int, g.V),
		s:      s,
	}

	dfp.dfs(g, s)

	return dfp
}

func (d *DeepFirstPaths) dfs(g undirectedgraph.Graph, v int) {
	d.marked[v] = true

	for w := range g.Adj[v] {
		if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		}
	}
}

func (d *DeepFirstPaths) HasPathTo(v int) bool {
	return d.marked[v]
}

func (d *DeepFirstPaths) PathTo(v int) []int {
	if !d.HasPathTo(v) {
		return nil
	}

	path := collections.NewStack[int]()

	for i := v; i != d.s; i = d.edgeTo[i] {
		path.Push(i)
	}

	path.Push(d.s)

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
