package undirectedgraph

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

func CreateTestGraph() *Graph {
	edges := make([][]int, 0, 13)

	edges = append(edges, []int{0, 5})
	edges = append(edges, []int{4, 3})
	edges = append(edges, []int{0, 1})
	edges = append(edges, []int{9, 12})
	edges = append(edges, []int{6, 4})
	edges = append(edges, []int{5, 4})
	edges = append(edges, []int{0, 2})
	edges = append(edges, []int{11, 12})
	edges = append(edges, []int{9, 10})
	edges = append(edges, []int{0, 6})
	edges = append(edges, []int{7, 8})
	edges = append(edges, []int{9, 11})
	edges = append(edges, []int{5, 3})

	return NewGraph(13, 13, edges)
}
