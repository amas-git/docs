package echosvc

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net"
	"time"

	"amas.org/echosvc/model"
	pb "amas.org/echosvc/model"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// EchoSVC is a helloworld service for gRPC
type EchoSVC struct {
	time time.Duration
	opts []grpc.ServerOption
	crt  string
	key  string
	svc  grpc.Server
	port string
}

// New
func New(port string) *EchoSVC {
	return &EchoSVC{port: port}
}

// Say is NOTING to say
func (s *EchoSVC) Say(ctx context.Context, msg *pb.Msg) (*pb.Msg, error) {
	log.Printf("[echosvc] : RECIVE <- %v\n", msg)
	r := new(pb.Msg)
	r.Id = msg.Id + 1
	r.Text = msg.Text
	return r, nil
}

// Count is count
func (s *EchoSVC) Count(_ *empty.Empty, stream model.Echo_CountServer) error {
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Second)
		if err := stream.Send(&wrappers.Int64Value{Value: int64(i)}); err != nil {
			return err
		}
	}
	return nil
}

// Ask is ack
func (s *EchoSVC) Ack(stream model.Echo_AckServer) error {

	return nil
}

// AddUnaryInterceptor is cool
func (s *EchoSVC) SetUnaryInterceptor(fn grpc.UnaryServerInterceptor) *EchoSVC {
	s.opts = append(s.opts, grpc.UnaryInterceptor(fn))
	return s
}

// WithTLS need cert file and key file to setup
func (s *EchoSVC) WithTLS(crt, key string) *EchoSVC {
	s.crt = crt
	s.key = key
	return s
}

func (s *EchoSVC) Start() error {
	if s.crt != "" {
		cert, err := tls.LoadX509KeyPair(s.crt, s.key)
		if err != nil {
			return err
		}
		s.opts = append(s.opts, grpc.Creds(credentials.NewServerTLSFromCert(&cert)))
	}

	l, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}
	svc := grpc.NewServer(s.opts...)
	pb.RegisterEchoServer(svc, s)
	if err := svc.Serve(l); err != nil {
		return err
	}

	fmt.Println("STARTED: ", s.port)
	return nil
}
