package search

import (
	"bytes"
	"testing"
)

func TestPrint(t *testing.T) {
	btree := BinaryTree{}
	btree.Put("S", 1)
	btree.Put("E", 2)
	btree.Put("X", 3)
	btree.Put("A", 4)
	btree.Put("R", 5)
	btree.Put("C", 6)
	btree.Put("H", 7)
	btree.Put("M", 8)

	var b bytes.Buffer
	btree.Print(&b)

	expect := "ACEHMRSX"
	if b.String() != expect {
		t.Errorf("expect: %s, but got: %s", expect, b.String())
	}
}

func TestGet(t *testing.T) {
	btree := BinaryTree{}
	btree.Put("S", 1)
	btree.Put("E", 2)
	btree.Put("X", 3)
	btree.Put("A", 4)
	btree.Put("R", 5)
	btree.Put("C", 6)
	btree.Put("H", 7)
	btree.Put("M", 8)

	expect := 5
	actual := btree.Get("R")

	if actual != expect {
		t.Errorf("expect: %d, get: %d", expect, actual)
	}
}

func TestDelete(t *testing.T) {
	btree := BinaryTree{}
	btree.Put("S", 1)
	btree.Put("E", 2)
	btree.Put("X", 3)
	btree.Put("A", 4)
	btree.Put("R", 5)
	btree.Put("C", 6)
	btree.Put("H", 7)
	btree.Put("M", 8)

	btree.Delete("E")
	var b bytes.Buffer
	btree.Print(&b)

	expect := "ACHMRSX"
	if b.String() != expect {
		t.Errorf("expect: %s, but got: %s", expect, b.String())
	}
}
