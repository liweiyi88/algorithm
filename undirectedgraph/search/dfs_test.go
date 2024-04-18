package undirectedgraph

import (
	"slices"
	"testing"

	"github.com/liweiyi88/algorithm/undirectedgraph"
	"github.com/stretchr/testify/assert"
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

	dfsPath := NewDeepFirstPaths(*graph, 11)
	path := dfsPath.PathTo(10)

	assert.Equal(t, 11, path[0])
	assert.Equal(t, 10, path[len(path)-1])

	dfsPath = NewDeepFirstPaths(*graph, 0)
	path = dfsPath.PathTo(4)

	assert.Equal(t, 0, path[0])
	assert.Equal(t, 4, path[len(path)-1])
}
