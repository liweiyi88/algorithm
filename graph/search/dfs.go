package graph

import "github.com/liweiyi88/algorithm/graph"

type DeepFirstSearch struct {
	marked []bool
	count  int
}

func NewDeepFirstSearch(g graph.Graph, s int) *DeepFirstSearch {
	search := &DeepFirstSearch{
		marked: make([]bool, g.V),
		count:  0,
	}

	search.dfs(g, s)

	return search
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
