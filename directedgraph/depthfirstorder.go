package directedgraph

import "github.com/liweiyi88/algorithm/collections"

type DepthFirstOrder struct {
	marked      []bool
	PreOrder    collections.Queue[int]
	PostOrder   collections.Queue[int]
	ReversePost collections.Stack[int]
}

func NewDepthFirstOrder(g Digraph) *DepthFirstOrder {
	d := &DepthFirstOrder{
		marked: make([]bool, g.V),
	}

	for v := 0; v < g.V; v++ {
		if !d.marked[v] {
			d.dfs(g, v)
		}
	}

	return d
}

func (d *DepthFirstOrder) dfs(g Digraph, v int) {
	d.PreOrder.Enqueue(v)

	d.marked[v] = true
	for w := range g.Adj[v] {
		if !d.marked[w] {
			d.dfs(g, w)
		}
	}
	d.PostOrder.Enqueue(v)
	d.ReversePost.Push(v)
}
