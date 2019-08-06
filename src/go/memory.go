package main

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"sync"
	"time"
)

func SystemMemory() uint64 {
	runtime.GC()
	var s runtime.MemStats
	runtime.ReadMemStats(&s)
	return s.Sys
}

func MemoryUsagePerGoKB(max int) {
	before := SystemMemory()
	for i := 0; i < max; i++ {
		go (func() {
			select {}
		})()
	}
	after := SystemMemory()
	numGo := runtime.NumGoroutine()

	mKb := float64(after-before) / 1000.0
	avg := mKb / float64(numGo)
	fmt.Printf("%d go : total %vKB , avg %2fKB\n", numGo, mKb, avg)
}

func testCond(max int) {
	c := sync.NewCond(&(sync.Mutex{}))
	for i := 0; i < max; i++ {
		go (func(n int) {
			c.L.Lock()
			fmt.Println(n)
		})(i)
	}
	fmt.Println("WAITING ...")
	time.Sleep(time.Second * 2)
	numGo := runtime.NumGoroutine()
	fmt.Printf("%d go\n", numGo)
	//c.Broadcast()
}

func PrintRuntimeInfo() {
}

func main() {
	//debug.SetGCPercent(1)

	debug.SetMaxThreads(8)
	//MemoryUsagePerGoKB(1000000)
	testCond(10)
	//_m := SystemMemory()
}
