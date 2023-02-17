package unionfind

type QuickUnion struct {
	Id    []int
	count int
}

func NewQuickUnion(n int) *QuickUnion {
	id := make([]int, 0, n)

	for i := 0; i < n; i++ {
		id = append(id, i)
	}

	return &QuickUnion{
		count: n,
		Id:    id,
	}
}

func (uf *QuickUnion) Union(p, q int) {
	pId := uf.Find(p)
	qId := uf.Find(q)

	if pId == qId {
		return
	}

	uf.Id[pId] = qId

	uf.count--
}

// Find the root node's value, root nodes'value is the same as its position(index).
func (uf *QuickUnion) Find(p int) int {
	for {
		if p != uf.Id[p] {
			p = uf.Id[p]
			continue
		}

		return p
	}
}

func (uf *QuickUnion) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *QuickUnion) Count() int {
	return uf.count
}
