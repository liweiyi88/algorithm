package unionfind

type WeightedQuickUnion struct {
	Id       []int
	rootSize []int
	count    int
}

func NewWeightedQuickUnion(n int) *WeightedQuickUnion {
	id, rootSize := make([]int, 0, n), make([]int, 0, n)

	for i := 0; i < n; i++ {
		id = append(id, i)
		rootSize = append(rootSize, 1)
	}

	return &WeightedQuickUnion{
		count:    n,
		Id:       id,
		rootSize: rootSize,
	}
}

func (uf *WeightedQuickUnion) Union(p, q int) {
	pId := uf.Find(p)
	qId := uf.Find(q)

	if pId == qId {
		return
	}

	if uf.rootSize[pId] < uf.rootSize[qId] {
		uf.Id[pId] = qId
		uf.rootSize[qId] += uf.rootSize[pId]
	} else {
		uf.Id[qId] = pId
		uf.rootSize[pId] += uf.rootSize[qId]
	}

	uf.count--
}

// Find the root node's value, root nodes'value is the same as its position(index).
func (uf *WeightedQuickUnion) Find(p int) int {
	for {
		if p != uf.Id[p] {
			p = uf.Id[p]
			continue
		}

		return p
	}
}

func (uf *WeightedQuickUnion) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *WeightedQuickUnion) Count() int {
	return uf.count
}
