package undirectedgraph

import (
	"github.com/liweiyi88/algorithm/undirectedgraph"
)

type DeepFirstSearch struct {
	marked []bool
	count  int
}

func NewDeepFirstSearch(g undirectedgraph.Graph, s int) *DeepFirstSearch {
	dfs := &DeepFirstSearch{
		marked: make([]bool, g.V),
		count:  0,
	}

	dfs.dfs(g, s)

	return dfs
}

func (d *DeepFirstSearch) dfs(g undirectedgraph.Graph, v int) {
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
