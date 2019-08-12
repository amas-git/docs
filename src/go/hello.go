package main

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
	"unsafe"
)

const PI = 3.14

func init() {
	fmt.Println("PI=", PI)
	fmt.Println("GOMAXPROCS=", runtime.GOMAXPROCS(-1))
}

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
	//chanTest()
	testWaitGroup()
	testOnce()
	simapleNumberStream(5)
	//deadLock()
	constTest()
	nouse()
	overflow()
	testDefer()
	fmt.Println("The result is", testDeferChangeResultVar())
	testSizeof()
	loopTrick()
	reflectTest()
	cat("/etc/passwd")
	scanTest("/etc/passwd")
	testBuffer()
	testGob()
	testGoto()
	fmt.Println(sum(1, 2, 5, 6, 7, 8, 99))
	testPointer()
	fmt.Println(OK, CREATED, ACCEPTED, NOAUTH, MC)
	testPanic()
	testStructCompose()
	testInterface()
}

func testInterface() {
	c1 := C1{"c1"}
	// c2 := C2{"c2"}

	// TODO:没搞定
	c1.Hello()

	TYPE_OF(c1)
	TYPE_OF(12)
}

func TYPE_OF(elem interface{}) {
	switch value := elem.(type) {
	default:
		fmt.Println(elem, "IS", value)
	}
}

type Hello interface {
	Hello()
}

type C1 struct {
	name string
}

func (c1 *C1) Hello() {
	fmt.Println("HELLO")
}

func (c1 *C1) Print() {
	fmt.Println("C2", "name=", c1.name)
}

func (c1 *C1) C1Print() {
	fmt.Println("C1P", "name=", c1.name)
}

type C2 struct {
	name string
}

func (c2 *C2) Print() {
	fmt.Println("C2", "name=", c2.name)
}

func (c2 *C2) C2Print() {
	fmt.Println("C2P", "name=", c2.name)
}

type P struct {
	name string
	C1
	C2
}

func (p *P) Print() {
	fmt.Println("P2", "name=", p.name)
}

func testStructCompose() {
	p := P{
		"p",
		C1{"c1"},
		C2{"c2"},
	}

	fmt.Println("struct compose: ", p)
	fmt.Println("struct compose: name=", p.name)
	p.C1Print()
	p.C2Print()
	p.C1.Print()
	p.C2.Print()
	p.Print()
}

func testPanic() {
	f := func(n int) {
		//panic("I'm panic")
		xs := [3]int{}
		xs[n] = 1
	}
	defer func() {
		fmt.Println("defer will still run after panic")
	}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("catch panic: ", err)
		}
	}()
	f(3)
}

const (
	OK       = 200       // iota = 0
	CREATED  = OK + iota // iota = 1
	ACCEPTED             // iota = 2
	NOAUTH               // iota = 3

	REDIRECT = 300             // iota = 4
	MC       = REDIRECT + iota // iota = 5
)

func testPointer() {
	addone := func(n *int) {
		(*n)++
	}
	i := 1
	addone(&i)
	addone(&i)
	addone(&i)
	addone(&i)
	fmt.Println("addone", i)
}

func testGoto() {
	i := 0
LOOP:
	i++
	fmt.Print(i)
	if i < 10 {
		goto LOOP
	}
}

func sum(args ...int) (r int) {
	for _, i := range args {
		r += i
	}
	return
}

func testGob() {
	type Phone struct {
		Tag    string `json:"tag"`
		Number string `json:"number"`
	}

	type Contact struct {
		Name   string  `json:"name"`
		Phones []Phone `json:"phones"`
	}

	contacts := []Contact{
		Contact{
			Name: "amas",
			Phones: []Phone{
				Phone{
					Tag:    "HOME",
					Number: "18876541230",
				},
			},
		},
		Contact{
			Name: "doudou",
			Phones: []Phone{
				{
					Tag:    "SCHOOL",
					Number: "04714565312",
				},
				{
					Tag:    "OFFICE",
					Number: "04714565312",
				},
			},
		},
	}

	file, err := os.Create("/tmp/contacts.dat")
	if err != nil {
		return
	}

	encoder := gob.NewEncoder(file)
	if err := encoder.Encode(contacts); err != nil {
		return
	}

	json := json.NewEncoder(os.Stdout)
	json.Encode(contacts)
}

func testBuffer() {
	buffer := bytes.Buffer{}
	buffer.WriteString("hello")
	buffer.WriteString(" ")
	buffer.WriteString("world")
	buffer.WriteTo(os.Stdout)
}

func scanTest(path string) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		xs := strings.Split(scanner.Text(), ":")
		fmt.Fprintf(os.Stdout, "USER: %v\n", xs[0])
	}
}

func cat(path string) {
	if file, ok := os.OpenFile(path, os.O_RDONLY, 0666); ok == nil {
		defer file.Close()
		n, err := io.Copy(os.Stdout, file)
		if err != nil {
			fmt.Printf("READ %d byte, ERROR : %v\n", n, err)
		}
	} else {
		fmt.Printf("File Not Existed : %v\n", path)
	}
}

func reflectTest() {
	type Person struct {
		Name string `default:amas`
		Age  int    `default:1`
	}

	amas := &Person{
		Name: "amas",
		Age:  199,
	}
	field, ok := reflect.TypeOf(amas).Elem().FieldByName("Name")
	if ok {
		fmt.Println(string(field.Tag))
	}
}

func loopTrick() {
	for i := range [5][]int{} {
		fmt.Printf("> %d\n", i)
	}
}

type Age int

func testSizeof() {
	fmt.Println("int", unsafe.Sizeof(int(1)))           // 8
	fmt.Println("rune", unsafe.Sizeof(rune(1)))         // 4
	fmt.Println("rune", unsafe.Sizeof(int8(1)))         // 1
	fmt.Println("byte", unsafe.Sizeof(byte(1)))         // 1
	fmt.Println("map", unsafe.Sizeof(map[string]int{})) // 8
}

func testDeferChangeResultVar() (n int) {
	defer func() {
		n = 111
	}()
	return 1
}

func testDefer() {
	defer fmt.Println("DEFER 2")
	defer fmt.Println("DEFER 1")
}

func overflow() {
	// a := 1 << 64 // ERRORS
	print("buildin")
}

const HELLO_WORLD = "hello world"

func init() {
	fmt.Println(HELLO_WORLD)
}

func nouse() {
	a := 1
	b := 2
	_, _ = a, b // fixed not ussed issue
}

func constTest() {
	const (
		A int = 1
		B
		C

		D int = iota
		E int = iota
		F int = iota
		G int = 199
		H int = iota
	)

	const (
		N0 int = iota
		N1
		N2
		N3
		N4
		N5
	)

	fmt.Println("constTest/1", A, B, C)
	fmt.Println("constTest/2", D, E, F, G, H)
	fmt.Println("constTest/3", N0, N1, N2, N3, N4, N5)
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
