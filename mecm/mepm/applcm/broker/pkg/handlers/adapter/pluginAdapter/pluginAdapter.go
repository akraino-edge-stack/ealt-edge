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
	"github.com/sirupsen/logrus"
	"time"
)

const (
	chunkSize = 1024
	rootCertificate  = ""
)

// Plugin adapter which decides a specific client based on plugin info
// TODO PluginInfo to have other information about plugins to find the client and implementation to handle accordingly.
type PluginAdapter struct {
	pluginInfo string
	logger *logrus.Logger
}

// Constructor of PluginAdapter
func NewPluginAdapter(pluginInfo string, logger *logrus.Logger) *PluginAdapter {
	return &PluginAdapter{pluginInfo: pluginInfo, logger: logger}
}

// Instantiate application
func (c *PluginAdapter) Instantiate(pluginInfo string, host string, deployArtifact string) (workloadId string, error error, status string) {
	c.logger.Infof("Instantation started")
	clientConfig := plugin.ClientGRPCConfig{Address: pluginInfo, ChunkSize: chunkSize, RootCertificate: rootCertificate, Logger: c.logger}
	var client, err = plugin.NewClientGRPC(clientConfig)
	if err != nil {
		c.logger.Errorf("failed to create client: %v", err)
		return "", err, "Failure"
	}

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()

	//Instantiate
	workloadId, status, err = client.Instantiate(ctx, deployArtifact, host)
	if err != nil {
		c.logger.Errorf("server failed to respond %s", err.Error())
		return "", err, "Failure"
	}
	c.logger.Infof("Instantiation completed with workloadId %s, status: %s ", workloadId, status)
	return workloadId, nil, status
}

// Query application
func (c *PluginAdapter) Query(pluginInfo string, host string, workloadId string) (status string, error error) {
	c.logger.Infof("Query started")
	clientConfig := plugin.ClientGRPCConfig{Address: pluginInfo, ChunkSize: chunkSize, RootCertificate: rootCertificate}
	var client, err = plugin.NewClientGRPC(clientConfig)
	if err != nil {
		c.logger.Errorf("failed to create client: %v", err)
		return "", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Query
	status, err = client.Query(ctx, host, workloadId)
	if err != nil {
		c.logger.Errorf("failed to query: %v", err)
		return "", err
	}
	c.logger.Infof("query status: ", status)
	return status, nil
}

// Terminate application
func (c *PluginAdapter) Terminate(pluginInfo string, host string, workloadId string) (status string, error error) {
	c.logger.Infof("Terminate started")
	clientConfig := plugin.ClientGRPCConfig{Address: pluginInfo, ChunkSize: chunkSize, RootCertificate: rootCertificate}
	var client, err = plugin.NewClientGRPC(clientConfig)
	if err != nil {
		c.logger.Errorf("failed to create client: %v", err)
		return "Failure", err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	//Terminate
	status, err = client.Terminate(ctx, host, workloadId)

	if err != nil {
		c.logger.Errorf("failed to instantiate: %v", err)
		return "Failure", err
	}

	c.logger.Infof("termination success with status: ", status)
	return status, nil
}
