package main

import (
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"sort"
	"sync"
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
	//randTest()
	funcTest()
	//letsgo()
	chanTest()
	//panic("I'm dead")
	testWaitGroup()
	testOnce()
	simapleNumberStream(5)
	//deadLock()
}

func deadLock() {
	c := make(chan int)

	fmt.Println(<-c) // dead lock
}

func simapleNumberStream(max int) {
	c := make(chan int)

	go func() {
		defer close(c) // 确保channel使用完毕之后关掉
		for i := 0; i < max; i++ {
			c <- i
			time.Sleep(time.Second)
		}
	}()

	go func() {
		for x := range c {
			fmt.Printf(" - %d\n", x)
		}
	}()

	for x := range c {
		fmt.Printf(" * %d\n", x)
	}
}

func testOnce() {
	once := sync.Once{}
	sum := 0
	inc := func() {
		sum++
	}

	go (func() {
		once.Do(inc)
	})()
	once.Do(inc)
	once.Do(inc)
	once.Do(inc)
	once.Do(inc)

	fmt.Println("sum: ", sum)
}

func testWaitGroup() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(s string) {
			defer wg.Done()
			fmt.Println(s)
		}(salutation)
	}
	wg.Wait()
}

func randTest() {
	for i := 0; i < 100; i++ {
		fmt.Println(randn(1, 200))
	}
}

func randn(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min)
}

func chanTest() {
	//timeout = time.After(5)

	c := make(chan string)

	fn := func(id int, c chan string) {
		n := randn(200, 2000)
		time.Sleep(time.Millisecond * time.Duration(n))
		fmt.Printf("TASK %d DONE (%d)\n", id, n)
		c <- "[" + string(id) + "]"
	}

	works := []int{1, 2, 3, 4, 5}
	for _, v := range works {
		go fn(v, c)
	}

	timeout := time.After(time.Second * 1)
	for _ = range works {
		select {
		case r := <-c:
			fmt.Printf("FINISHED: %s \n", r)
		case <-timeout:
			fmt.Printf("TIMOUT EXIT")
			return
		}
	}

}

func letsgo() {
	fn := func() {
		for n := 0; n < 10; n++ {
			fmt.Println(time.Now().Format(time.RFC850))
			time.Sleep(time.Millisecond * 500)
		}
	}
	go fn()
	fmt.Println("letsgo over")
	time.Sleep(time.Second * 8)
}

func funcTest() {
	fn := func(x int) {
		fmt.Printf("=== [%v] ===\n", x)
	}

	for _, v := range map[string]int{"a": 1, "b": 2} {
		fn(v)
	}

	sum := func(a int, b int) int {
		return a + b
	}(1, 2)
	fmt.Println(sum)
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p person) print() {
	fmt.Printf("[NAME:%8s | AGE: %2d]\n", p.Name, p.Age)
}

func structTest() {
	var point struct {
		x int
		y int
	}

	point.x = 1
	point.y = 2

	fmt.Printf("%v\n", point)

	type box struct {
		width  int
		height int
	}
	var a box
	a.width = 1
	a.height = 2

	b := box{11, 12}

	c := box{height: 1}

	fmt.Println(a, b, c)

	students := []person{
		{"zhou", 11},
		{"bob", 12},
		{"amas", 13},
	}
	bytes, err := json.Marshal(students)
	fmt.Println(bytes, err)
	fmt.Println(string(bytes)) // FIXME: NOT WORK???

	amas := person{"amas", 110}
	amas.print()

	double := func(p person) {
		p.Age *= 2
		p.print()
	}

	double(amas)
	amas.print()

	func(p *person) {
		p.Age *= 3
		p.print()
	}(&amas)
	amas.print()

	ab := AB{A{"amas"}, B{19}}
	fmt.Println(ab)
	ab.FA()
	ab.A.FA()
	ab.B.FB()
	ab.FB()

	var one Say = A{"one"}
	iSay, ok := one.(Say)
	if ok {
		fmt.Println(iSay.say())
	}
}

type Say interface {
	say() string
}

type AB struct {
	A
	B
}

type A struct {
	name string
}

type B struct {
	age int
}

func (a A) say() string {
	return "hahaha"
}

func (a A) FA() {
	fmt.Printf("FA name %v\n", a.name)
}
func (b B) FB() {
	fmt.Printf("FB age %v\n", b.age)
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
