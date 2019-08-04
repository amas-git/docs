package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Printf("Hello World! %s", runtime.Version())
}
