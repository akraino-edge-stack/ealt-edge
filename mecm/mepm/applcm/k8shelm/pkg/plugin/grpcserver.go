/*
 * Copyright 2020 Huawei Technologies Co., Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package plugin

import (
	"bytes"
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/status"
	"io"
	"k8shelm/internal/lcmservice"
	"net"
	"os"
	"strconv"
)

// GRPC server
type ServerGRPC struct {
	server      *grpc.Server
	port        int
	certificate string
	key         string
	logger      *logrus.Logger
}

// GRPC service configuration used to create GRPC server
type ServerGRPCConfig struct {
	Certificate string
	Key         string
	Port        int
	Logger      *logrus.Logger
}

// Constructor to GRPC server
func NewServerGRPC(cfg ServerGRPCConfig) (s ServerGRPC) {
	s.logger = cfg.Logger
	s.port = cfg.Port
	s.certificate = cfg.Certificate
	s.key = cfg.Key
	s.logger.Infof("Binding is successful")
	return
}

// Start GRPC server and start listening on the port
func (s *ServerGRPC) Listen() (err error) {
	var (
		listener  net.Listener
		grpcOpts  = []grpc.ServerOption{}
		grpcCreds credentials.TransportCredentials
	)

	// Listen announces on the network address
	listener, err = net.Listen("tcp", ":"+strconv.Itoa(s.port))
	if err != nil {
		s.logger.Fatalf("failed to listen on specified port")
	}
	s.logger.Infof("Server started listening on specified port")

	// Secure connection if asked
	if s.certificate != "" && s.key != "" {
		grpcCreds, err = credentials.NewServerTLSFromFile(
			s.certificate, s.key)
		if err != nil {
			s.logger.Fatalf("failed to create tls grpc server using given cert and key")
		}
		grpcOpts = append(grpcOpts, grpc.Creds(grpcCreds))
	}

	// Register server with GRPC
	s.server = grpc.NewServer(grpcOpts...)
	lcmservice.RegisterAppLCMServer(s.server, s)

	s.logger.Infof("Server registered with GRPC")

	// Server start serving
	err = s.server.Serve(listener)
	if err != nil {
		s.logger.Fatalf("failed to listen for grpc connections. Err: %s", err)
		return err
	}
	return
}

// Query HELM chart
func (s *ServerGRPC) Query(ctx context.Context, req *lcmservice.QueryRequest) (resp *lcmservice.QueryResponse, err error) {

	// Input validation
	if (req.GetHostIp() == "") || (req.GetWorkloadId() == "") {
		return nil, s.logError(status.Errorf(codes.InvalidArgument, "HostIP & WorkloadId can't be null", err))
	}

	// Create HELM Client
	hc, err := NewHelmClient(req.GetHostIp(), s.logger)
	if os.IsNotExist(err) {
		return nil, s.logError(status.Errorf(codes.InvalidArgument, "Kubeconfig corresponding to given Edge can't be found. " +
			"Err: %s", err))
	}

	// Query Chart
	r, err := hc.queryChart(req.GetWorkloadId())
	if (err != nil) {
		return nil, s.logError(status.Errorf(codes.NotFound, "Chart not found for workloadId: %s. Err: %s",
			req.GetWorkloadId(), err))
	}
	resp = &lcmservice.QueryResponse{
		Status: r,
	}
	return resp, nil
}

// Terminate HELM charts
func (s *ServerGRPC) Terminate(ctx context.Context, req *lcmservice.TerminateRequest) (resp *lcmservice.TerminateResponse, err error) {
	// Input validation
	if (req.GetHostIp() == "") || (req.GetWorkloadId() == "") {
		return nil, s.logError(status.Errorf(codes.InvalidArgument, "HostIP & WorkloadId can't be null", err))
	}

	// Create HELM client
	hc, err := NewHelmClient(req.GetHostIp(), s.logger)
	if os.IsNotExist(err) {
		return nil, s.logError(status.Errorf(codes.InvalidArgument, "Kubeconfig corresponding to given Edge can't be found. " +
			"Err: %s", err))
	}

	// Uninstall chart
	err = hc.uninstallChart(req.GetWorkloadId())

	if (err != nil) {
		resp = &lcmservice.TerminateResponse{
			Status: "Failure",
		}
		return resp, s.logError(status.Errorf(codes.NotFound, "Chart not found for workloadId: %s. Err: %s",
			req.GetWorkloadId(), err))
	} else {
		resp = &lcmservice.TerminateResponse{
			Status: "Success",
		}
		return resp, nil
	}
}

// Instantiate HELM Chart
func (s *ServerGRPC) Instantiate(stream lcmservice.AppLCM_InstantiateServer) (err error) {

	// Recieve metadata which is host ip
	req, err := stream.Recv()
	if err != nil {
		s.logger.Errorf("Cannot receive package metadata. Err: %s", err)
		return
	}

	hostIP := req.GetHostIp()
	s.logger.Info("Recieved instantiate request")

	// Host validation
	if (hostIP == "") {
		return s.logError(status.Errorf(codes.InvalidArgument, "HostIP & WorkloadId can't be null", err))
	}

	// Receive package
	helmPkg := bytes.Buffer{}
	for {
		err := s.contextError(stream.Context())
		if err != nil {
			return err
		}

		s.logger.Debug("Waiting to receive more data")

		req, err := stream.Recv()
		if err == io.EOF {
			s.logger.Debug("No more data")
			break
		}
		if err != nil {
			return s.logError(status.Errorf(codes.Unknown, "cannot receive chunk data: %v", err))
		}

		// Receive chunk and write to helm package
		chunk := req.GetPackage()

		s.logger.Info("Recieved chunk")

		_, err = helmPkg.Write(chunk)
		if err != nil {
			return s.logError(status.Errorf(codes.Internal, "cannot write chunk data: %v", err))
		}
	}

	// Create HELM client
	hc, err := NewHelmClient(req.GetHostIp(), s.logger)
	if os.IsNotExist(err) {
		return s.logError(status.Errorf(codes.InvalidArgument, "Kubeconfig corresponding to edge can't be found. " +
			"Err: %s", err))
	}

	relName, err := hc.installChart(helmPkg)

	var res lcmservice.InstantiateResponse
	res.WorkloadId = relName

	if (err != nil) {
		res.Status = "Failure"
	} else {
		res.Status = "Success"
	}

	err = stream.SendAndClose(&res)
	if err != nil {
		return s.logError(status.Errorf(codes.Unknown, "cannot send response: %v", err))
	}
	s.logger.Info("Successful Instantiation")
	return
}

func (s *ServerGRPC) contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		return s.logError(status.Error(codes.Canceled, "request is canceled"))
	case context.DeadlineExceeded:
		return s.logError(status.Error(codes.DeadlineExceeded, "deadline is exceeded"))
	default:
		return nil
	}
}

func (s *ServerGRPC) logError(err error) error {
	if err != nil {
		s.logger.Errorf("Error Information: ", err)
	}
	return err
}