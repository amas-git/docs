package main

import (
	"fmt"
)

func Fib(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}

func FibIter(n int) int {
	switch n {
	case 1:
		return 1
	case 2:
		return 1
	default:
		n1 := 1
		n2 := 2
		for i := 2; i < n; i++ {
			n1, n2 = n2, n2+n1
		}
		return n2
	}
}

func add(n int) func(int) int {
	return func(m int) int {
		return n + m
	}
}

func SumTCO(xs []int) int {
	if len(xs) == 0 {
		return 0
	}

	return xs[0] + SumTCO(xs[1:])
}

type FontSize int

const (
	_              = iota
	SMALL FontSize = iota*4 + 10
	MEDIUM
	LARGE
	XLARGE
)

func main() {
	fmt.Println(SumTCO([]int{1, 2, 3}))
	fmt.Println(add(1)(2))
	fmt.Println(SMALL, MEDIUM, LARGE, XLARGE)
}
