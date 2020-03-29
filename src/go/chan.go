package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func times(fun interface{}, times int) {
	for i := 0; i < times; i++ {
		fun.(func())()
	}
}

func deadlock_01() {
	ch := make(chan interface{})
	go func() {
		if 0 < 1 {
			return
		}
		ch <- "hello"
	}()
	fmt.Println(<-ch)
}

func seq(max int) (r <-chan interface{}) {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for i := 0; i < max; i++ {
			ch <- i
		}
	}()
	return ch
}

func block(max int) {
	ch := make(chan interface{})
	var wg sync.WaitGroup

	for i := 0; i < max; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-ch
			fmt.Printf("[START]%v/%v\n", i, max)
		}(i)
	}
	fmt.Println("UNLOCK")
	close(ch)
	wg.Wait()
}

func close_nil_chan() {
	//ch := make(chan interface{})
	var ch chan interface{}
	close(ch)
}

func write_closed() {
	ch := make(chan interface{})
	close(ch)
	ch <- 1
}

func close_closed_ch() {
	ch := make(chan interface{})
	close(ch)
	close(ch)
}

func unclose_chan() (r chan int) {
	r = make(chan int)
	go func() {
		for {
			select {
			case <-r:
				return
			default:
				{
					fmt.Printf("...\n")
				}
			}
		}
	}()
	return
}

func wait(sec time.Duration) {
	start := time.Now()
	ch := make(chan interface{})
	go func() {
		time.Sleep(sec * time.Second)
		close(ch)
	}()

	select {
	case <-ch:
		fmt.Println("wait: ", time.Since(start))
	}
}

func wait2(sec time.Duration) {
	start := time.Now()
	select {
	case <-time.After(sec * time.Second):
		fmt.Println("wait2: ", time.Since(start))
	}
}

func wait3(sec time.Duration) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		time.Sleep(sec * time.Second)
		fmt.Println("wait3 done")
	}()
	return ch
}

func select_closed_chan() {
	ch1 := make(chan interface{})
	close(ch1)

	n1 := 0
	for i := 0; i < 100; i++ {
		select {
		case <-ch1:
			n1++
		}
	}
	fmt.Println("n1=", n1)
}

func select_default() {
	var ch chan int
	select {
	case <-ch:
	default:
		fmt.Printf("select: no chans ready")
	}
}

func select_closed() {
	ch_closed := make(chan int)
	close(ch_closed)

	select {
	case v, r := <-ch_closed:
		fmt.Printf("select: closed, value=%v, ret=%v\n", v, r)
	}
}

func select_uninit() {
	var ch chan int
	go func() {
		time.Sleep(1 * time.Second)
		ch = make(chan int)
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for i := 0; i < 10; i++ {
		select {
		case v, r := <-ch:
			{
				fmt.Printf("select ch: %v (%v)", v, r)

			}
		default:
			fmt.Println("select_uninit: delay")
			time.Sleep(2 * time.Second)
		}
	}
}

func nil_select() {
	select {}
}

func long_running() chan interface{} {
	done := make(chan interface{})

	go func() {
		n := 0
		run := true

		for run {
			select {
			case <-done:
				{
					fmt.Println("READ DONE")
					run = false
				}
			default:
				{
					time.Sleep(1 * time.Second)
					n++
					fmt.Println("do job:", n)
				}
			}
		}
		fmt.Println("JOB DONE")
	}()
	return done
}

func n() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for {
			ch <- rand.Int()
		}
	}()
	return ch
}

func simple_pipline() {
	gen := func(done <-chan interface{}, xs ...int) <-chan int {
		rs := make(chan int)
		go func() {
			defer close(rs)
			for _, n := range xs {
				select {
				case <-done:
				case rs <- n:
				}
			}
		}()
		return rs
	}

	add := func(done <-chan interface{}, xs <-chan int, n int) <-chan int {
		rs := make(chan int)
		go func() {
			defer close(rs)
			for i := range xs {
				select {
				case <-done:
					return
				case rs <- i + n:
				}
			}
		}()
		return rs
	}

	mul := func(done <-chan interface{}, xs <-chan int, n int) <-chan int {
		rs := make(chan int)
		go func() {
			defer close(rs)
			for i := range xs {
				select {
				case <-done:
					return
				case rs <- i * n:
				}
			}
		}()
		return rs
	}

	done := make(chan interface{})
	defer close(done)

	for n := range mul(done, add(done, gen(done, 1, 2, 3, 4), 4), 2) {
		fmt.Println(">", n)
	}
}

func main() {
	times(func() {
		fmt.Print("HELLO")
	}, 10)
	// deadlock_01()
	for i := range seq(3) {
		fmt.Printf("%v\n", i)
	}

	//block(5)
	//close_nil_chan()
	//write_closed()
	//close_closed_ch()
	// wait(1)
	// select_closed_chan()
	// wait2(2)

	//select_default()
	//select_uninit()

	// ctrl := long_running()
	// time.Sleep(5 * time.Second)
	// fmt.Println("CLOSE")
	// ctrl <- 0
	//close(ctrl)

	// ctrl := unclose_chan()
	// time.Sleep(2 * time.Second)
	// close(ctrl)

	// fmt.Println("start wait3")
	// <-wait3(2)

	simple_pipline()
}
