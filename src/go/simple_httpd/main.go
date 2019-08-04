package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func main() {
	//http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	port := os.Getenv("PORT")
	if port == "" {
		port = "1983"
	}
	server2(port)
}

func server(port string) {
	fmt.Println("START HTTP ON PORT", port)
	http.HandleFunc("/hello", hello)
	http.Handle("/", http.FileServer(http.Dir("publicjjj")))
	http.ListenAndServe(":"+port, nil)
}

func server2(port string) {
	r := httprouter.New()
	r.GET("/hello/:id", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintln(rw, "hello", p.ByName("id"))
	})

	http.ListenAndServe(":"+port, r)
}

func hello(rw http.ResponseWriter, r *http.Request) {
	method := r.Method
	// body := r.FormValue("body")
	fmt.Println(method)
	rw.Write([]byte("Hello boy " + r.URL.String()))
}
