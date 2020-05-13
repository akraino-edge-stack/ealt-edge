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
	"broker/pkg/util"
	"github.com/sirupsen/logrus"
	"os"
)

const (
	logFile = "/home/root1/code/ealt-edge/mecm/mepm/applcm/k8shelm/pkg/plugin/logfile"
	loggerLevel = logrus.InfoLevel
)

func main() {
	// Prepare logger
	file, err := os.Create(logFile)
	if err != nil {
		logrus.Fatal(err)
	}
	defer file.Close()

	var logger = util.GetLogger(logFile, loggerLevel, file)

	handler := &handlers.Handlers{}
	handler.Initialize(logger)
	//handler.Run("127.0.0.1:8081")
	handler.Run("0.0.0.0:8081")
}
