package directedgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopologicalSort(t *testing.T) {
	edges := make([][]int, 0, 13)

	edges = append(edges, []int{0, 5})
	edges = append(edges, []int{0, 1})
	edges = append(edges, []int{0, 6})
	edges = append(edges, []int{5, 4})
	edges = append(edges, []int{2, 3})
	edges = append(edges, []int{2, 0})
	edges = append(edges, []int{3, 5})
	edges = append(edges, []int{8, 7})
	edges = append(edges, []int{7, 6})
	edges = append(edges, []int{6, 9})
	edges = append(edges, []int{6, 4})
	edges = append(edges, []int{9, 11})
	edges = append(edges, []int{9, 12})
	edges = append(edges, []int{9, 10})
	edges = append(edges, []int{11, 12})

	g := NewDiGraph(13, edges)
	ts := NewTopologicalSort(*g)

	assert.True(t, ts.HasOrder())
	assert.Equal(t, []int{8, 7, 2, 3, 0, 6, 9, 10, 11, 12, 1, 5, 4}, ts.Order)
}
