package echosvc

import (
	"context"
	"log"
	"net"
	"time"

	"amas.org/echosvc/model"
	pb "amas.org/echosvc/model"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	"google.golang.org/grpc"
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
func (s *EchoSVC) Count(stream model.Echo_CountServer) error {

	for i := 0; i < 100; i++ {
		stream.Send(&wrappers.Int64Value{Value: i})
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
