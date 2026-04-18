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
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func GcScheduleCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "schedule",
		Short: "View the GC schedule",
		Long:  `Display the current Garbage Collection schedule configuration in Harbor.`,
		Args:  cobra.MaximumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			logrus.Info("Retrieving Garbage Collection schedule configuration")

			schedule, err := api.GetGCSchedule()

			// Flawless error handling, exactly like we planned!
			if err != nil {
				return fmt.Errorf("failed to retrieve GC schedule: %v", utils.ParseHarborErrorMsg(err))
			}

			// Automatically format the output safely for now
			FormatFlag := viper.GetString("output-format")
			if FormatFlag == "" {
				FormatFlag = "json"
			}

			err = utils.PrintFormat(schedule, FormatFlag)
			if err != nil {
				return err
			}

			return nil
		},
	}

	return cmd
}
