package echosvc

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"amas.org/echosvc/model"
	pb "amas.org/echosvc/model"
	empty "github.com/golang/protobuf/ptypes/empty"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	"github.com/grpc-ecosystem/go-grpc-middleware"

	// "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// EchoSVC is a helloworld service for gRPC
type EchoSVC struct {
	time     time.Duration
	opts     []grpc.ServerOption
	crt      string
	key      string
	svc      grpc.Server
	port     string
	hostname string
	ius      []grpc.UnaryServerInterceptor
	iss      []grpc.StreamServerInterceptor
}

// New
func New(port string) *EchoSVC {
	return &EchoSVC{port: port}
}

// Say is NOTING to say
func (s *EchoSVC) Say(ctx context.Context, msg *pb.Msg) (*pb.Msg, error) {
	if msg.Id < 0 {
		return nil, status.Error(codes.InvalidArgument, "Id must > 0")
	}

	timestamp := md(ctx, "timestamp")[0]
	logrus.WithField("hostname", s.hostname).Info("RECV:", msg, "timestamp", timestamp)

	r := new(pb.Msg)
	r.Id = msg.Id + 1
	r.Text = msg.Text
	r.From = s.hostname
	return r, nil
}

func md(ctx context.Context, key string) []string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md[key]) < 1 {
		return []string{""}
	}
	return md[key]
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

// AddUnaryInterceptor add unary server interceptors
func (s *EchoSVC) AddUnaryInterceptor(fns ...grpc.UnaryServerInterceptor) *EchoSVC {
	if len(fns) == 0 {
		return s
	}

	if len(s.ius) < 1 {
		s.ius = make([]grpc.UnaryServerInterceptor, 3)
	}
	s.ius = append(s.ius, fns...)
	return s
}

// WithTLS need cert file and key file to setup
func (s *EchoSVC) WithTLS(crt, key string) *EchoSVC {
	s.crt = crt
	s.key = key
	return s
}

func (s *EchoSVC) SetHostname(hostname string) *EchoSVC {
	s.hostname = hostname
	return s
}

func (s *EchoSVC) appendOptions(opts ...grpc.ServerOption) *EchoSVC {
	s.opts = append(s.opts, opts...)
	return s
}

func (s *EchoSVC) cred() (grpc.ServerOption, error) {
	if s.crt != "" {
		cert, err := tls.LoadX509KeyPair(s.crt, s.key)
		if err != nil {
			return nil, err
		}
		return grpc.Creds(credentials.NewServerTLSFromCert(&cert)), nil
	}
	return nil, nil
}

// Start EchoSvc
func (s *EchoSVC) Start() error {
	cred, err := s.cred()
	if err != nil {
		return err
	}

	s.appendOptions(cred)
	s.appendOptions(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(s.ius...)))

	l, err := net.Listen("tcp", s.port)
	if err != nil {
		return err
	}
	svc := grpc.NewServer(s.opts...)
	pb.RegisterEchoServer(svc, s)
	if err := svc.Serve(l); err != nil {
		return err
	}
	return nil
}
