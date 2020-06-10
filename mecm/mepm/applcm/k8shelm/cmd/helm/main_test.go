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
