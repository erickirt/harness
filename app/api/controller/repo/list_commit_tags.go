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

package repo

import (
	"context"

	"github.com/harness/gitness/app/api/controller"
	"github.com/harness/gitness/app/auth"
	"github.com/harness/gitness/git"
	"github.com/harness/gitness/git/sha"
	"github.com/harness/gitness/types"
	"github.com/harness/gitness/types/enum"
)

type CommitTag struct {
	Name        string            `json:"name"`
	SHA         sha.SHA           `json:"sha"`
	IsAnnotated bool              `json:"is_annotated"`
	Title       string            `json:"title,omitempty"`
	Message     string            `json:"message,omitempty"`
	Tagger      *types.Signature  `json:"tagger,omitempty"`
	SignedData  *types.SignedData `json:"-"`
	Commit      *types.Commit     `json:"commit,omitempty"`
}

// ListCommitTags lists the commit tags of a repo.
func (c *Controller) ListCommitTags(ctx context.Context,
	session *auth.Session,
	repoRef string,
	includeCommit bool,
	filter *types.TagFilter,
) ([]CommitTag, error) {
	repo, err := c.getRepoCheckAccess(ctx, session, repoRef, enum.PermissionRepoView)
	if err != nil {
		return nil, err
	}

	rpcOut, err := c.git.ListCommitTags(ctx, &git.ListCommitTagsParams{
		ReadParams:    git.CreateReadParams(repo),
		IncludeCommit: includeCommit,
		Query:         filter.Query,
		Sort:          mapToRPCTagSortOption(filter.Sort),
		Order:         mapToRPCSortOrder(filter.Order),
		Page:          int32(filter.Page), //nolint:gosec
		PageSize:      int32(filter.Size), //nolint:gosec
	})
	if err != nil {
		return nil, err
	}

	tags := make([]CommitTag, len(rpcOut.Tags))
	for i := range rpcOut.Tags {
		tags[i] = mapCommitTag(rpcOut.Tags[i])
	}

	return tags, nil
}

func mapToRPCTagSortOption(o enum.TagSortOption) git.TagSortOption {
	switch o {
	case enum.TagSortOptionDate:
		return git.TagSortOptionDate
	case enum.TagSortOptionName:
		return git.TagSortOptionName
	case enum.TagSortOptionDefault:
		return git.TagSortOptionDefault
	default:
		// no need to error out - just use default for sorting
		return git.TagSortOptionDefault
	}
}

func mapCommitTag(t git.CommitTag) CommitTag {
	var tagger *types.Signature
	if t.Tagger != nil {
		tagger = &types.Signature{}
		*tagger = controller.MapSignature(*t.Tagger)
	}

	return CommitTag{
		Name:        t.Name,
		SHA:         t.SHA,
		IsAnnotated: t.IsAnnotated,
		Title:       t.Title,
		Message:     t.Message,
		Tagger:      tagger,
		SignedData:  (*types.SignedData)(t.SignedData),
		Commit:      controller.MapCommit(t.Commit),
	}
}
