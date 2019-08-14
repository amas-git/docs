package main

import (
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// MakeFile make specify size file
func MakeFile(path string, size int) (int, error) {
	if size < 0 {
		size = 0
	}

	target, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return -1, err
	}
	defer target.Close()

	n := size
	for n > 0 {
		buf := make([]byte, func() int {
			if n > 1024 {
				return 1024
			}
			return n
		}())
		rand.Read(buf)
		x, _ := target.Write(buf)
		n -= x
	}
	return size, nil
}

func main() {
	MakeFile("/tmp/100k", 100*1024+19)
}
