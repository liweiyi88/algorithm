package undirectedgraph

import (
	"slices"
	"testing"

	"github.com/liweiyi88/algorithm/undirectedgraph"
)

func TestDFS(t *testing.T) {
	graph := undirectedgraph.CreateTestGraph()

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
