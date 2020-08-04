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
	"fmt"
	"strings"

	setup "ealt/cmd/setup"

	"github.com/spf13/cobra"
)

func NewEdgeCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "edge",
		Short: "Install Edge Node",
		Long:  `Command to Install Edge Node only For Example : ealt init edge`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("Installation of Edge components")
			setupModeFlag := strings.ToLower(cmd.Flag("mode").Value.String())
			var err error
			if setupModeFlag == "dev" {
				err = setup.EaltInstall("edge")
			} else if setupModeFlag == "prod" {
				err = setup.EaltInstall("ssledge")
			}
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("mode", "m", "dev", "Deployment Mode")
	return cmd
}
