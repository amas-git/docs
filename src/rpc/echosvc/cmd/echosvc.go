package main

import (
	echosvc "amas.org/echosvc/svc"
	"amas.org/echosvc/utils"
)

const (
	crt = "cert/svc.crt"
	key = "cert/svc.key"
)

func main() {
	utils.HelloWorld()

	svc := echosvc.New()
	svc.StartSEC(":8888", crt, key)
	//svc.Start(":8888")

}
