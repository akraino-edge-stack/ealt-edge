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
package init

import (
	setup "ealt/cmd/setup"

	"github.com/spf13/cobra"
)

// allCmd represents the all command
func NewAllCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "all",
		Short: "Install Complete EALT Deployment Environment",
		Long:  `Install Complete EALT Deployment Environment`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// setupModeFlag := strings.ToLower(cmd.Flag("mode").Value.String())
			var err error
			err = setup.EaltInstall("all")
			// if setupModeFlag == "dev" {
			// 	err = setup.EaltInstall("all")
			// } else if setupModeFlag == "prod" {
			// 	err = setup.EaltInstall("secure")
			// }
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("mode", "m", "dev", "Deployment Mode")
	//cmd.MarkFlagRequired("mode")

	return cmd
}
