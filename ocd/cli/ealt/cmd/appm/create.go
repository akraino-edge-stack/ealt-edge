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
package appm

import (
	"ealt/cmd/adapter"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
func NewAppCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "To create the application on MEP Node",
		Long:  `To create the application on MEP Node`,
		RunE: func(cmd *cobra.Command, args []string) error {
			theFlags := []string{cmd.Flag("packagefile").Value.String()}
			err := adapter.BuilderRequest(theFlags, "NewAppCreateCommand")
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("packagefile", "f", "", "Application Package File to be onboarded to MEP")
	cmd.MarkFlagRequired("packagefile")
	return cmd
}
