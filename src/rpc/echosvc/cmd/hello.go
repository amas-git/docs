package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
)

var version string
var (
	app_hello_processed_ops_total = promauto.NewCounter(prometheus.CounterOpts{
		Name: "app_hello_processed_ops_total",
		Help: "The total number of hello",
	})
)

// make smaller bin
// go build -ldflags="-s -w" cmd/hello.go
// 进一步可以使用upx工具减少包的大小： https://blog.filippo.io/shrink-your-go-binaries-with-this-one-weird-trick/
// 更实用的HTTP库: https://www.gorillatoolkit.org/pkg/mux#overview
// curl -i http://localhost:6666
func main() {
	addr := func() string {
		if len(os.Args) < 2 {
			return ":6666"
		}
		return os.Args[1]
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "HELLO")
		app_hello_processed_ops_total.Inc()
	})
	http.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("templates/info-template.html")
		if err != nil {
			logrus.WithError(err)
			return
		}
		err = t.Execute(w, struct{ Name string }{Name: "hello"})
		if err != nil {
			fmt.Errorf("%v\n", err)
		}
	})
	logrus.WithField("app", "hello").WithField("ver", version).Info("listen on ", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
