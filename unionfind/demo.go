package unionfind

import "fmt"

type UnionFind interface {
	Union(p, q int)
	Find(p int) int
	Connected(p, q int) bool
	Count() int
}

type Demo struct {
	NumberOfComponentIds int
	Method               Method
}

type Method int64

const (
	QuickFind Method = iota
	QuickUnion
	Weight
)

func NewDemo(n int, method Method) *Demo {
	return &Demo{
		NumberOfComponentIds: n,
		Method:               method,
	}
}

func (d *Demo) algorithm() UnionFind {
	switch d.Method {
	case QuickFind:
		return NewUFQuickFind(d.NumberOfComponentIds)
	default:
		return NewUFQuickFind(d.NumberOfComponentIds)
	}
}

func (d *Demo) Run() {
	data := make([][2]int, 0, 11)
	data = append(data, [2]int{4, 3})
	data = append(data, [2]int{3, 8})
	data = append(data, [2]int{6, 5})
	data = append(data, [2]int{9, 4})
	data = append(data, [2]int{2, 1})
	data = append(data, [2]int{8, 9})
	data = append(data, [2]int{5, 0})
	data = append(data, [2]int{7, 2})
	data = append(data, [2]int{6, 1})
	data = append(data, [2]int{1, 0})
	data = append(data, [2]int{6, 7})

	al := d.algorithm()

	for _, v := range data {
		if al.Connected(v[0], v[1]) {
			fmt.Printf("%+v\n", al)
			continue
		}
		al.Union(v[0], v[1])
		fmt.Printf("%+v\n", al)
	}
}
