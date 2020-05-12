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
package pluginAdapter

import (
	"broker/pkg/plugin"
	"context"
	"log"
	"os"
	"time"
)

func Instantiate(pluginInfo string, host string, deployArtifact string) (workloadId string, error error) {
	logger := log.New(os.Stdout, "broker ", log.LstdFlags|log.Lshortfile)
	clientConfig := plugin.ClientGRPCConfig{Address: pluginInfo, ChunkSize: 1024, RootCertificate: ""}
	var client, err = plugin.NewClientGRPC(clientConfig)
	if err != nil {
		logger.Fatalf("failed to create client: %v", err)
		return "", err
	}
	log.Printf("pluginInfo: ", pluginInfo)
	log.Printf("host: ", host)
	log.Printf("deployArtifact: ", deployArtifact)
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	//Instantiate
	workloadId, status, err := client.Instantiate(ctx, deployArtifact, host)
	if err != nil {
		logger.Println("server failed to respond %s", err.Error())
	} else {
		logger.Println(workloadId, status)
		return workloadId, nil
	}
	return "", err
}

func Query(pluginInfo string, host string, workloadId string) (status string, error error) {
	logger := log.New(os.Stdout, "broker ", log.LstdFlags|log.Lshortfile)
	clientConfig := plugin.ClientGRPCConfig{Address: pluginInfo, ChunkSize: 1024, RootCertificate: ""}
	var client, err = plugin.NewClientGRPC(clientConfig)
	if err != nil {
		logger.Fatalf("failed to create client: %v", err)
		return "", err
	}
	log.Printf("pluginInfo: ", pluginInfo)
	log.Printf("host: ", host)
	log.Printf("workloadId: ", workloadId)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Query
	stats := client.Query(ctx, host, workloadId)
	if err != nil {
		logger.Fatalf("failed to instantiate: %v", err)
		return stats, err
	}
	logger.Println("query status: ", stats)
	return stats, nil
}

func Terminate(pluginInfo string, host string, workloadId string) (status string, error error) {
	logger := log.New(os.Stdout, "broker ", log.LstdFlags|log.Lshortfile)
	clientConfig := plugin.ClientGRPCConfig{Address: pluginInfo, ChunkSize: 1024, RootCertificate: ""}
	var client, err = plugin.NewClientGRPC(clientConfig)
	if err != nil {
		logger.Fatalf("failed to create client: %v", err)
		return
	}
	log.Printf("pluginInfo: ", pluginInfo)
	log.Printf("host: ", host)
	log.Printf("workloadId: ", workloadId)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Terminate
	ts := client.Terminate(ctx, host, workloadId)
	if err != nil {
		logger.Fatalf("failed to instantiate: %v", err)
		return ts, err
	}

	logger.Println("termination status: ", ts)
	return ts, nil
}
