package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func arg(i int, value string) string {
	if i > len(os.Args)-1 || os.Args[i] == "" {
		return value
	}
	return os.Args[i]
}

func hostname() string {
	if s, err := os.Hostname(); err == nil {
		return s
	}
	return "n/a"
}

func main() {
	addr := arg(1, ":6666")
	tags := arg(2, "v1")

	fmt.Printf("START WITH %v [%v]\n", addr, tags)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("HELLO %s FROM %s\n", tags, hostname())))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Fatal(http.ListenAndServe(addr, nil))
}
