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
package api

import (
	"github.com/goharbor/go-client/pkg/sdk/v2.0/client/gc"
	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
	"github.com/goharbor/harbor-cli/pkg/utils"
)

// GetGCSchedule fetches the current Garbage Collection schedule from Harbor
func GetGCSchedule() (*models.GCHistory, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	response, err := client.GC.GetGCSchedule(ctx, &gc.GetGCScheduleParams{})
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

// ListGCExecutions fetches the history of Garbage Collection executions
func ListGCExecutions(opts ...ListFlags) ([]*models.GCHistory, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return nil, err
	}

	var listFlags ListFlags
	if len(opts) > 0 {
		listFlags = opts[0]
	}

	response, err := client.GC.GetGCHistory(ctx, &gc.GetGCHistoryParams{
		Page:     &listFlags.Page,
		PageSize: &listFlags.PageSize,
	})
	if err != nil {
		return nil, err
	}

	return response.Payload, nil
}

// GetGCLog fetches the raw text logs for a specific garbage collection execution
func GetGCLog(id int64) (string, error) {
	ctx, client, err := utils.ContextWithClient()
	if err != nil {
		return "", err
	}

	response, err := client.GC.GetGCLog(ctx, &gc.GetGCLogParams{
		GCID: id,
	})
	if err != nil {
		return "", err
	}

	return response.Payload, nil
}
