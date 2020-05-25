/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	applcmCmds "ealt/cmd/applcmpkg"

	"github.com/spf13/cobra"
)

// applcmCmd represents the applcm command
var applcmCmd = &cobra.Command{
	Use:   "applcm",
	Short: "Commands to send request to the APPLCM Application",
	Long: `To manage the application running on the MEP Node, APPLCM exposes
	some API which can be used to manage the Applicaton running on the MEP Node
	The command have following options :
	1. Create
	2. Start
	3. Delete
	4. Stop.`,
}

func init() {

	applcmCmd.AddCommand(applcmCmds.NewApplcmCreateCommand())
	applcmCmd.AddCommand(applcmCmds.NewApplcmStartCommand())
	applcmCmd.AddCommand(applcmCmds.NewApplcmDeleteCommand())
	applcmCmd.AddCommand(applcmCmds.NewApplcmTerminateCommand())

	rootCmd.AddCommand(applcmCmd)

}
