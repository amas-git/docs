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

## Form

## CSFR

## 

## Authentication 和 Sessions

验证用户发起的http请求是否有获取资源的权限。

Why?

 	1. HTTP请求是无状态的，就是我们没办法区分一个请求是不是与其他请求相关，如果回到认证的问题上，假如HTTP允许我们知道当前的请求是和之前某个登录成功请求相关的，那我们就不必在当前请求里验证这用户是不是登录过。
 	2. 没等用户访问资源，需要提供用户名/密码POST过来，我们与数据库中的进行比对，如果通过就返回资源，否则就拒绝，但每个HTTP请求都加入这样的validate环节会影响性能。

可以通过sessionID解决这个问提， 我们可以让用户登录成功请求之后设置一个唯一的sessionID到Cookie中， 后续用户发来的请求携带这个sessionID, 这样就不用频繁validate了。 可以使用`gorilla/sessions`包来管理sessionID:

```go
import (
 "net/http"
 "github.com/gorilla/sessions"
)
```

```go
//Store the cookie store which is going to store session data in the cookie
var Store = sessions.NewCookieStore([]byte("secret-password"))

 //IsLoggedIn will check if the user has an active session and return True
func IsLoggedIn(r *http.Request) bool {
	session, _ := Store.Get(r, "session")
	if session.Values["loggedin"] == "true" {
 		return true
 	}
	return false
}
```



## Cookies

Cookies是浏览器的一种存储机制，常用来存放一些用户相关的数据，CSFR Token, 用户名密码都可能保存在Cookie中，如果你不做限制，别人网站可以用js读取你网站的Cookies, 这是件可怕的事情。所以你要把Cookie设置成HttpOnly的，只有你的域名才可以访问你的Cookies. 不要把你的饼干风非别人吃。。。



## 注册

注册的过程大概是这样的

我们把username/password等信息通过HTTP Form POST到服务端， 然后存储起来，这样就完成了注册。

```go
func SignUpFunc(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        r.ParseForm()
        username := r.Form.Get("username")
        password := r.Form.Get("password")
        // 保存到数据库中
        if err := createUser(username, passord); err != nil {
            http.Error(w, "Sign up error", http.StatusInternalServerError)	
        }  else {
            http.Redirect(w, r, "/login", 302)
				}
    }
}
```





## 登录

我们把username/password等信息通过HTTP Form POST到服务端， 通过Validate后，我们需要建立session, 方法是写入cookies:

```go
session.Values["loggedin"] = "true"
session.Values["username"] = "amas"
session.Save(r, w)
```



## 中间件

## API

API与用户通过浏览器访问网页只有一个小小的区别，向我们发送请求的客户端需要自己管理session相关的数据，这种情况下我们一般叫做token.

	1. username/password通过HTTP发送过来，得到一个token
 	2. 在后续的请求中带上这个token, 一般保存在请求的Header中

## JWT

> JSON Web Tokens



















## 参考

