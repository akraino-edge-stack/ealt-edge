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

	initCmds "ealt/cmd/init"

	"github.com/spf13/cobra"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "This command is used to install various components of EALT",
	Long: `Command has multiple options to handle installation of each components.
	Options :
	"all" : To install the complete EALT Environment.
	"mecm": To install the MECM - Controller Node
	"mep" : To install the MEP - Edge Node`,
}

func init() {
	//Adding the various sub-commands of init command
	//Adding all subcommand to init 
	initCmd.AddCommand(initCmds.NewAllCommand())
	//ealt init infra
	initCmd.AddCommand(initCmds.NewInfraCommand())
	//ealt init manager
	initCmd.AddCommand(initCmds.NewMecmCommand())
	//ealt init edge
	initCmd.AddCommand(initCmds.NewEdgeCommand())
	//ealt init k8s
	initCmd.AddCommand(initCmds.NewK8SCommand())
	//ealt init k3s
	initCmd.AddCommand(initCmds.NewK3SCommand())

	//Add init subcommand to root command.
	rootCmd.AddCommand(initCmd)

}
