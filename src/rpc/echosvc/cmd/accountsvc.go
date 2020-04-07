package main

import (
	_ "database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
)

// GET simple url
func GET(url string) string {
	r, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer r.Body.Close()

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return ""
	}
	return string(b)
}

func main() {

	//fmt.Println(GET("http://localhost:8888"))
	s := "hello"
	ts := reflect.TypeOf(s)
	fmt.Println(ts.Kind())
}
