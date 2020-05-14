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
package applcmpkg

import (
	"ealt/cmd/adapter"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command
func NewApplcmCreateCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "create",
		Short: "To create the application instance on MEP Node",
		Long:  `To create the application instance on MEP Node`,
		RunE: func(cmd *cobra.Command, args []string) error {
			theFlags := []string{cmd.Flag("appdid").Value.String(),
				cmd.Flag("name").Value.String(),
				cmd.Flag("desc").Value.String()}
			err := adapter.BuilderRequest(theFlags, "NewApplcmCreateCommand")
			if err != nil {
				return err
			}
			return nil
		},
	}
	cmd.Flags().StringP("appdid", "i", "", "Application  ID to be created MEP")
	cmd.Flags().StringP("name", "n", "", "Application Instance Name")
	cmd.Flags().StringP("desc", "d", "", "Application Instance Descriptor")
	cmd.MarkFlagRequired("appdid")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("desc")
	return cmd
}
