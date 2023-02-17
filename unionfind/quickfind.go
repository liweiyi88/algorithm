package unionfind

type QuickFind struct {
	Id    []int
	count int
}

func NewQuickFind(n int) *QuickFind {
	id := make([]int, 0, n)

	for i := 0; i < n; i++ {
		id = append(id, i)
	}

	return &QuickFind{
		count: n,
		Id:    id,
	}
}

func (uf *QuickFind) Union(p, q int) {
	pId := uf.Find(p)
	qId := uf.Find(q)

	if pId == qId {
		return
	}

	for i := 0; i < len(uf.Id); i++ {
		if uf.Id[i] == pId {
			uf.Id[i] = qId
		}
	}

	uf.count--
}

func (uf *QuickFind) Find(p int) int {
	return uf.Id[p]
}

func (uf *QuickFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *QuickFind) Count() int {
	return uf.count
}
