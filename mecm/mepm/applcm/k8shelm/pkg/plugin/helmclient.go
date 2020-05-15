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
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/kube"
)

// Variables to be defined in deployment file
var (
	chartPath        = os.Getenv("CHART_PATH")
	kubeconfigPath   = os.Getenv("KUBECONFIG_DIR_PATH")
	releaseNamespace = os.Getenv("RELEASE_NAMESPACE")
)

// Helm client
type HelmClient struct {
	hostIP     string
	kubeconfig string
	logger     *logrus.Logger
}

// Constructor of helm client for a given host IP
func NewHelmClient(hostIP string, logger *logrus.Logger) (*HelmClient, error) {
	// Kubeconfig file will be picked based on host IP and will be check for existence
	exists, err := fileExists(kubeconfigPath + hostIP)
	if exists {
		return &HelmClient{hostIP: hostIP, kubeconfig: kubeconfigPath + hostIP, logger: logger}, nil
	} else {
		logger.Errorf("No file exist with name: %s. Err: %s", kubeconfigPath+hostIP)
		return nil, err
	}
}

// Install a given helm chart
func (hc *HelmClient) installChart(helmPkg bytes.Buffer) (string, error) {
	hc.logger.Debug("Inside helm client")

	// Create temporary file to hold helm chart
	file, err := os.Create(chartPath + "temp.tar.gz")
	if err != nil {
		hc.logger.Errorf("Unable to create file: %s. Err: %s", chartPath+"temp.tar.gz", err)
		return "", err
	}
	defer os.Remove(chartPath + "temp.tar.gz")

	// Write input bytes to temp file
	_, err = helmPkg.WriteTo(file)
	if err != nil {
		hc.logger.Errorf("Unable to write to file: %s. Err: %s", chartPath+"temp.tar.gz", err)
		return "", err
	}

	// Load the file to chart
	chart, err := loader.Load(chartPath + "temp.tar.gz")
	if err != nil {
		hc.logger.Errorf("Unable to load chart from file: %s. Err: %s", chartPath+"temp.tar.gz", err)
		return "", err
	}

	// Release name will be taken from the name in chart's metadata
	relName := chart.Metadata.Name

	// Initialize action config
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kube.GetConfig(hc.kubeconfig, "", releaseNamespace), releaseNamespace,
		os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
			fmt.Sprintf(format, v)
		}); err != nil {
		hc.logger.Errorf("Unable to initialize action config Err: %s", err)
		return "", err
	}

	// Prepare chart install action and install chart
	installer := action.NewInstall(actionConfig)
	installer.Namespace = releaseNamespace
	installer.ReleaseName = relName
	rel, err := installer.Run(chart, nil)
	if err != nil {
		hc.logger.Errorf("Unable to install chart with release name: %s. Err: %s", relName, err)
		return "", err
	}
	hc.logger.Infof("Successfully create chart with release name: %s", relName)
	return rel.Name, err
}

// Un-Install a given helm chart
func (hc *HelmClient) uninstallChart(relName string) error {
	// Prepare action config and uninstall chart
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kube.GetConfig(hc.kubeconfig, "", releaseNamespace), releaseNamespace,
		os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
			fmt.Sprintf(format, v)
		}); err != nil {
		hc.logger.Errorf("Unable to initialize action config Err: %s", err)
		return err
	}

	ui := action.NewUninstall(actionConfig)
	res, err := ui.Run(relName)
	if err != nil {
		hc.logger.Errorf("Unable to uninstall chart with release name: %s. Err: %s", relName, err)
		return err
	}
	hc.logger.Infof("Successfully uninstalled chart with release name: %s. Response Info: %s", res.Release.Name, res.Info)
	return nil
}

// Query a given chart
func (hc *HelmClient) queryChart(relName string) (string, error) {
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kube.GetConfig(hc.kubeconfig, "", releaseNamespace), releaseNamespace,
		os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
			fmt.Sprintf(format, v)
		}); err != nil {
		hc.logger.Errorf("Unable to initialize action config Err: %s", err)
		return "", err
	}
	s := action.NewStatus(actionConfig)
	res, err := s.Run(relName)
	if err != nil {
		hc.logger.Errorf("Unable to query chart with release name: %s. Err: %s", relName, err)
		return "", err
	}
	return res.Info.Status.String(), nil
}

// fileExists checks if a file exists and is not a directory before we
// try using it to prevent further errors.
func fileExists(filename string) (bool, error) {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, err
	}
	return !info.IsDir(), nil
}
