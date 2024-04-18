package undirectedgraph

import (
	"testing"

	"github.com/liweiyi88/algorithm/undirectedgraph"
	"github.com/stretchr/testify/assert"
)

func TestBFSPath(t *testing.T) {
	graph := undirectedgraph.CreateTestGraph()
	bfsPath := NewBreadthFirstPaths(*graph, 11)
	path := bfsPath.PathTo(10)

	assert := assert.New(t)

	assert.Equal(11, path[0])
	assert.Equal(10, path[len(path)-1])
	assert.Equal([]int{11, 9, 10}, path)

	bfsPath = NewBreadthFirstPaths(*graph, 0)
	path = bfsPath.PathTo(1)
	assert.Equal([]int{0, 1}, path)
}
