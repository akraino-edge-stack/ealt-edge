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
	"broker/pkg/handlers"
	"log"
	"os"
)
/*
var (
	GcukCertFile    = os.Getenv("GCUK_CERT_FILE")
	GcukKeyFile     = os.Getenv("GCUK_KEY_FILE")
	GcukServiceAddr = os.Getenv("GCUK_SERVICE_ADDR")
)*/

func main() {
	logger := log.New(os.Stdout, "broker ", log.LstdFlags|log.Lshortfile)

	handler := &handlers.Handlers{}
	handler.Initialize(logger)
	//handler.Run("127.0.0.1:8081")
	handler.Run("0.0.0.0:8081")
}
