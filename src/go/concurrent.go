package main

import (
	"fmt"
	"sync"
	"time"
)

func coprint(xs []string) {
	var wg sync.WaitGroup
	for _, x := range xs {
		wg.Add(1)
		go func(x string) {
			defer wg.Done()
			fmt.Println(x)
		}(x)
	}
	wg.Wait()
}

func cond(fn func() bool) {
	c := sync.NewCond(&sync.Mutex{})
	c.L.Lock()
	defer c.L.Unlock()

	for fn() == false {
		c.Wait() // Wait()中会调用L.Unlock(), 将锁释放
	}
}

func once_play() {
	TITLE("once_play")
	n := 0
	inc := func() {
		n++
	}

	var wg sync.WaitGroup
	var once sync.Once

	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			once.Do(inc)
		}()
	}
	wg.Wait()
	fmt.Println("n:", n)
}

func TITLE(msg string) {
	fmt.Println("====================", "[", msg, "]")
}

func cond_play() {
	TITLE("cond_play")

	c := sync.NewCond(&sync.Mutex{})
	q := make([]interface{}, 0, 10)

	pop := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		if len(q) > 1 {
			q = q[1:]
		}
		fmt.Println("> POP")
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(q) == 2 {
			c.Wait()
		}
		fmt.Println("> ADD", len(q))
		q = append(q, struct{}{})

		go pop(1 * time.Second)
		go pop(1 * time.Second)
		c.L.Unlock()
	}
}

func add(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("Hello")

	coprint([]string{"a", "b", "c", "d", "e"})

	cond(func() bool {
		return true
	})

	//cond_play()
	//once_play()

}
