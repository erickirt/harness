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

package gitspace

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/harness/gitness/store"
	"github.com/harness/gitness/store/database/dbtx"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"

	"github.com/rs/zerolog/log"
)

func (c *Service) FindWithLatestInstanceWithSpacePath(
	ctx context.Context,
	spacePath string,
	identifier string,
) (*types.GitspaceConfig, error) {
	space, err := c.spaceFinder.FindByRef(ctx, spacePath)
	if err != nil {
		return nil, fmt.Errorf("failed to find space: %w", err)
	}
	return c.FindWithLatestInstance(ctx, space.ID, spacePath, identifier)
}

func (c *Service) FindWithLatestInstance(
	ctx context.Context,
	spaceID int64,
	spacePath string,
	identifier string,
) (*types.GitspaceConfig, error) {
	var gitspaceConfigResult *types.GitspaceConfig
	txErr := c.tx.WithTx(ctx, func(ctx context.Context) error {
		gitspaceConfig, err := c.gitspaceConfigStore.FindByIdentifier(ctx, spaceID, identifier)
		if err != nil {
			return fmt.Errorf("failed to find gitspace config: %w", err)
		}
		gitspaceConfigResult = gitspaceConfig
		if err = c.setInstance(ctx, gitspaceConfigResult); err != nil {
			return err
		}
		return nil
	}, dbtx.TxDefaultReadOnly)
	if txErr != nil {
		return nil, txErr
	}
	gitspaceConfigResult.BranchURL = c.GetBranchURL(ctx, gitspaceConfigResult)

	if spacePath == "" {
		space, err := c.spaceFinder.FindByID(ctx, spaceID)
		if err != nil {
			log.Warn().Err(err).Msgf("failed to find space path for id %d", spaceID)
		} else {
			spacePath = space.Path
		}
	}

	gitspaceConfigResult.SpacePath = spacePath
	return gitspaceConfigResult, nil
}

func (c *Service) setInstance(
	ctx context.Context,
	gitspaceConfig *types.GitspaceConfig,
) error {
	instance, err := c.gitspaceInstanceStore.FindLatestByGitspaceConfigID(ctx, gitspaceConfig.ID)
	if err != nil && !errors.Is(err, store.ErrResourceNotFound) {
		return err
	}
	if instance != nil {
		gitspaceConfig.GitspaceInstance = instance
		instance.SpacePath = gitspaceConfig.SpacePath
		gitspaceStateType, err := instance.GetGitspaceState()
		if err != nil {
			return err
		}
		gitspaceConfig.State = gitspaceStateType
	} else {
		gitspaceConfig.State = enum.GitspaceStateUninitialized
	}
	return nil
}

func (c *Service) FindWithLatestInstanceByID(
	ctx context.Context,
	id int64,
	includeDeleted bool,
) (*types.GitspaceConfig, error) {
	var gitspaceConfigResult *types.GitspaceConfig
	txErr := c.tx.WithTx(ctx, func(ctx context.Context) error {
		gitspaceConfig, err := c.gitspaceConfigStore.Find(ctx, id, includeDeleted)
		gitspaceConfigResult = gitspaceConfig
		if err != nil {
			return fmt.Errorf("failed to find gitspace config: %w", err)
		}
		space, err := c.spaceFinder.FindByID(ctx, gitspaceConfigResult.SpaceID)
		if err != nil {
			return fmt.Errorf("failed to find space: %w", err)
		}
		gitspaceConfigResult.SpacePath = space.Path
		return c.setInstance(ctx, gitspaceConfigResult)
	}, dbtx.TxDefaultReadOnly)
	if txErr != nil {
		return nil, txErr
	}
	return gitspaceConfigResult, nil
}

func (c *Service) FindAll(
	ctx context.Context,
	ids []int64,
) ([]*types.GitspaceConfig, error) {
	var gitspaceConfigResult []*types.GitspaceConfig
	txErr := c.tx.WithTx(ctx, func(ctx context.Context) error {
		// TODO join and set gitspace instance, space from cache
		gitspaceConfigs, err := c.gitspaceConfigStore.FindAll(ctx, ids)
		if err != nil {
			return fmt.Errorf("failed to find gitspace config: %w", err)
		}
		for _, gitspaceConfig := range gitspaceConfigs {
			// FindByRef method is backed by cache as opposed to Find
			space, err := c.spaceFinder.FindByRef(ctx, strconv.FormatInt(gitspaceConfig.SpaceID, 10))
			if err != nil {
				return fmt.Errorf("failed to find space: %w", err)
			}
			gitspaceConfig.SpacePath = space.Path
			gitspaceConfigResult = append(gitspaceConfigResult, gitspaceConfig)
		}
		return nil
	}, dbtx.TxDefaultReadOnly)
	if txErr != nil {
		return nil, txErr
	}
	return gitspaceConfigResult, nil
}

func (c *Service) FindAllByIdentifier(
	ctx context.Context,
	spaceID int64,
	identifiers []string,
) ([]types.GitspaceConfig, error) {
	var gitspaceConfigResult []types.GitspaceConfig
	txErr := c.tx.WithTx(ctx, func(ctx context.Context) error {
		gitspaceConfigs, err := c.gitspaceConfigStore.FindAllByIdentifier(ctx, spaceID, identifiers)
		if err != nil {
			return fmt.Errorf("failed to find gitspace config: %w", err)
		}

		gitspaceConfigResult = gitspaceConfigs
		return nil
	}, dbtx.TxDefaultReadOnly)
	if txErr != nil {
		return nil, txErr
	}
	return gitspaceConfigResult, nil
}

func (c *Service) FindInstanceByIdentifier(
	ctx context.Context,
	identifier string,
) (*types.GitspaceInstance, error) {
	var gitspaceInstanceResult *types.GitspaceInstance
	txErr := c.tx.WithTx(ctx, func(ctx context.Context) error {
		gitspaceInstance, err := c.gitspaceInstanceStore.FindByIdentifier(ctx, identifier)
		if err != nil {
			return fmt.Errorf("failed to find gitspace instance: %w", err)
		}
		gitspaceInstanceResult = gitspaceInstance
		return nil
	}, dbtx.TxDefaultReadOnly)
	if txErr != nil {
		return nil, txErr
	}
	return gitspaceInstanceResult, nil
}
