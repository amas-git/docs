package main

import (
	"fmt"
	"math/big"
	"sort"
	"time"
)

const PI = 3.14

func main() {
	fmt.Println("let's go", add(100, 2009))
	fmt.Println("2 is odd", isOdd(2))
	forever()
	//simpleLoop(10)
	fmt.Println(weekday(1, 0))
	floatNum()
	printType()
	bigInt()
	stringTest()
	typeCovert()
	defaultType()
	array()
	sortTest()
	put("aaa")
	put("hello", "word")
	mapTest()
	structTest()
}

func structTest() {
	var point struct {
		x int
		y int
	}

	point.x = 1
	point.y = 2

	fmt.Printf("%v\n", point)
}

func mapTest() {
	people := map[string]int{
		"amas": 1,
		"boo":  2,
		"car":  3,
		"dog":  4,
		"exc":  5,
	}

	fmt.Printf("%v\n", people)

	for k, v := range people {
		fmt.Printf("%v : %v\n", k, v)
	}

	if value, ok := people["dog"]; ok {
		fmt.Printf("dog is esisted %d\nm", value)
	} else {
		fmt.Printf("dog is NOT esisted\nm")
	}

	// AS SET
	set := map[string]bool{}
	set["a"] = true
	set["b"] = true
	set["c"] = true

	if set["a"] {
		fmt.Println("HAS 'a'")
	}
}

func sortTest() {
	name := []string{
		"bob",
		"amas",
		"join",
		"kate",
		"blade",
	}

	sort.StringSlice(name).Sort()
	fmt.Println(name)
}

func put(words ...string) {
	for i, x := range words {
		fmt.Printf("%2d %s\n", i, x)
	}
}

func dump(label string, slice []string) {
	fmt.Printf("%v: length %v, capacity %v %v\n", label, len(slice), cap(slice), slice)
}

func array() {
	var xs [8]int
	xs[0] = 1
	ys := []int{1, 2, 3, 4, 7, 8, 9, 9}
	ys = append(ys, 1998)

	fmt.Println(ys)
	fmt.Println(ys[0:3])
	fmt.Println(ys[0 : 2*2])
	//xs[9] = 2 // out of bound

	board := [][]int{
		{1, 2, 3},
		{1, 2, 3},
		{1, 2, 3},
	}

	fmt.Println(board)

	for _, x := range board {
		for _, y := range x {
			fmt.Printf("%2v\n", y)
		}
	}

	zs := []string{"a", "bb", "ccc", "ddd", "EEE"}
	dump("zs", zs)
	dump("zs[0:2]", zs[0:2])
	dump("zs[0:2]", zs[0:2:2])
	dump("make([]int, 0,10)", make([]string, 0, 100))
}

func defaultType() {
	typeOf(1)
	typeOf(1.0)
	typeOf("hello")
	typeOf(isOdd)
	typeOf([]int{1, 2, 3})
	typeOf([3]int{1, 2, 3})
	typeOf([...]int{1, 2, 3})
}

func isOdd(n int) bool {
	var r = n % 2
	if r != 0 {
		return true
	} else {
		return false
	}
}

// 求余数，谁用谁知道
func DivMod(dvdn, dvsr int) (q, r int) {
	r = dvdn
	for r >= dvsr {
		q += 1
		r = r - dvsr
	}
	return
}

func simpleLoop(n int) {
	for n > 0 {
		fmt.Println(n)
		n--
		time.Sleep(time.Second * 5)
	}
}

func stringTest() {
	fmt.Println(len("123"))
	str := "hello world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%c\n", str[i])
	}

	for i, c := range str {
		fmt.Printf("%0d : %c\n", i, c)
	}

	for i := range []int{1, 2, 3, 4, 5} {
		fmt.Println(i)
	}
}

func typeCovert() {
	f := 1111999990.99999
	n := int16(f)

	fmt.Printf("%d\n", n)
}

func typeOf(i interface{}) {
	fmt.Printf("%8v :: %T\n", i, i)
}

func floatNum() {
	// pi := 3.14
	f1 := 0.1
	f2 := f1 + 0.2
	fmt.Println(f2) // ieee 754

	third := 1 / 3.0
	fmt.Println(third + third + third)
}

func bigInt() {
	n := big.NewInt(989898989898)
	n.SetString("123333333333333333333333333333333333333333333333333333333333", 10)
	r := new(big.Int)
	r.Mul(n, big.NewInt(1000))
	fmt.Println(n.String())
	fmt.Println(r.String())
}

func printType() {
	r, g, b := 1, 2, 3
	n := 100
	f := 1.27
	fmt.Printf("%v is %T\n", n, n)
	fmt.Printf("%v is %T\n", f, f)
	fmt.Printf("RGB(%d,%d,%d)\n", r, g, b)
}

func weekday(n int, offset int) string {
	r := ""
	switch n += offset; n {
	case 1:
		r = "星期一"
	case 2:
		r = "星期二"
	case 3:
		r = "星期三"
	default:
		r = "NA"
	}
	return r
}

func forever() {
	var n = 100
	for {
		fmt.Println("HELLO", n)
		n--
		if n < 0 {
			break
		}
	}
}

//-----------------------------------------------------[ func ]
func add(a int, b int) int {
	return a + b
}

func pair(a int, b int) (int, int) {
	return a, b
}
