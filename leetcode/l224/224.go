package l224

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	closed   bool
	negative bool
	numbers  []string
	next     *node
}

func (n *node) calculate() int {
	results := 0
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
	first *node
}

func (c *Calculator) Calculate() int {
	result := 0

	for first := c.first; first != nil; first = first.next {
		result = first.calculate()

		if first.next != nil {
			first.next.numbers = append(first.next.numbers, strconv.Itoa(result))
		}

		fmt.Println(first.closed, first.negative, first.numbers, result)

	}

	return result
}

func (c *Calculator) Add(numbers []string, negative bool, closed bool) {
	c.first = &node{
		closed:   closed,
		negative: negative,
		numbers:  numbers,
		next:     c.first,
	}
}

func (c *Calculator) FirstUnclosedNode() *node {
	for first := c.first; first != nil; first = first.next {
		if !first.closed {
			return first
		}
	}

	return nil
}

func calculate(s string) int {
	calculator := &Calculator{}

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
			unclosedNode := calculator.FirstUnclosedNode()
			// if unclosedNode.next != nil {
			// 	total := unclosedNode.calculate()
			// 	if unclosedNode.negative {
			// 		number := strconv.Itoa(-total)
			// 		unclosedNode.negative = false
			// 		unclosedNode.numbers = []string{number}
			// 	}
			// }

			unclosedNode.closed = true

			continue
		}

		if index > 0 && isNumber(digit) && isNumber(string(s[index-1])) {
			firstUnclosedNode := calculator.FirstUnclosedNode()
			firstNumbers := firstUnclosedNode.numbers
			lastNumber := firstNumbers[len(firstNumbers)-1]
			firstNumbers[len(firstNumbers)-1] = lastNumber + digit
			continue
		}

		if isNumber(digit) && index == 0 {
			calculator.Add([]string{digit}, false, false)
			continue
		}

		if index > 0 && isNumber(digit) && string(s[index-1]) == "-" {
			firstUnclosedNode := calculator.FirstUnclosedNode()

			if firstUnclosedNode == nil {
				calculator.Add([]string{"-" + digit}, false, false)
			} else {
				firstUnclosedNode.numbers = append(firstUnclosedNode.numbers, "-"+digit)
			}

			continue
		}

		if isNumber(digit) {
			if calculator.first == nil {
				calculator.Add([]string{digit}, false, false)
			} else {
				firstUnclosedNode := calculator.FirstUnclosedNode()

				if firstUnclosedNode != nil {
					firstUnclosedNode.numbers = append(firstUnclosedNode.numbers, digit)
				} else {
					calculator.first.numbers = append(calculator.first.numbers, digit)
				}
			}
		}
	}

	// fmt.Printf("%+v", calculator.first.calculate())
	// fmt.Println()
	// fmt.Printf("%+v", calculator.first.next.calculate())
	// fmt.Println()
	// fmt.Printf("%+v", calculator.first.next.next.calculate())

	return calculator.Calculate()
}

func isNumber(s string) bool {
	if _, err := strconv.Atoi(s); err == nil {
		return true
	}

	return false
}
