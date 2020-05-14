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

package main

import (
	"github.com/sirupsen/logrus"
	"k8shelm/pkg/plugin"
	"os"
	"strconv"
)

// Variables to be defined in deployment file
var (
	serverPort = os.Getenv("HELM_PLUGIN_PORT")
	logFile = os.Getenv("LOGFILE_PATH")
	loggerLevel = os.Getenv("LOGGER_LEVEL")
	certificate = os.Getenv("CERTIFICATE_PATH")
	key = os.Getenv("KEY_PATH")
)

func main() {
	// Prepare logger
	file, err := os.Create(logFile)
	if err != nil {
		logrus.Fatal(err)
	}
	defer file.Close()

	level, err := logrus.ParseLevel(loggerLevel)
	var logger = plugin.GetLogger(logFile, level, file)

	// Create GRPC server
	sp, err := strconv.Atoi(serverPort)
	serverConfig := plugin.ServerGRPCConfig{Certificate: certificate, Port:sp, Key:key, Logger:logger}
	server := plugin.NewServerGRPC(serverConfig)

	// Start listening
	err = server.Listen()
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
}
