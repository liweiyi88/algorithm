package graph

import (
	"testing"
)

func TestGraph(t *testing.T) {
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

	graph := NewGraph(13, 13, edges)
	t.Log(graph.ToString())
}
