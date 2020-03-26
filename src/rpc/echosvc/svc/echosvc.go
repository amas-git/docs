package echosvc

import (
	"context"
	"crypto/tls"
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
}

func New() *EchoSVC {
	return new(EchoSVC)
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

// Start is start
func (s *EchoSVC) Start(port string) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("FAILED TO CREATE SERVER @%v : %v", port, err)
	}
	svc := grpc.NewServer()
	pb.RegisterEchoServer(svc, s)
	log.Printf("GRPC SVC START LISTEN @%v\n", port)
	if err := svc.Serve(l); err != nil {
		log.Fatalf("FAILED TO START: %v", err)
	}
}

// StartSEC is start
// create crt:
// $ openssl req -nodes -x509 -newkey rsa:4096 -keyout svc.key -out svc.crt -days 365
func (s *EchoSVC) StartSEC(port string, crt string, key string) {
	l, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("FAILED TO CREATE SERVER @%v : %vn", port, err)
		return
	}

	cert, err := tls.LoadX509KeyPair(crt, key)
	if err != nil {
		log.Fatalf("FAILED TO LOAD CERT : %v\n", err)
		return
	}

	opts := []grpc.ServerOption{
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
	}
	svc := grpc.NewServer(opts...)
	pb.RegisterEchoServer(svc, s)
	log.Printf("GRPC SVC START LISTEN @%v\n", port)
	if err := svc.Serve(l); err != nil {
		log.Fatalf("FAILED TO START: %v", err)
	}
}
