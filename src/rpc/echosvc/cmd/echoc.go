package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"amas.org/echosvc/model"
	"google.golang.org/grpc"
)

const (
	addr = "localhost:8888"
)

func say(id int32, text string) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatalf("DID NOT CONNECT: %v", err)
	}

	c := model.NewEchoClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	// go!!!
	msg, err := c.Say(ctx, &model.Msg{
		Id:   id,
		Text: text,
	})

	if err != nil {
		log.Fatalf("CALL ERROR: %v", err)
	}

	fmt.Printf("%v\n", msg)
}

func main() {
	say(1, "a")
	say(2, "b")
}
