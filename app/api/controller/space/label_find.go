// Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package space

import (
	"context"
	"fmt"

	"github.com/harness/gitness/app/auth"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"
)

// FindLabel finds a label for the specified space.
func (c *Controller) FindLabel(
	ctx context.Context,
	session *auth.Session,
	spaceRef string,
	key string,
	includeValues bool,
) (*types.LabelWithValues, error) {
	space, err := c.getSpaceCheckAuth(ctx, session, spaceRef, enum.PermissionSpaceView)
	if err != nil {
		return nil, fmt.Errorf("failed to acquire access to space: %w", err)
	}

	var labelWithValues *types.LabelWithValues
	if includeValues {
		labelWithValues, err = c.labelSvc.FindWithValues(ctx, &space.ID, nil, key)
		if err != nil {
			return nil, fmt.Errorf("failed to find space label with values: %w", err)
		}
	} else {
		label, err := c.labelSvc.Find(ctx, &space.ID, nil, key)
		if err != nil {
			return nil, fmt.Errorf("failed to find space label: %w", err)
		}
		labelWithValues = &types.LabelWithValues{
			Label: *label,
		}
	}

	return labelWithValues, nil
}
