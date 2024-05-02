package directedgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDirectedCycle(t *testing.T) {
	edges := make([][]int, 0, 5)
	edges = append(edges, []int{0, 5})
	edges = append(edges, []int{1, 2})
	edges = append(edges, []int{5, 4})
	edges = append(edges, []int{4, 3})
	edges = append(edges, []int{3, 5})

	g := NewDiGraph(6, edges)
	dc := NewDirectedCycle(*g)

	assert.True(t, dc.HasCycle())

	cycle := make([]int, 0)
	for !dc.Cycle.IsEmpty() {
		v, ok := dc.Cycle.Pop()
		if ok {
			cycle = append(cycle, v)
		}
	}

	assert.Equal(t, []int{3, 5, 4, 3}, cycle)
}
