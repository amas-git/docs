package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	echosvc "amas.org/echosvc/svc"
	"amas.org/echosvc/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
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

func loginInterceptor(ctx context.Context, r interface{}, i *grpc.UnaryServerInfo, h grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.InvalidArgument, "NEED METADATA")
	}

	valid := func(auth []string) bool {
		if len(auth) < 1 {
			return false
		}
		token := strings.TrimPrefix(auth[0], "Basic ")
		return token == base64.StdEncoding.EncodeToString([]byte("amas:888888"))
	}

	if !valid(md["authorization"]) {
		return nil, status.Errorf(codes.Unauthenticated, "INVALID CREDNTIALS")
	}

	return h(ctx, r)
}

func main() {
	utils.HelloWorld()

	svc := echosvc.New(":8888")
	svc.WithTLS(crt, key)
	svc.SetUnaryInterceptor(loginInterceptor)
	if err := svc.Start(); err != nil {
		fmt.Errorf("START FAILED %v\n", err)
	}
}
