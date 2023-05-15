package l224

import (
	"testing"
)

func TestCalculate10(t *testing.T) {
	expect := 2
	actual := calculate("1-(-2-(-2+1))")

	if expect != actual {
		t.Errorf("expect %d, actual got: %d", expect, actual)
	}
}

func TestCalculate9(t *testing.T) {
	expect := 11
	actual := calculate("(7)-(0)+(4)")

	if expect != actual {
		t.Errorf("expect %d, actual got: %d", expect, actual)
	}
}

func TestCalculate(t *testing.T) {
	expect := 2
	actual := calculate("1 + 1")

	if expect != actual {
		t.Errorf("expect %d, actual got: %d", expect, actual)
	}

	expect = 3
	actual = calculate(" 2-1 + 2 ")

	if expect != actual {
		t.Errorf("expect %d, actual got: %d", expect, actual)
	}

	expect = 23
	actual = calculate("(1+(4+5+2)-3)+(6+8)")

	if expect != actual {
		t.Errorf("expect %d, actual got: %d", expect, actual)
	}

	expect = 2147483647
	actual = calculate("2147483647")

	if expect != actual {
		t.Errorf("expect %d, actual got: %d", expect, actual)
	}

	expect = -1
	actual = calculate("-2+ 1")

	if expect != actual {
		t.Errorf("expect %d, actual got: %d", expect, actual)
	}

	expect = -2
	actual = calculate("(6)-(8)-(7)+(1+(6))")

	if expect != actual {
		t.Errorf("expect %d, actual got: %d", expect, actual)
	}
}

func TestCalculate2(t *testing.T) {
	expect := -12
	actual := calculate("- (3 + (4 + 5))")

	if expect != actual {
		t.Errorf("expect %d, actual got: %d", expect, actual)
	}
}

func TestCalculate3(t *testing.T) {
	expect := -15
	actual := calculate("2-4-(8+2-6+(8+4-(1)+8-10))")

	if expect != actual {
		t.Errorf("expect %d, actual got: %d", expect, actual)
	}
}

func TestCalculate4(t *testing.T) {
	expect := 23
	actual := calculate("(3-(5-(8)-(13)-4))")

	if expect != actual {
		t.Errorf("expect %d, actual got: %d", expect, actual)
	}
}
