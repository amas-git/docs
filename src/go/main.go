package main

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func main() {

	for n := range seq(100) {
		fmt.Println(n)
	}
}
