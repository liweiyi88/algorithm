package unionfind

type UFQuickFind struct {
	Id    []int
	count int
}

func NewUFQuickFind(n int) *UFQuickFind {
	id := make([]int, 0, n)

	for i := 0; i < n; i++ {
		id = append(id, i)
	}

	return &UFQuickFind{
		count: n,
		Id:    id,
	}
}

func (uf *UFQuickFind) Union(p, q int) {
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

func (uf *UFQuickFind) Find(p int) int {
	return uf.Id[p]
}

func (uf *UFQuickFind) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UFQuickFind) Count() int {
	return uf.count
}
