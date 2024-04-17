package graph

import (
	"slices"
	"testing"

	"github.com/liweiyi88/algorithm/graph"
)

func TestDFS(t *testing.T) {
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

	graph := graph.NewGraph(13, 13, edges)

	search := NewDeepFirstSearch(*graph, 0)

	connected := []int{}
	for i := 0; i < graph.V; i++ {
		if search.Marked(i) {
			connected = append(connected, i)
		}
	}

	if !slices.Equal(connected, []int{0, 1, 2, 3, 4, 5, 6}) {
		t.Error("unexpected search result.")
	}

	search = NewDeepFirstSearch(*graph, 9)

	connected = []int{}
	for i := 0; i < graph.V; i++ {
		if search.Marked(i) {
			connected = append(connected, i)
		}
	}

	if !slices.Equal(connected, []int{9, 10, 11, 12}) {
		t.Error("unexpected search result.")
	}
}
