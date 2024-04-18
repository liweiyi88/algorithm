package graph

import (
	"github.com/liweiyi88/algorithm/collections"
	"github.com/liweiyi88/algorithm/graph"
)

type DeepFirstSearch struct {
	marked []bool
	count  int
}

func NewDeepFirstSearch(g graph.Graph, s int) *DeepFirstSearch {
	dfs := &DeepFirstSearch{
		marked: make([]bool, g.V),
		count:  0,
	}

	dfs.dfs(g, s)

	return dfs
}

func (d *DeepFirstSearch) dfs(g graph.Graph, v int) {
	d.marked[v] = true
	d.count++
	for w := range g.Adj[v] {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
}

func (d *DeepFirstSearch) Marked(v int) bool {
	return d.marked[v]
}

func (d *DeepFirstSearch) Count() int {
	return d.count
}

type DeepFirstPaths struct {
	marked []bool
	edgeTo []int
	s      int
}

func NewDeepFirstPaths(g graph.Graph, s int) *DeepFirstPaths {
	dfp := &DeepFirstPaths{
		marked: make([]bool, g.V),
		edgeTo: make([]int, g.V),
		s:      s,
	}

	dfp.dfs(g, s)

	return dfp
}

func (d *DeepFirstPaths) dfs(g graph.Graph, v int) {
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
