package main

import (
	"github.com/liweiyi88/algorithm/unionfind"
)

func main() {
	unionfindDemo := unionfind.NewDemo(10, unionfind.QuickFind)
	unionfindDemo.Run()
}
