package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func longTImeTask(ch chan string, max int) {
	time.Sleep(time.Second * time.Duration(max))
	ch <- strconv.Itoa(max)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	ch := make(chan string)

	go longTImeTask(ch, rand.Intn(6))
	select {
	case t := <-ch:
		fmt.Printf("WORK DONE in %v\n", t)
	case <-time.After(time.Second * time.Duration(3)):
		fmt.Print("TIMEOUT")
		close(ch)
	}
}
