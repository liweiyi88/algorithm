package collections

import "testing"

func TestBag(t *testing.T) {
	bag := NewBag[int]()

	bag.Add(2)
	bag.Add(2)
	bag.Add(3)

	if bag.Size() != 2 {
		t.Error("unexpected bag size")
	}

	bag.Delete(1)

	if bag.Size() != 2 {
		t.Error("unexpected bag size")
	}

	bag.Delete(2)

	if bag.Size() != 1 {
		t.Error("unexpected bag size")
	}

	bag.Delete(3)

	if !bag.IsEmpty() {
		t.Error("unexpected bag size")
	}

	if bag.Count(1) != 0 {
		t.Error("unexpected count")
	}

	bag.Add(1)
	bag.Add(1)
	bag.Add(1)

	if bag.Count(1) != 3 {
		t.Error("unexpected count")
	}

	for k, v := range bag {
		t.Log(k, v)
	}
}
