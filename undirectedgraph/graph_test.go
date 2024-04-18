package undirectedgraph

import (
	"testing"
)

func TestGraph(t *testing.T) {
	graph := CreateTestGraph()
	t.Log(graph.ToString())
}
