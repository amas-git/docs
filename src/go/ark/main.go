package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	lru "github.com/hashicorp/golang-lru"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var cache *lru.Cache

func init() {
	fmt.Println("INIT...")
	cache, _ = lru.NewWithEvict(2,
		func(key interface{}, value interface{}) {
			// 移除缓存时调用
			fmt.Printf("Evicted: key=%v value=%v\n", key, value)
		},
	)
}

func checkCache() {
	cache.Add(1, "a")
	cache.Add(2, "b")
	// adds 1
	// adds 2; cache is now at capacity
	fmt.Println(cache.Get(1)) // "a true"; 1 now most recently used
	cache.Add(3, "c")         // adds 3, evicts key 2
	fmt.Println(cache.Get(2)) // "<nil> false" (not found)
}

func shutdown(s os.Signal) {
	fmt.Println(s)
}

func main() {
	fmt.Printf("hello ark")
	//checkCache()

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Print("hello world")
	log.Print("hello world")
	<-make(chan int)
}

func init() {
	go func() {
		// Create a quit channel which carries os.Signal values.
		quit := make(chan os.Signal, 1)
		// Use signal.Notify() to listen for incoming SIGINT and SIGTERM signals and
		// relay them to the quit channel. Any other signals will not be caught by
		// signal.Notify() and will retain their default behavior.
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		// Read the signal from the quit channel. This code will block until a signal is
		// received.
		s := <-quit
		shutdown(s)
		// Log a message to say that the signal has been caught. Notice that we also
		// call the String() method on the signal to get the signal name and include it
		// in the log entry properties.
		// Exit the application with a 0 (success) status code.
		os.Exit(0)
	}()
}
