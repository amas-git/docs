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

func main() {
	addr := arg(1, ":6666")
	tags := arg(2, "v1")

	fmt.Printf("START WITH %v [%v]\n", addr, tags)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(fmt.Sprintf("HELLO %s", tags)))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	log.Fatal(http.ListenAndServe(addr, nil))
}