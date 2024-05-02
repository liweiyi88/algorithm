package directedgraph

import "github.com/liweiyi88/algorithm/collections"

type DirectedCycle struct {
	marked  []bool
	edgeTo  []int
	Cycle   collections.Stack[int]
	onStack []bool
}

func NewDirectedCycle(g Digraph) *DirectedCycle {
	marked := make([]bool, g.V)
	edgeTo := make([]int, g.V)
	onStack := make([]bool, g.V)
	Cycle := collections.NewStack[int]()

	directedCycle := &DirectedCycle{
		marked,
		edgeTo,
		*Cycle,
		onStack,
	}

	for v := 0; v < g.V; v++ {
		if !directedCycle.marked[v] {
			directedCycle.dfs(g, v)
		}
	}

	return directedCycle
}

func (d *DirectedCycle) dfs(g Digraph, v int) {
	d.onStack[v] = true
	d.marked[v] = true

	for w := range g.Adj[v] {
		if d.HasCycle() {
			return
		}
		if !d.marked[w] {
			d.edgeTo[w] = v
			d.dfs(g, w)
		} else if d.onStack[w] {
			for x := v; x != w; x = d.edgeTo[x] {
				d.Cycle.Push(x)
			}
			d.Cycle.Push(w)
			d.Cycle.Push(v)	
		}
	}

	d.onStack[v] = false
}

func (d *DirectedCycle) HasCycle() bool {
	return d.Cycle.Size() != 0
}
