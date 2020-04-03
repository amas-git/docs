package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func repeat(done <-chan interface{}, values ...interface{}) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case ch <- v:
				}
			}
		}
	}()
	return ch
}

func repeatX(ctx context.Context, values ...interface{}) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			for _, v := range values {
				select {
				case <-ctx.Done():
					return
				case ch <- v:
				}
			}
		}
	}()
	return ch
}

func take(done <-chan interface{}, input <-chan interface{}, n int) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case ch <- <-input:
			}
		}
	}()
	return ch
}

func repeatFn(done <-chan interface{}, fn func() interface{}) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case ch <- fn():
			}
		}
	}()
	return ch
}

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func fanIn(done <-chan interface{}, chs ...<-chan interface{}) <-chan interface{} {
	ch := make(chan interface{})
	var wg sync.WaitGroup

	read := func(c <-chan interface{}) {
		defer wg.Done()
		for {
			select {
			case <-done:
				return
			case ch <- <-c:
			}
		}
	}

	wg.Add(len(chs))
	for _, c := range chs {
		go read(c)
	}

	go func() {
		defer close(ch)
		wg.Wait()
	}()
	return ch
}

func play_fanin() {
	done := make(chan interface{})
	a := repeatFn(done, func() interface{} {
		time.Sleep(1 * time.Second)
		return "a"
	})

	b := repeatFn(done, func() interface{} {
		time.Sleep(5 * time.Second)
		return "b"
	})

	c := repeatFn(done, func() interface{} {
		time.Sleep(1 * time.Second)
		return randInt(999, 9999)
	})

	i := 0
	for v := range take(done, fanIn(done, a, b, c), 100) {
		fmt.Println(i, v)
		i++
	}
	close(done)
}

func orDone(done <-chan interface{}, in <-chan interface{}) <-chan interface{} {
	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if ok == false {
					return
				}
				select {
				case ch <- v:
				case <-done:
				}
			}
		}
	}()
	return ch
}

func tee(done <-chan interface{}, in <-chan interface{}) (<-chan interface{}, <-chan interface{}) {
	out1 := make(chan interface{})
	out2 := make(chan interface{})

	go func() {
		defer close(out1)
		defer close(out2)

		for v := range in {
			o1, o2 := out1, out2
			for i := 0; i < 2; i++ {
				select {
				case <-done:
					return
				case o1 <- v:
					{
						o1 = nil
					}
				case o2 <- v:
					{
						o2 = nil
					}
				}
			}
		}
	}()

	return out1, out2
}

func split(done <-chan interface{}, in <-chan interface{}, router map[string](func(interface{}) bool)) map[string](<-chan interface{}) {
	chmap := make(map[string](chan interface{}), len(router))
	rmap := make(map[string](<-chan interface{}), len(router))
	for key := range router {
		chmap[key] = make(chan interface{})
		rmap[key] = chmap[key]
	}

	go func() {
		for _, ch := range chmap {
			defer close(ch)
		}

		for v := range in {
			fmt.Println(v)
			for key, fn := range router {
				//fmt.Printf("%v <- %v", key, v)
				if fn(v) {
					tch := chmap[key]
					fmt.Println(key, v)
					select {
					case <-done:
						return
					case tch <- v:
						{
							tch = nil
						}
					}
				}
			}
		}
	}()
	return rmap
}

func play_split() {
	done := make(chan interface{})
	defer close(done)
	isodd := func(n interface{}) bool {
		return n.(int)%2 == 0
	}

	iseven := func(n interface{}) bool {
		return !isodd(n)
	}

	chs := split(done, seq(100), map[string](func(interface{}) bool){
		"odd":  isodd,
		"even": iseven,
	})
	fmt.Print(chs)

	for k := range chs {
		for x := range chs[k] {
			fmt.Printf("[%v] : %v\n", k, x)
		}
	}
}

func slowFn(fn func() interface{}, t time.Duration) interface{} {
	time.Sleep(t * time.Second)
	return fn()
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

func play_tee() {
	done := make(chan interface{})
	defer close(done)

	ch1, ch2 := tee(done, seq(10))

	for v := range ch1 {
		fmt.Println("[1]", v, "[2]", <-ch2)
	}
}

func play_pipline() {
	done := make(chan interface{})
	for v := range take(done, repeat(done, 1, 2, 3), 10) {
		fmt.Println(v)
	}
	close(done)
}

func play_forselect_break_infinitloop() {
	for {
		fmt.Println("hello")
		select {
		default:
			break
		}
	}
}

func play_ordone() {
	seq := func(n int) <-chan int {
		ch := make(chan int)
		go func() {
			defer close(ch)
			for i := 0; i < n; i++ {
				ch <- i
			}
		}()
		return ch
	}

	n10 := seq(10)
	for n := range n10 {
		fmt.Println("play_ordone 1:", n)
	}
	for n := range n10 {
		fmt.Println("play_ordone 2:", n)
	}
}

func play_ctx() {
	ctx, _ := context.WithTimeout(context.TODO(), 1*time.Second)
	//defer cancel()
	ctx.Err()
	for v := range repeatX(ctx, 1) {
		fmt.Println("ctx:", v)
	}
}

func div(a, b int) {

}

func pipe(ctx context.Context, in <-chan interface{}, fns ...func(interface{}) interface{}) <-chan interface{} {

	if len(fns) == 1 {
		return func(ctx context.Context, in <-chan interface{}) <-chan interface{} {
			rch := make(chan interface{})
			go func() {
				defer close(rch)
				for {
					select {
					case rch <- func() interface{} {
						if in == nil {
							return fns[0](nil)
						}
						return fns[0](<-in)
					}():
					case <-ctx.Done():
						return
					}
				}
			}()
			return rch
		}(ctx, in)
	}
	x, xs := fns[0], fns[1:]
	return pipe(ctx, pipe(ctx, in, x), xs...)
}

func play_pipe() {
	ctx, cancel := context.WithCancel(context.TODO())
	num := 0
	double := func(n interface{}) interface{} {
		time.Sleep(1 * time.Second)
		if n == nil {
			return nil
		}
		return (n.(int)) * 2
	}

	for x := range pipe(ctx, nil,
		func(_ interface{}) interface{} {
			num++
			if num > 20 {
				cancel()
			}
			return num
		},
		double,
		double,
		double,
		double,
	) {
		fmt.Printf("%v\n", x)
	}
	cancel()
}

func play_nil_chan() {
	var ch chan int
	<-ch
}

const a = 1
const f float32 = a
const d float64 = a

func printType(a interface{}) {
	s := ""
	switch a.(type) {
	case int:
		s = "int"
	case string:
		s = "string"
	default:
		fmt.Printf("Unknown: %T\n", a)
	}
	fmt.Printf("%v is %T (%v)", a, a, s)
}

func play_ordone1() {

	ch := seq(10)
	for i := 0; i < 15; i++ {
		x, ok := <-ch
		fmt.Println("->", x, "(", ok, ")")
	}
}

func main() {
	fmt.Println("hello")
	//play_pipline()

	// done := make(chan interface{})
	// for n := range take(done, repeatFn(done, func() interface{} {
	// 	return randInt(100, 999)
	// }), 10) {
	// 	fmt.Println(">", n)
	// }
	// close(done)
	//play_fanin()
	//play_forchan()
	// play_ordone()
	//play_tee()
	//play_ctx()
	//play_pipe()
	//play_nil_chan()
	// printType(1)
	// printType(1i)
	// printType(nil)
	// printType(false)
	// printType(context.TODO())
	// for range []int{1, 1, 1, 1} {
	// 	fmt.Println("Looping")
	// }

	//play_tee()
	//play_ordone1()

	play_split()
	// for v := range seq(10) {
	// 	fmt.Println(v)
	// }
}
