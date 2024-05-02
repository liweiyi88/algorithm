package directedgraph

type TopologicalSort struct {
	Order []int
}

func NewTopologicalSort(g Digraph) *TopologicalSort {
	cycle := NewDirectedCycle(g)

	ts := &TopologicalSort{
		Order: make([]int, 0),
	}

	if !cycle.HasCycle() {
		order := NewDepthFirstOrder(g)

		for !order.ReversePost.IsEmpty() {
			v, ok := order.ReversePost.Pop()
			if ok {
				ts.Order = append(ts.Order, v)
			}
		}
	}

	return ts
}

func (t *TopologicalSort) HasOrder() bool {
	return len(t.Order) > 0
}
