package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"amas.org/echosvc/model"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	addr = "localhost:8888"
)

// EchoClient is gRPC client
type EchoClient struct {
	addr     string
	opts     []grpc.DialOption
	crt      string
	hostname string
	client   model.EchoClient
}

// New create new echo client
func NewEchoClient(addr string) EchoClient {
	c := EchoClient{addr: addr}
	return c
}

// WithTLS set crt file and hostname
func (r *EchoClient) WithTLS(crt, hostname string) *EchoClient {
	//cred, err := credentials.NewClientTLSFromFile(crt, hostname)
	//append(r.opts, grpc.WithTransportCredentials(cred))
	r.crt = crt
	r.hostname = hostname
	return r
}

// WithInsecure NOT USE TLS
func (r *EchoClient) WithInsecure() *EchoClient {
	r.opts = append(r.opts, grpc.WithInsecure())
	return r
}

func (r *EchoClient) dial() error {
	if r.client != nil {
		return nil
	}

	if r.crt != "" {
		cred, err := credentials.NewClientTLSFromFile(r.crt, r.hostname)
		if err != nil {
			return err
		}

		r.opts = append(r.opts, grpc.WithTransportCredentials(cred))
	}

	conn, err := grpc.Dial(r.addr, r.opts...)
	if err != nil {
		return fmt.Errorf("DIAL FAILED: %v", err)
	}

	r.client = model.NewEchoClient(conn)
	return nil
}

func (r *EchoClient) call() {

}

// Say with id & text
func (r *EchoClient) Say(id int32, text string) (*model.Msg, error) {
	if err := r.dial(); err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	msg, err := r.client.Say(ctx, &model.Msg{
		Id:   id,
		Text: text,
	})
	return msg, err
}

func count() {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	defer conn.Close()

	if err != nil {
		log.Fatalf("DID NOT CONNECT: %v", err)
	}
	c := model.NewEchoClient(conn)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	countc, err := c.Count(ctx, &empty.Empty{})
	for {
		v, err := countc.Recv()
		if err != nil {
			log.Fatalln(err)
			return
		}
		fmt.Printf("[count] : received %v\n", v)
	}
}

func main() {
	//count()
	//echoc := EchoClient.N

	echoc := NewEchoClient(addr)
	//echoc.WithInsecure()
	echoc.WithTLS("cert/svc.crt", "localhost")

	for i := 0; i < 10; i++ {
		msg, err := echoc.Say(int32(i), fmt.Sprintf("HELLO %v", i))
		if err != nil {
			log.Fatalf("ERROR: %v\n", err)
		}
		fmt.Printf("%v\n", msg)
	}
}
