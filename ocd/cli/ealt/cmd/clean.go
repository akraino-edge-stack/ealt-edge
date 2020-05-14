/*
Copyright 2020 Huawei Technologies Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	cleancmds "ealt/cmd/clean"
	"github.com/spf13/cobra"
)

// cleanCmd represents the clean command
var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "To uninstall ealt environment or specific component or node.",
	Long:  `To uninstall ealt environment or specific component or node.`,
}

func init() {
	cleanCmd.AddCommand(cleancmds.NewAllCommand())
	//ealt init infra
	cleanCmd.AddCommand(cleancmds.NewInfraCommand())
	//ealt init manager
	cleanCmd.AddCommand(cleancmds.NewMecmCommand())
	//ealt init edge
	cleanCmd.AddCommand(cleancmds.NewEdgeCommand())
	//ealt init k8s
	cleanCmd.AddCommand(cleancmds.NewK8SCommand())
	//ealt init k3s
	cleanCmd.AddCommand(cleancmds.NewK3SCommand())

	//Add init subcommand to root command.
	rootCmd.AddCommand(cleanCmd)

}
