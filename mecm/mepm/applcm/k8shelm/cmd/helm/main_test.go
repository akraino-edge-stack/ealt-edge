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
	"strconv"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	file, err := os.Create(logFile)
	assert.NotNil(t, file, "File should not nil")
	if err != nil {
		t.Errorf("Expected value, received %v", err)
	}
}

func TestParseLevel(t *testing.T) {
	level, err := logrus.ParseLevel(loggerLevel)
	assert.NotNil(t, level, "Level should not nil")
	if err != nil {
		t.Errorf("Expected value, received %v", err)
	}
}

func TestAtoi(t *testing.T) {
	sp, err := strconv.Atoi(serverPort)
	assert.Equal(t, sp, sp, "Both should equal")
	if err != nil {
		t.Errorf("Expected value, received %v", err)
	}
}

func TestServerGRPCConfig(t *testing.T) {
	//var logger = plugin.GetLogger(logFile, _, _)
	sp, err := strconv.Atoi(serverPort)
	serverConfig := plugin.ServerGRPCConfig{Certificate: certificate, Port: sp, Key: key, Logger: nil}
	assert.NotNil(t, serverConfig, "The server should not nil")
	assert.Equal(t, sp, serverConfig.Port, "The port is not matching")
	log.Print(serverConfig)
	//log.Print(t)
	if err != nil {
		t.Errorf("Expected value, received %v", serverConfig)
	}
}
