package l224

import (
	"strconv"
	"strings"
)

type node struct {
	closed   bool
	negative bool
	numbers  []string
	nodes    []*node
}

func newNode(numbers []string, negative, cloased bool) *node {
	return &node{
		closed:   cloased,
		negative: negative,
		numbers:  numbers,
	}
}

func (n *node) nextUnclosedNode() *node {
	if len(n.nodes) > 0 {
		for i := len(n.nodes) - 1; i >= 0; i-- {
			if !n.nodes[i].closed {
				return n.nodes[i].nextUnclosedNode()
			}
		}
	}

	if !n.closed {
		return n
	}

	return nil
}

func (n *node) append(node *node) {
	n.nodes = append(n.nodes, node)
}

func (n *node) calculate() int {
	results := 0

	if len(n.nodes) > 0 {
		for _, n := range n.nodes {
			results += n.calculate()
		}
	}

	for _, number := range n.numbers {
		num, _ := strconv.Atoi(number)
		results += num
	}

	if n.negative {
		return -results
	}

	return results
}

type Calculator struct {
	nodes []*node
}

func (c *Calculator) Calculate() int {
	result := 0

	for _, node := range c.nodes {
		result += node.calculate()
	}

	return result
}

func (c *Calculator) Add(numbers []string, negative bool, closed bool) {
	nextUnclosedNode := c.nextUnclosedNode()

	if nextUnclosedNode == nil {
		c.nodes = append(c.nodes, &node{
			closed:   closed,
			negative: negative,
			numbers:  numbers,
		})
	} else {
		nextUnclosedNode.append(newNode(numbers, negative, closed))
	}
}

func (c *Calculator) nextUnclosedNode() *node {
	for i := len(c.nodes) - 1; i >= 0; i-- {
		if !c.nodes[i].closed {
			return c.nodes[i].nextUnclosedNode()
		}
	}

	return nil
}

func calculate(s string) int {
	calculator := &Calculator{
		nodes: make([]*node, 0),
	}

	s = strings.ReplaceAll(s, " ", "")

	for index, char := range s {
		digit := string(char)

		if digit == "+" {
			continue
		}

		if digit == "(" {
			if index > 0 && string(s[index-1]) == "-" {
				calculator.Add([]string{}, true, false)
			} else {
				calculator.Add([]string{}, false, false)
			}

			continue
		}

		if digit == ")" {
			unclosedNode := calculator.nextUnclosedNode()
			unclosedNode.closed = true
			continue
		}

		if index > 0 && isNumber(digit) && isNumber(string(s[index-1])) {
			nextUnclosedNode := calculator.nextUnclosedNode()
			firstNumbers := nextUnclosedNode.numbers
			lastNumber := firstNumbers[len(firstNumbers)-1]
			firstNumbers[len(firstNumbers)-1] = lastNumber + digit
			continue
		}

		if isNumber(digit) && index == 0 {
			calculator.Add([]string{digit}, false, false)
			continue
		}

		if index > 0 && isNumber(digit) && string(s[index-1]) == "-" {
			nextUnclosedNode := calculator.nextUnclosedNode()

			if nextUnclosedNode == nil {
				calculator.Add([]string{"-" + digit}, false, false)
			} else {
				nextUnclosedNode.numbers = append(nextUnclosedNode.numbers, "-"+digit)
			}

			continue
		}

		if isNumber(digit) {
			nextUnclosedNode := calculator.nextUnclosedNode()

			if nextUnclosedNode == nil {
				calculator.Add([]string{digit}, false, true)
			} else {
				nextUnclosedNode.numbers = append(nextUnclosedNode.numbers, digit)
			}
		}
	}

	return calculator.Calculate()
}

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}
