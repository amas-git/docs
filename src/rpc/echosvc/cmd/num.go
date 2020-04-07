package main

import (
	"fmt"

	"gonum.org/v1/gonum/mat"
)

func main() {
	m := mat.NewVecDense(5, []float64{1, 2, 3, 4, 5})

	formated := mat.Formatted(m, mat.Prefix(""), mat.Squeeze())
	fmt.Println(formated)
}
