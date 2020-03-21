package main

import (
	echosvc "amas.org/echosvc/svc"
	"amas.org/echosvc/utils"
)

func main() {
	utils.HelloWorld()

	svc := echosvc.New()
	svc.Start(":8888")

}
