package main

import (
	"fmt"

	lru "github.com/hashicorp/golang-lru"
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

func main() {
	fmt.Printf("hello ark")
	//checkCache()
}
