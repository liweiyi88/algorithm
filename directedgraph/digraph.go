package directedgraph

import (
	"github.com/liweiyi88/algorithm/collections"
)

type Digraph struct {
	V   int                    // number of vertices
	E   int                    // number of edges
	Adj []collections.Bag[int] // adjacency list
}

func NewDiGraph(v int, edges [][]int) *Digraph {
	g := &Digraph{
		V: v,
		E: 0,
	}

	g.Adj = make([]collections.Bag[int], g.V)

	for v := 0; v < g.V; v++ {
		g.Adj[v] = collections.NewBag[int]()
	}

	for _, edge := range edges {
		v, w := edge[0], edge[1]

		g.AddEdge(v, w)
	}

	return g
}

func (d *Digraph) AddEdge(v int, w int) {
	d.Adj[v].Add(w)
	d.E++
}

func (d *Digraph) reverse() *Digraph {
	newGraph := &Digraph{V: d.V}

	for v := 0; v < d.V; v++ {
		for w := range d.Adj[v] {
			newGraph.AddEdge(w, v)
		}
	}

	return newGraph
}
