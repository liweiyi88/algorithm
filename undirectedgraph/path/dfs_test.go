package undirectedgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/liweiyi88/algorithm/undirectedgraph"
)

func TestDFSPath(t *testing.T) {
	graph := undirectedgraph.CreateTestGraph()
	dfsPath := NewDeepFirstPaths(*graph, 11)
	path := dfsPath.PathTo(10)

	assert.Equal(t, 11, path[0])
	assert.Equal(t, 10, path[len(path)-1])

	dfsPath = NewDeepFirstPaths(*graph, 0)
	path = dfsPath.PathTo(4)

	assert.Equal(t, 0, path[0])
	assert.Equal(t, 4, path[len(path)-1])
}
