package main

import (
	"context"
	"fmt"

	echosvc "amas.org/echosvc/svc"
	"amas.org/echosvc/utils"
	"google.golang.org/grpc"
)

const (
	crt = "cert/svc.crt"
	key = "cert/svc.key"
)

func logInterceptor(ctx context.Context, r interface{}, i *grpc.UnaryServerInfo, h grpc.UnaryHandler) (interface{}, error) {
	fmt.Printf("[PRE  CALL]: %v\n", i)
	m, err := h(ctx, r)
	fmt.Printf("[POST CALL]: %v\n", m)
	return m, err
}

func main() {
	utils.HelloWorld()

	svc := echosvc.New(":8888")
	svc.WithTLS(crt, key)
	svc.SetUnaryInterceptor(logInterceptor)
	svc.Start()
}
