package undirectedgraph

import (
	"testing"

	"github.com/liweiyi88/algorithm/undirectedgraph"
	"github.com/stretchr/testify/assert"
)

func TestBFS(t *testing.T) {
	graph := undirectedgraph.CreateTestGraph()

	search := NewBreadthFirstSearch(*graph, 0)

	connected := []int{}
	for i := 0; i < graph.V; i++ {
		if search.Marked(i) {
			connected = append(connected, i)
		}
	}

	assert.Equal(t, connected, []int{0, 1, 2, 3, 4, 5, 6})

	search = NewBreadthFirstSearch(*graph, 9)

	connected = []int{}
	for i := 0; i < graph.V; i++ {
		if search.Marked(i) {
			connected = append(connected, i)
		}
	}

	assert.Equal(t, connected, []int{9, 10, 11, 12})
}
