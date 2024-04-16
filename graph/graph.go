package graph

import (
	"fmt"
	"strings"

	"github.com/liweiyi88/algorithm/collections"
)

type Graph struct {
	V   int                    // number of vertices
	E   int                    // number of edges
	Adj []collections.Bag[int] // adjacency list
}

func NewGraph(v int, e int, edges [][]int) *Graph {
	g := &Graph{
		V: v,
		E: e,
	}

	adj := make([]collections.Bag[int], g.V)

	for v := 0; v < g.V; v++ {
		adj[v] = collections.NewBag[int]()
	}

	g.Adj = adj

	for _, edge := range edges {
		v, w := edge[0], edge[1]
		g.AddEdge(v, w)
	}

	return g
}

func (g *Graph) AddEdge(v int, w int) {
	g.Adj[v].Add(w)
	g.Adj[w].Add(v)
	g.E++
}

func (g *Graph) ToString() string {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("%d vertices, %d edges\n", g.V, g.E))

	for v := 0; v < g.V; v++ {
		sb.WriteString(fmt.Sprintf("%d :", v))
		for w := range g.Adj[v] {
			sb.WriteString(fmt.Sprintf("%d ", w))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}
