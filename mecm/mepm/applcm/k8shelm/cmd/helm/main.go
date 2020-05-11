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
	"k8shelm/pkg/plugin"
	"log"
	"os"
)

func main() {
	logger := log.New(os.Stdout, "helm ", log.LstdFlags|log.Lshortfile)
	serverConfig := plugin.ServerGRPCConfig{Certificate:"", Port:50051, Key:""}
	server, err := plugin.NewServerGRPC(serverConfig)
	if err != nil {
		logger.Fatalf("failed to create server: %v", err)
	}
	error := server.Listen()
	if error != nil {
		logger.Fatalf("failed to listen: %v", err)
	}
}
