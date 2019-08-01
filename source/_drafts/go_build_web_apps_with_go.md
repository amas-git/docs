# Building web apps with go



## 需要安装的包

````go
$ go get -u github.com/julienschmidt/httprouter 
$ go get -u github.com/codegangsta/negroni 
$ go get -u github.com/russross/blackfriday
$ go get -u gopkg.in/unrolled/render.v1
$ go get -u github.com/mattn/go-sqlite3
````



## net/http

这个是go自带的http包，我们先从最基本的开始。

这个就是处理http请求的接口

```go
type Handler interface {
		ServeHTTP(ResponseWriter, *Request)
}

type ResponseWriter interface {
  Header() Header
  Write([]byte) (int, error)
  WriteHeader(int)
}
```



simple_httpd/main.go:

```go
package main

import "net/http"

func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
```

- http.ListenAndServe(addr string, handler Handler)
- http.Dir就是string

编译:

```go
$ cd simple_httpd
$ ls
main.go
$ go build
$ ./simple_httpd
// 现在就可以通过 http://localhost:8080 访问服务了
```

添加路由

```go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	//http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
	server()
}

func server() {
	http.HandleFunc("/hello", hello)
	http.Handle("/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8080", nil)
}

func hello(rw http.ResponseWriter, r *http.Request) {
  r.Header // HTTP头，本质就是map
  
	 rw.Write([]byte("Hello boy"))
}
```

使用http router

```go
func server2(port string) {
  r := httprouter.New()
  r.GET("/hello/:id", func(rw http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintln(rw, "hello", p.ByName("id"))
  })

  http.ListenAndServe(":"+port, r)
}
```

