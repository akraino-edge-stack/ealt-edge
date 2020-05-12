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
	"broker/internal/lcmservice"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	_ "google.golang.org/grpc/encoding/gzip"
	"io"
	"log"
	"os"
)

// ClientGRPC provides the implementation of a file
// uploader that streams chunks via protobuf-encoded
// messages.
type ClientGRPC struct {
	conn      *grpc.ClientConn
	client    lcmservice.AppLCMClient
	chunkSize int
}

type ClientGRPCConfig struct {
	Address         string
	ChunkSize       int
	RootCertificate string
}

func NewClientGRPC(cfg ClientGRPCConfig) (c ClientGRPC, err error) {

	logger := log.New(os.Stdout, "broker ", log.LstdFlags|log.Lshortfile)

	var (
		grpcOpts  = []grpc.DialOption{}
		grpcCreds credentials.TransportCredentials
	)

	if cfg.Address == "" {
		logger.Fatalf("address must be specified: ", err)
	}

	if cfg.RootCertificate != "" {
		grpcCreds, err = credentials.NewClientTLSFromFile(cfg.RootCertificate, "localhost")
		if err != nil {
			logger.Fatalf("failed to create grpc tls client via root-cert: ", err)
		}

		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(grpcCreds))
	} else {
		grpcOpts = append(grpcOpts, grpc.WithInsecure())
	}

	switch {
	case cfg.ChunkSize == 0:
		logger.Fatalf("ChunkSize must be specified")
	case cfg.ChunkSize > (1 << 22):
		logger.Fatalf("ChunkSize must be < than 4MB")
	default:
		c.chunkSize = cfg.ChunkSize
	}

	c.conn, err = grpc.Dial(cfg.Address, grpcOpts...)
	if err != nil {
		logger.Fatalf("failed to start grpc connection with address: ", cfg.Address)
	}

	c.client = lcmservice.NewAppLCMClient(c.conn)
	return
}

func (c *ClientGRPC) Instantiate(ctx context.Context, f string, hostIP string) (workloadId string, status string, error error) {
	var (
		writing = true
		buf     []byte
		n       int
		file    *os.File
	)
	log.Printf("hostIP: ", hostIP)
	log.Printf("deployArtifact: ", f)
	logger := log.New(os.Stdout, "broker ", log.LstdFlags|log.Lshortfile)

	// Get a file handle for the file we
	// want to upload
	file, err := os.Open(f)
	if err != nil {
		logger.Fatalf("failed to open file: ", err.Error())
	}
	defer file.Close()

	// Open a stream-based connection with the
	// gRPC server
	stream, err := c.client.Instantiate(ctx)

	if err != nil {
		logger.Fatalf("failed to create upload stream for file: ", err)
	}
	defer stream.CloseSend()

    //send metadata information
	req := &lcmservice.InstantiateRequest{

		Data: &lcmservice.InstantiateRequest_HostIp{
				HostIp:  hostIP,
		},
	}

	err = stream.Send(req)
	if err != nil {
		logger.Fatalf("failed to send metadata information: ", f)
	}

	// Allocate a buffer with `chunkSize` as the capacity
	// and length (making a 0 array of the size of `chunkSize`)
	buf = make([]byte, c.chunkSize)
	for writing {
		// put as many bytes as `chunkSize` into the
		// buf array.
		n, err = file.Read(buf)
		if err != nil {
			// ... if `eof` --> `writing=false`...
			if err == io.EOF {
				writing = false
				err = nil
				continue
			}
			logger.Fatalf("errored while copying from file to buf: ", err)
		}

		req := &lcmservice.InstantiateRequest {
			Data: &lcmservice.InstantiateRequest_Package {
				Package: buf[:n],
			},
		}

		err = stream.Send(req)

		if err != nil {
			logger.Fatalf("failed to send chunk via stream: ", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		logger.Fatalf("failed to receive upstream status response: ", err)
		return "", "", err
	}
	log.Printf("response", res)
	return res.WorkloadId, res.Status, err
}

func (c *ClientGRPC) Query(ctx context.Context, hostIP string, workloadId string) (status string) {

	req := &lcmservice.QueryRequest{
		HostIp: hostIP,
		WorkloadId: workloadId,
	}
	resp, _ := c.client.Query(ctx, req)
	return resp.Status
}

func (c *ClientGRPC) Terminate(ctx context.Context, hostIP string, workloadId string) (status string) {

	req := &lcmservice.TerminateRequest{
		HostIp: hostIP,
		WorkloadId: workloadId,
	}
	resp, _ := c.client.Terminate(ctx, req)
	return resp.Status
}

func (c *ClientGRPC) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

