package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"amas.org/echosvc/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	//addr = "localhost:8888"
	addr = "minikube:30100"
)

// EchoClient is gRPC client
type EchoClient struct {
	addr     string
	opts     []grpc.DialOption
	crt      string
	hostname string
	client   model.EchoClient
}

// NewEchoClient new echo client
func NewEchoClient(addr string) EchoClient {
	c := EchoClient{addr: addr}
	return c
}

// RPC is CallBuilder
func (r *EchoClient) RPC() *CallBuilder {
	return &CallBuilder{client: r, timeout: 2 * time.Second}
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

type basicAuth struct {
	username string
	password string
}

func (r basicAuth) GetRequestMetadata(ctx context.Context, in ...string) (map[string]string, error) {
	auth := r.username + ":" + r.password
	return map[string]string{
		"authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(auth)),
	}, nil
}

func (r basicAuth) RequireTransportSecurity() bool {
	return true
}

func (r *EchoClient) WithAuthBasic(user string, pwd string) *EchoClient {
	r.opts = append(r.opts, grpc.WithPerRPCCredentials(basicAuth{username: user, password: pwd}))
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

// CallBuilder is method for easy build rpc call
type CallBuilder struct {
	client   *EchoClient
	header   metadata.MD
	trailer  metadata.MD
	md       metadata.MD
	timeout  time.Duration
	deadline time.Time
	opts     []grpc.CallOption
}

// WithHeader add new header to header
func (r *CallBuilder) WithHeader(key string, value ...string) *CallBuilder {
	if r.header == nil {
		r.header = metadata.MD{}
	}
	r.header.Set(key, value...)
	return r
}

// WithTrailer is call build func
func (r *CallBuilder) WithTrailer(key string, value ...string) *CallBuilder {
	if r.trailer == nil {
		r.trailer = metadata.MD{}
	}
	r.trailer.Set(key, value...)
	return r
}

// WithMetadata is call build func
func (r *CallBuilder) WithMetadata(key string, value ...string) *CallBuilder {
	if r.md == nil {
		r.md = metadata.MD{}
	}
	r.md.Set(key, value...)
	return r
}

// WithTimeout set the gRPC timeout
func (r *CallBuilder) WithTimeout(t time.Duration) *CallBuilder {
	r.timeout = t
	return r
}

// WithDeadline set the gRPC call Deadline
func (r *CallBuilder) WithDeadline(t time.Duration) *CallBuilder {
	r.timeout = t
	return r
}

func (r *CallBuilder) buildContext() (ctx context.Context, cancel context.CancelFunc) {
	ctx, cancel = context.WithTimeout(context.Background(), r.timeout)
	if len(r.md) > 0 {
		ctx = metadata.NewOutgoingContext(ctx, r.md)
	}

	r.opts = append(r.opts, grpc.Header(&r.header))
	r.opts = append(r.opts, grpc.Trailer(&r.trailer))
	return
}

// WithGZip is enable gzip
func (r *CallBuilder) WithGZip() *CallBuilder {
	r.opts = append(r.opts, grpc.UseCompressor("HELLO"))
	return r
}

// Say with id & text
func (r *CallBuilder) Say(id int32, text string) (*model.Msg, error) {
	if err := r.client.dial(); err != nil {
		return nil, err
	}

	ctx, cancel := r.buildContext()
	defer cancel()

	msg, err := r.client.client.Say(ctx, &model.Msg{
		Id:   id,
		Text: text,
	})
	return msg, err
}

// func count() {
// 	conn, err := grpc.Dial(addr, grpc.WithInsecure())
// 	defer conn.Close()

// 	if err != nil {
// 		log.Fatalf("DID NOT CONNECT: %v", err)
// 	}
// 	c := model.NewEchoClient(conn)

// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	countc, err := c.Count(ctx, &empty.Empty{})
// 	for {
// 		v, err := countc.Recv()
// 		if err != nil {
// 			log.Fatalln(err)
// 			return
// 		}
// 		fmt.Printf("[count] : received %v\n", v)
// 	}
// }

func main() {
	//count()
	//echoc := EchoClient.N

	echoc := NewEchoClient(addr)
	//echoc.WithInsecure()
	echoc.WithTLS("cert/svc.crt", "localhost")
	echoc.WithAuthBasic("amas", "888888")

	for i := 0; i < 10; i++ {
		msg, err := echoc.RPC().WithGZip().WithMetadata("timestamp", time.Now().String()).Say(int32(i), fmt.Sprintf("HELLO %v", i))
		if err != nil {
			log.Fatalf("ERROR: %v\n", err)
		}
		fmt.Printf("%v\n", msg)
	}
}
