package plugin

import (
	"bytes"
	"fmt"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart/loader"
	"helm.sh/helm/v3/pkg/kube"
	"log"
	"os"
)

const releaseNamespace  = "default"
const chartPath  = "/go/release/charts/"
const kubeconfigPath  = "/go/release/kubeconfig/"

//const chartPath  = "/home/root1/code/mecm/mepm/applcm/k8shelm/pkg/plugin/"
//const kubeconfigPath  = "/home/root1/"

func installChart(helmPkg bytes.Buffer, hostIP string) string {
	logger := log.New(os.Stdout, "helmplugin ", log.LstdFlags|log.Lshortfile)
	logger.Println("Inside helm client")

	file, err := os.Create(chartPath + "temp.tar.gz")
	if err != nil {
		logger.Printf("unable to create file")
	}

	_, err = helmPkg.WriteTo(file)
	if err != nil {
		logger.Printf("uanble to write to file")
	}

	chart, err := loader.Load(chartPath + "temp.tar.gz")
	if err != nil {

		panic(err)
	}

	releaseName := chart.Metadata.Name

	kubeconfig := kubeconfigPath + hostIP

	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kube.GetConfig(kubeconfig, "", releaseNamespace), releaseNamespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Sprintf(format, v)
	}); err != nil {
		panic(err)
	}

	iCli := action.NewInstall(actionConfig)
	iCli.Namespace = releaseNamespace
	iCli.ReleaseName = releaseName
	rel, err := iCli.Run(chart, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully installed release: ", rel.Name)
	return rel.Name
}

func uninstallChart(relName string, hostIP string) {
	kubeconfig := kubeconfigPath + hostIP
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kube.GetConfig(kubeconfig, "", releaseNamespace), releaseNamespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Sprintf(format, v)
	}); err != nil {
		panic(err)
	}
	iCli := action.NewUninstall(actionConfig)
	res, err := iCli.Run(relName);
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully uninstalled release: ", res.Info)
}

func queryChart(relName string, hostIP string) string  {
	kubeconfig := kubeconfigPath + hostIP
	actionConfig := new(action.Configuration)
	if err := actionConfig.Init(kube.GetConfig(kubeconfig, "", releaseNamespace), releaseNamespace, os.Getenv("HELM_DRIVER"), func(format string, v ...interface{}) {
		fmt.Sprintf(format, v)
	}); err != nil {
		panic(err)
	}
	iCli := action.NewStatus(actionConfig)
	res, err := iCli.Run(relName)
	if err != nil {
		panic(err)
	}
	return res.Info.Status.String()
}

