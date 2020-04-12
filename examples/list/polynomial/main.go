package main

import (
	"fmt"
	"strconv"

	"github.com/fancxxy/algorithm/list/singlylinkedlist"
)

/*

f1(x) = 5x^2 + 4x^1 + 2x^0
f2(x) = 5x^1 + 5x^0
f1(x) + f2(x) = 5x^2 + 9x^1 +7x^0
f1(x) * f2(x) = 25x^3 + 45x^2 + 30x^1 + 10x^0

*/

type polynomial struct {
	Coefficient int
	Exponent    int
}

func (p *polynomial) String() string {
	return strconv.Itoa(p.Coefficient) + "x^" + strconv.Itoa(p.Exponent)
}

func addPolynomial(poly1, poly2 *singlylinkedlist.List) *singlylinkedlist.List {
	var (
		poly3 = singlylinkedlist.New()
		node1 = poly1.First()
		node2 = poly2.First()
		node3 = poly3.Head()
	)

	for node1 != nil && node2 != nil {
		value1, value2 := node1.Value.(*polynomial), node2.Value.(*polynomial)
		if value1.Exponent > value2.Exponent {
			node3 = poly3.Insert(&polynomial{
				Coefficient: value1.Coefficient,
				Exponent:    value1.Exponent,
			}, node3)
			node1 = node1.Next
		} else if value1.Exponent < value2.Exponent {
			node3 = poly3.Insert(&polynomial{
				Coefficient: value2.Coefficient,
				Exponent:    value2.Exponent,
			}, node3)
			node2 = node2.Next
		} else {
			node3 = poly3.Insert(&polynomial{
				Coefficient: value1.Coefficient + value2.Coefficient,
				Exponent:    value1.Exponent,
			}, node3)
			node1 = node1.Next
			node2 = node2.Next
		}
	}

	for node1 != nil {
		value1 := node1.Value.(*polynomial)
		node3 = poly3.Insert(&polynomial{
			Coefficient: value1.Coefficient,
			Exponent:    value1.Exponent,
		}, node3)
		node1 = node1.Next
	}
	for node2 != nil {
		value2 := node2.Value.(*polynomial)
		node3 = poly3.Insert(&polynomial{
			Coefficient: value2.Coefficient,
			Exponent:    value2.Exponent,
		}, node3)
		node2 = node2.Next
	}

	return poly3
}

func multiplyPolynomial(poly1, poly2 *singlylinkedlist.List) *singlylinkedlist.List {
	var (
		poly3 = singlylinkedlist.New()
		node1 = poly1.First()
		node2 = poly2.First()
		node3 = poly3.Head()
	)

	for node1 != nil {
		value1 := node1.Value.(*polynomial)
		node2 = poly2.First()
		for node2 != nil {
			value2 := node2.Value.(*polynomial)
			node3 = poly3.Insert(&polynomial{
				Coefficient: value1.Coefficient * value2.Coefficient,
				Exponent:    value1.Exponent + value2.Exponent,
			}, node3)
			node2 = node2.Next
		}
		node1 = node1.Next
	}

	curr := poly3.Head().Next
	for curr != nil && curr.Next != nil {
		currValue := curr.Value.(*polynomial)
		dup := curr
		for dup.Next != nil {
			dupValue := dup.Next.Value.(*polynomial)
			if currValue.Exponent == dupValue.Exponent {
				currValue.Coefficient += dupValue.Coefficient
				poly3.Remove(dup)
			} else {
				dup = dup.Next
			}
		}
		curr = curr.Next
	}

	return poly3
}

func main() {
	poly1 := singlylinkedlist.New()
	node := poly1.Insert(&polynomial{Coefficient: 5, Exponent: 2}, poly1.Head())
	node = poly1.Insert(&polynomial{Coefficient: 4, Exponent: 1}, node)
	node = poly1.Insert(&polynomial{Coefficient: 2, Exponent: 0}, node)

	poly2 := singlylinkedlist.New()
	node = poly2.Insert(&polynomial{Coefficient: 5, Exponent: 1}, poly2.Head())
	node = poly2.Insert(&polynomial{Coefficient: 5, Exponent: 0}, node)

	for _, value := range addPolynomial(poly1, poly2).Values() {
		fmt.Printf("%v ", value)
	}

	fmt.Println()

	for _, value := range multiplyPolynomial(poly1, poly2).Values() {
		fmt.Printf("%v ", value)
	}
}
