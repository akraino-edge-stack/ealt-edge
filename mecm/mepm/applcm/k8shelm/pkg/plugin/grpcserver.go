package plugin

import (
	"bytes"
	"context"
	"k8shelm/internal/lcmservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"net"
	"os"
	"strconv"
)

type ServerGRPC struct {
	server      *grpc.Server
	port        int
	certificate string
	key         string
}

type ServerGRPCConfig struct {
	Certificate string
	Key         string
	Port        int
}

func NewServerGRPC(cfg ServerGRPCConfig) (s ServerGRPC, err error) {
	logger := log.New(os.Stdout, "helmplugin ", log.LstdFlags|log.Lshortfile)
	if cfg.Port == 0 {
		logger.Fatalf("Port must be specified")
	}
	s.port = cfg.Port
	s.certificate = cfg.Certificate
	s.key = cfg.Key
	logger.Println("Binding is successful")
	return
}

func (s *ServerGRPC) Listen() (err error) {
	logger := log.New(os.Stdout, "helmplugin ", log.LstdFlags|log.Lshortfile)
	var (
		listener  net.Listener
		grpcOpts  = []grpc.ServerOption{}
		grpcCreds credentials.TransportCredentials
	)

	logger.Println("Listening start")

	listener, err = net.Listen("tcp", ":"+strconv.Itoa(s.port))

	logger.Println("Listening end")

	if err != nil {
		logger.Fatalf("failed to listen on port: ", s.port)
	}

	if s.certificate != "" && s.key != "" {
		grpcCreds, err = credentials.NewServerTLSFromFile(
			s.certificate, s.key)
		if err != nil {
			logger.Fatalf("failed to create tls grpc server using cert %s and key: ", s.certificate, s.key)
		}

		grpcOpts = append(grpcOpts, grpc.Creds(grpcCreds))
	}

	logger.Println("New server creation")

	s.server = grpc.NewServer(grpcOpts...)

	logger.Println("New server creation success")

	lcmservice.RegisterAppLCMServer(s.server, s)

	logger.Println("New server registration success")


	err = s.server.Serve(listener)
	if err != nil {
		logger.Fatalf("errored listening for grpc connections")
	}

	logger.Println("Server is serving")


	return
}

func (s *ServerGRPC) Query(ctx context.Context, req *lcmservice.QueryRequest) (resp *lcmservice.QueryResponse, err error) {
	r := queryChart(req.GetWorkloadId(), req.GetHostIp())
	resp = &lcmservice.QueryResponse{
		Status: r,
	}
	return resp, nil
}

func (s *ServerGRPC) Terminate(ctx context.Context, req *lcmservice.TerminateRequest) (resp *lcmservice.TerminateResponse, err error) {
	logger := log.New(os.Stdout, "helmplugin ", log.LstdFlags|log.Lshortfile)
	uninstallChart(req.GetWorkloadId(), req.GetHostIp())
	resp = &lcmservice.TerminateResponse{
		Status: "Success",
	}
	logger.Printf("Termination completed")
	return resp, nil
}

func (s *ServerGRPC) Instantiate(stream lcmservice.AppLCM_InstantiateServer) (err error) {
	logger := log.New(os.Stdout, "helmplugin ", log.LstdFlags|log.Lshortfile)

	req, err := stream.Recv()
	if err != nil {
		logger.Fatalf("cannot receive package metadata")
	}

	hostIP := req.GetHostIp()
	logger.Printf("receive an upload-image request for package ip %s", hostIP)

	helmPkg := bytes.Buffer{}

	for {
		err := contextError(stream.Context())
		if err != nil {
			return err
		}

		logger.Printf("waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			logger.Printf("no more data")
			break
		}
		if err != nil {
			return logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		chunk := req.GetPackage()

		logger.Printf("received a chunk ")

		// write slowly
		// time.Sleep(time.Second)
		_, err = helmPkg.Write(chunk)
		if err != nil {
			return logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	relName := installChart(helmPkg, hostIP)

	res := &lcmservice.InstantiateResponse{
		WorkloadId: relName,
		Status:    "Success",
	}

	err = stream.SendAndClose(res)
	if err != nil {
		return logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}

	logger.Printf("Instantation completed")
	return
}


func (s *ServerGRPC) Close() {
	if s.server != nil {
		s.server.Stop()
	}
	return
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func logError(err error) error {
	if err != nil {
		log.Print(err)
	}
	return err
}