package undirectedgraph

import (
	"github.com/liweiyi88/algorithm/collections"
	"github.com/liweiyi88/algorithm/undirectedgraph"
)

type BreadthFirstSearch struct {
	marked []bool
	count  int
}

func NewBreadthFirstSearch(g undirectedgraph.Graph, s int) *BreadthFirstSearch {
	bfs := &BreadthFirstSearch{
		marked: make([]bool, g.V),
		count:  0,
	}

	bfs.bfs(g, s)

	return bfs
}

func (b *BreadthFirstSearch) bfs(g undirectedgraph.Graph, s int) {
	queue := collections.NewQueue[int]()
	b.marked[s] = true
	queue.Enqueue(s)

	for {
		v, ok := queue.Dequeue()
		if !ok {
			break
		}

		for w := range g.Adj[v] {
			if !b.marked[w] {
				b.marked[w] = true
				queue.Enqueue(w)
			}
		}
	}
}

func (b *BreadthFirstSearch) Marked(v int) bool {
	return b.marked[v]
}

func (b *BreadthFirstSearch) Count() int {
	return b.count
}
