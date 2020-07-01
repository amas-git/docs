package main

import (
	"fmt"
	"strings"
)

func f(x int) int {
	return 2 * x
}

func g(x int) int {
	return 1 + x
}

func id(x interface{}) interface{} {
	return x
}

func a(s string) (string, string) {
	return "a()", fmt.Sprintf("a:%v:a", s)
}

func b(s string) (string, string) {
	return "b()", fmt.Sprintf("b:%v:b", s)
}

func toUpper(s string) (string, string) {
	return "toUpper()", strings.ToUpper(s)
}

func compose(f1 func(string) (string, string), f2 func(string) (string, string), s string) (string, string) {
	log1, r1 := f1(s)
	log2, r2 := f2(r1)
	return fmt.Sprintf("%v\n%v", log1, log2), r2
}

func return0(s string) (string, string) {
	return "", s
}

func print(x interface{}) {
	fmt.Println(x)
}

type Nothing int

type Maybe struct {
}

type Optional struct {
}

func main() {
	print(f(g(1)))
	print(id(1) == 1)
	print(id("cool") == "cool")

	log, r := compose(a, b, "hello")
	print(r)
	print(log)

}
