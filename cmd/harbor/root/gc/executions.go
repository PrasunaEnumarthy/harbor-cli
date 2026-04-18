// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package gc

import (
	"fmt"

	"github.com/goharbor/harbor-cli/pkg/api"
	"github.com/goharbor/harbor-cli/pkg/utils"
	view "github.com/goharbor/harbor-cli/pkg/views/gc/list"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GcExecutionsCommand() *cobra.Command {
	var opts api.ListFlags

	cmd := &cobra.Command{
		Use:   "executions",
		Short: "List GC executions",
		Long:  `List historical Garbage Collection executions in Harbor with pagination.`,
		Args:  cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			logrus.Info("Retrieving Garbage Collection execution history")

			// Validation!
			if opts.Page < 1 {
				return fmt.Errorf("page number must be greater than or equal to 1")
			}
			if opts.PageSize < 0 {
				return fmt.Errorf("page size must be greater than or equal to 0")
			}
			if opts.PageSize > 100 {
				return fmt.Errorf("page size should be less than or equal to 100")
			}

			executions, err := api.ListGCExecutions(opts)
			if err != nil {
				return fmt.Errorf("failed to retrieve GC executions: %v", utils.ParseHarborErrorMsg(err))
			}

			FormatFlag := viper.GetString("output-format")
			if FormatFlag != "" {
				err = utils.PrintFormat(executions, FormatFlag)
				if err != nil {
					return err
				}
			} else {
				view.ListGCExecutions(executions)
			}
			return nil
		},
	}

	flags := cmd.Flags()
	flags.Int64Var(&opts.Page, "page", 1, "Page number for pagination")
	flags.Int64Var(&opts.PageSize, "page-size", 10, "Size of elements in a page")

	return cmd
}
