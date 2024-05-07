package validparentheses

type Stack[T any] []T

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(item T) {
	*s = append(*s, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var v T
		return v, false
	}

	item := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]

	return item, true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Size() int {
	return len(*s)
}

func isValid(s string) bool {
	stack := NewStack[rune]()

	for _, char := range s {
		if char == '(' || char == '[' || char == '{' {
			stack.Push(char)
		} else {
			c, ok := stack.Pop()
			if !ok {
				return false
			}

			if c == '(' && char != ')' {
				return false
			}

			if c == '[' && char != ']' {
				return false
			}

			if c == '{' && char != '}' {
				return false
			}
		}
	}

	return stack.IsEmpty()
}
