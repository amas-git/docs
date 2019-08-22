package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var balance = 1000
var taken = 0

// get money rnadom
func getMoney(seq, x int) (r bool) {
	defer func() {
		fmt.Printf("[%02d] : %5v (%02d/%03d)\n", seq, r, x, balance)
	}()

	if balance-x < 0 {
		r = false
		return
	}
	balance -= x
	r = true
	taken += x
	return
}

var mutex = new(sync.Mutex)

func getMoney2(seq, x int) (r bool) {
	wg.Add(1)
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()

	defer func() {
		fmt.Printf("[%02d] : %5v (%02d/%03d)\n", seq, r, x, balance)
	}()

	if balance-x < 0 {
		r = false
		return
	}
	balance -= x
	r = true
	taken += x
	return
}

var wg sync.WaitGroup

func main() {
	//runtime.GOMAXPROCS(2)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		go getMoney2(i, rand.Intn(25))
	}

	// 简单的等待2秒
	//time.Sleep(time.Second * 2)
	wg.Wait()
	fmt.Printf("FINAL BALANCE : %d TAKEN %d", balance, taken)
}
