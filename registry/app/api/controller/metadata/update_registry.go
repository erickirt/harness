//  Copyright 2023 Harness, Inc.
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

package metadata

import (
	"context"
	"fmt"
	"net/http"

	apiauth "github.com/harness/gitness/app/api/auth"
	"github.com/harness/gitness/app/api/request"
	"github.com/harness/gitness/app/auth"
	"github.com/harness/gitness/audit"
	"github.com/harness/gitness/registry/app/api/openapi/contracts/artifact"
	"github.com/harness/gitness/registry/types"
	types2 "github.com/harness/gitness/types"
	gitnessenum "github.com/harness/gitness/types/enum"

	"github.com/rs/zerolog/log"
)

func (c *APIController) ModifyRegistry(
	ctx context.Context,
	r artifact.ModifyRegistryRequestObject,
) (artifact.ModifyRegistryResponseObject, error) {
	regInfo, err := c.RegistryMetadataHelper.GetRegistryRequestBaseInfo(ctx, "", string(r.RegistryRef))
	if err != nil {
		return artifact.ModifyRegistry400JSONResponse{
			BadRequestJSONResponse: artifact.BadRequestJSONResponse(
				*GetErrorResponse(http.StatusBadRequest, err.Error()),
			),
		}, err
	}
	space, err := c.SpaceFinder.FindByRef(ctx, regInfo.ParentRef)
	if err != nil {
		return artifact.ModifyRegistry400JSONResponse{
			BadRequestJSONResponse: artifact.BadRequestJSONResponse(
				*GetErrorResponse(http.StatusBadRequest, err.Error()),
			),
		}, err
	}

	session, _ := request.AuthSessionFrom(ctx)
	permissionChecks := c.RegistryMetadataHelper.GetPermissionChecks(space,
		regInfo.RegistryIdentifier, gitnessenum.PermissionRegistryEdit)
	if err = apiauth.CheckRegistry(
		ctx,
		c.Authorizer,
		session,
		permissionChecks...,
	); err != nil {
		return artifact.ModifyRegistry403JSONResponse{
			UnauthorizedJSONResponse: artifact.UnauthorizedJSONResponse(
				*GetErrorResponse(http.StatusForbidden, err.Error()),
			),
		}, err
	}

	repoEntity, err := c.RegistryRepository.GetByParentIDAndName(ctx, regInfo.ParentID, regInfo.RegistryIdentifier)
	if err != nil {
		return throwModifyRegistry500Error(err), err
	}

	if string(repoEntity.Type) == string(artifact.RegistryTypeVIRTUAL) {
		return c.updateVirtualRegistry(ctx, r, repoEntity, err, regInfo, session)
	}
	upstreamproxyEntity, err := c.UpstreamProxyStore.GetByRegistryIdentifier(
		ctx, regInfo.ParentID,
		regInfo.RegistryIdentifier,
	)
	if len(upstreamproxyEntity.RepoKey) == 0 {
		return artifact.ModifyRegistry404JSONResponse{
			NotFoundJSONResponse: artifact.NotFoundJSONResponse(
				*GetErrorResponse(http.StatusNotFound, "registry doesn't exist with this key"),
			),
		}, nil
	}
	if err != nil {
		return throwModifyRegistry500Error(err), err
	}
	registry, upstreamproxy, err := c.UpdateUpstreamProxyEntity(
		ctx,
		artifact.RegistryRequest(*r.Body),
		regInfo.ParentID, regInfo.RootIdentifierID, upstreamproxyEntity,
	)
	registry.ID = repoEntity.ID
	upstreamproxy.ID = upstreamproxyEntity.ID
	upstreamproxy.RegistryID = repoEntity.ID
	if err != nil {
		return throwModifyRegistry500Error(err), err
	}
	err = c.tx.WithTx(
		ctx, func(ctx context.Context) error {
			err = c.updateRegistryWithAudit(ctx, repoEntity, registry, session.Principal, regInfo.ParentRef)

			if err != nil {
				return fmt.Errorf("failed to update registry: %w", err)
			}

			err = c.updateUpstreamProxyWithAudit(
				ctx, upstreamproxy, session.Principal, regInfo.ParentRef, registry.Name,
			)

			if err != nil {
				return fmt.Errorf("failed to update upstream proxy: %w", err)
			}
			return nil
		},
	)
	if err != nil {
		return throwModifyRegistry500Error(err), err
	}
	modifiedRepoEntity, err := c.UpstreamProxyStore.Get(ctx, upstreamproxyEntity.RegistryID)
	if err != nil {
		return throwModifyRegistry500Error(err), err
	}
	if registry.PackageType == artifact.PackageTypeRPM {
		c.PostProcessingReporter.BuildRegistryIndex(ctx, registry.ID, make([]types.SourceRef, 0))
	}
	return artifact.ModifyRegistry200JSONResponse{
		RegistryResponseJSONResponse: *CreateUpstreamProxyResponseJSONResponse(modifiedRepoEntity),
	}, nil
}

func (c *APIController) updateVirtualRegistry(
	ctx context.Context, r artifact.ModifyRegistryRequestObject, repoEntity *types.Registry, err error,
	regInfo *types.RegistryRequestBaseInfo, session *auth.Session,
) (artifact.ModifyRegistryResponseObject, error) {
	if len(repoEntity.Name) == 0 {
		return artifact.ModifyRegistry404JSONResponse{
			NotFoundJSONResponse: artifact.NotFoundJSONResponse(
				*GetErrorResponse(http.StatusNotFound, "registry doesn't exist with this key"),
			),
		}, nil
	}
	if err != nil {
		return throwModifyRegistry500Error(err), err
	}
	registry, err := UpdateRepoEntity(
		artifact.RegistryRequest(*r.Body),
		repoEntity.ParentID,
		repoEntity.RootParentID,
		repoEntity,
	)
	if err != nil {
		return artifact.ModifyRegistry400JSONResponse{
			BadRequestJSONResponse: artifact.BadRequestJSONResponse(
				*GetErrorResponse(http.StatusInternalServerError, err.Error()),
			),
		}, nil
	}
	err = c.setUpstreamProxyIDs(ctx, registry, artifact.RegistryRequest(*r.Body), regInfo.ParentID)
	if err != nil {
		return throwModifyRegistry500Error(err), nil
	}
	if registry.PackageType == artifact.PackageTypeRPM {
		c.PostProcessingReporter.BuildRegistryIndex(ctx, registry.ID, make([]types.SourceRef, 0))
	}
	err = c.updateRegistryWithAudit(ctx, repoEntity, registry, session.Principal, regInfo.ParentRef)

	if err != nil {
		return throwModifyRegistry500Error(err), nil
	}
	err = c.updateCleanupPolicy(ctx, r.Body, registry.ID)
	if err != nil {
		return throwModifyRegistry500Error(err), nil
	}
	modifiedRepoEntity, err := c.RegistryRepository.Get(ctx, registry.ID)
	if err != nil {
		return throwModifyRegistry500Error(err), nil
	}
	cleanupPolicies, err := c.CleanupPolicyStore.GetByRegistryID(ctx, repoEntity.ID)
	if err != nil {
		return throwModifyRegistry500Error(err), nil
	}
	return artifact.ModifyRegistry200JSONResponse{
		RegistryResponseJSONResponse: *CreateVirtualRepositoryResponse(
			modifiedRepoEntity,
			c.getUpstreamProxyKeys(ctx, modifiedRepoEntity.UpstreamProxies), cleanupPolicies,
			c.URLProvider.RegistryURL(ctx, regInfo.RootIdentifier, regInfo.RegistryIdentifier),
		),
	}, nil
}

func (c *APIController) updateUpstreamProxyWithAudit(
	ctx context.Context, upstreamProxy *types.UpstreamProxyConfig,
	principal types2.Principal, parentRef string, registryName string,
) error {
	existingUpstreamProxy, err := c.UpstreamProxyStore.Get(ctx, upstreamProxy.RegistryID)
	if err != nil {
		log.Ctx(ctx).Warn().Msgf(
			"failed to fig upstream proxy config for: %d",
			upstreamProxy.RegistryID,
		)
	}

	err = c.UpstreamProxyStore.Update(ctx, upstreamProxy)
	if err != nil {
		return err
	}
	if existingUpstreamProxy != nil {
		auditErr := c.AuditService.Log(
			ctx,
			principal,
			audit.NewResource(audit.ResourceTypeRegistryUpstreamProxy, registryName),
			audit.ActionUpdated,
			parentRef,
			audit.WithOldObject(
				audit.RegistryUpstreamProxyConfigObject{
					ID:         existingUpstreamProxy.ID,
					RegistryID: existingUpstreamProxy.RegistryID,
					Source:     existingUpstreamProxy.Source,
					URL:        existingUpstreamProxy.RepoURL,
					AuthType:   existingUpstreamProxy.RepoAuthType,
					CreatedAt:  existingUpstreamProxy.CreatedAt,
					UpdatedAt:  existingUpstreamProxy.UpdatedAt,
					CreatedBy:  existingUpstreamProxy.CreatedBy,
					UpdatedBy:  existingUpstreamProxy.UpdatedBy,
				},
			),
			audit.WithNewObject(
				audit.RegistryUpstreamProxyConfigObject{
					ID:         upstreamProxy.ID,
					RegistryID: upstreamProxy.RegistryID,
					Source:     upstreamProxy.Source,
					URL:        upstreamProxy.URL,
					AuthType:   upstreamProxy.AuthType,
					CreatedAt:  upstreamProxy.CreatedAt,
					UpdatedAt:  upstreamProxy.UpdatedAt,
					CreatedBy:  upstreamProxy.CreatedBy,
					UpdatedBy:  upstreamProxy.UpdatedBy,
				},
			),
		)
		if auditErr != nil {
			log.Ctx(ctx).Warn().Msgf(
				"failed to insert audit log for update upstream proxy "+
					"config operation: %s", auditErr,
			)
		}
	}
	return err
}

func (c *APIController) updateRegistryWithAudit(
	ctx context.Context, oldRegistry *types.Registry,
	newRegistry *types.Registry, principal types2.Principal, parentRef string,
) error {
	err := c.RegFinder.Update(ctx, newRegistry)
	if err != nil {
		return err
	}
	auditErr := c.AuditService.Log(
		ctx,
		principal,
		audit.NewResource(audit.ResourceTypeRegistry, newRegistry.Name),
		audit.ActionUpdated,
		parentRef,
		audit.WithOldObject(oldRegistry),
		audit.WithNewObject(newRegistry),
	)
	if auditErr != nil {
		log.Ctx(ctx).Warn().Msgf("failed to insert audit log for update registry operation: %s", auditErr)
	}

	return err
}

func throwModifyRegistry500Error(err error) artifact.ModifyRegistry500JSONResponse {
	return artifact.ModifyRegistry500JSONResponse{
		InternalServerErrorJSONResponse: artifact.InternalServerErrorJSONResponse(
			*GetErrorResponse(http.StatusInternalServerError, err.Error()),
		),
	}
}

func (c *APIController) updateCleanupPolicy(
	ctx context.Context, config *artifact.ModifyRegistryJSONRequestBody, registryID int64,
) error {
	existingCleanupPolicies, err := c.CleanupPolicyStore.GetIDsByRegistryID(ctx, registryID)
	if err != nil {
		return err
	}
	currentCleanupPolicyEntities := CreateCleanupPolicyEntity(config, registryID)

	err = c.CleanupPolicyStore.ModifyCleanupPolicies(ctx, currentCleanupPolicyEntities, existingCleanupPolicies)

	return err
}

func UpdateRepoEntity(
	dto artifact.RegistryRequest,
	parentID int64,
	rootParentID int64,
	existingRepo *types.Registry,
) (*types.Registry, error) {
	allowedPattern, blockedPattern, description, labels := getRepoEntityFields(dto)
	e := ValidatePackageTypeChange(string(existingRepo.PackageType), string(dto.PackageType))
	if e != nil {
		return nil, e
	}
	e = ValidateRepoTypeChange(string(existingRepo.Type), string(dto.Config.Type))
	if e != nil {
		return nil, e
	}
	e = ValidateIdentifierChange(existingRepo.Name, dto.Identifier)
	if e != nil {
		return nil, e
	}
	entity := &types.Registry{
		Name:           dto.Identifier,
		ID:             existingRepo.ID,
		ParentID:       parentID,
		RootParentID:   rootParentID,
		Description:    description,
		AllowedPattern: allowedPattern,
		BlockedPattern: blockedPattern,
		PackageType:    existingRepo.PackageType,
		Type:           existingRepo.Type,
		Labels:         labels,
		CreatedAt:      existingRepo.CreatedAt,
	}
	return entity, nil
}

//nolint:gocognit,cyclop
func (c *APIController) UpdateUpstreamProxyEntity(
	ctx context.Context, dto artifact.RegistryRequest, parentID int64, rootParentID int64, u *types.UpstreamProxy,
) (*types.Registry, *types.UpstreamProxyConfig, error) {
	allowedPattern := []string{}
	if dto.AllowedPattern != nil {
		allowedPattern = *dto.AllowedPattern
	}
	blockedPattern := []string{}
	if dto.BlockedPattern != nil {
		blockedPattern = *dto.BlockedPattern
	}
	e := ValidatePackageTypeChange(string(u.PackageType), string(dto.PackageType))
	if e != nil {
		return nil, nil, e
	}
	e = ValidateIdentifierChange(u.RepoKey, dto.Identifier)
	if e != nil {
		return nil, nil, e
	}
	repoEntity := &types.Registry{
		ID:             u.RegistryID,
		Name:           dto.Identifier,
		ParentID:       parentID,
		RootParentID:   rootParentID,
		AllowedPattern: allowedPattern,
		BlockedPattern: blockedPattern,
		PackageType:    dto.PackageType,
		Type:           artifact.RegistryTypeUPSTREAM,
		CreatedAt:      u.CreatedAt,
	}
	config, _ := dto.Config.AsUpstreamConfig()
	CleanURLPath(config.Url)
	upstreamProxyConfigEntity := &types.UpstreamProxyConfig{
		URL:        *config.Url,
		AuthType:   string(config.AuthType),
		RegistryID: u.RegistryID,
		CreatedAt:  u.CreatedAt,
	}
	if config.Source != nil && len(string(*config.Source)) > 0 {
		err := ValidateUpstreamSource(string(*config.Source))
		if err != nil {
			return nil, nil, err
		}
		upstreamProxyConfigEntity.Source = string(*config.Source)
	}
	if string(artifact.UpstreamConfigSourceDockerhub) == string(*config.Source) {
		upstreamProxyConfigEntity.URL = ""
	}
	if u.ID != -1 {
		upstreamProxyConfigEntity.ID = u.ID
	}
	switch {
	case config.AuthType == artifact.AuthTypeUserPassword:
		res, err := config.Auth.AsUserPassword()
		if err != nil {
			return nil, nil, err
		}
		upstreamProxyConfigEntity.UserName = res.UserName
		if res.SecretIdentifier == nil {
			return nil, nil, fmt.Errorf("failed to create upstream proxy: secret_identifier missing")
		}

		if res.SecretSpacePath != nil && len(*res.SecretSpacePath) > 0 {
			upstreamProxyConfigEntity.SecretSpaceID, err = c.RegistryMetadataHelper.GetSecretSpaceID(ctx,
				res.SecretSpacePath)
			if err != nil {
				return nil, nil, err
			}
		} else if res.SecretSpaceId != nil {
			upstreamProxyConfigEntity.SecretSpaceID = *res.SecretSpaceId
		}
		upstreamProxyConfigEntity.SecretIdentifier = *res.SecretIdentifier
	case config.AuthType == artifact.AuthTypeAccessKeySecretKey:
		res, err := config.Auth.AsAccessKeySecretKey()
		if err != nil {
			return nil, nil, err
		}
		switch {
		case res.AccessKey != nil && len(*res.AccessKey) > 0:
			upstreamProxyConfigEntity.UserName = *res.AccessKey
		case res.AccessKeySecretIdentifier == nil:
			return nil, nil, fmt.Errorf("failed to create upstream proxy: access_key_secret_identifier missing")
		default:
			if res.AccessKeySecretSpacePath != nil && len(*res.AccessKeySecretSpacePath) > 0 {
				upstreamProxyConfigEntity.UserNameSecretSpaceID, err =
					c.RegistryMetadataHelper.GetSecretSpaceID(ctx, res.AccessKeySecretSpacePath)
				if err != nil {
					return nil, nil, err
				}
			} else if res.AccessKeySecretSpaceId != nil {
				upstreamProxyConfigEntity.UserNameSecretSpaceID = *res.AccessKeySecretSpaceId
			}
			upstreamProxyConfigEntity.UserNameSecretIdentifier = *res.AccessKeySecretIdentifier
		}

		if res.SecretKeySpacePath != nil && len(*res.SecretKeySpacePath) > 0 {
			upstreamProxyConfigEntity.SecretSpaceID, err =
				c.RegistryMetadataHelper.GetSecretSpaceID(ctx, res.SecretKeySpacePath)
			if err != nil {
				return nil, nil, err
			}
		} else if res.SecretKeySpaceId != nil {
			upstreamProxyConfigEntity.SecretSpaceID = *res.SecretKeySpaceId
		}
		upstreamProxyConfigEntity.SecretIdentifier = res.SecretKeyIdentifier
	default:
		upstreamProxyConfigEntity.UserName = ""
		upstreamProxyConfigEntity.SecretIdentifier = ""
		upstreamProxyConfigEntity.SecretSpaceID = 0
	}
	return repoEntity, upstreamProxyConfigEntity, nil
}
