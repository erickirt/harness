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

package response

type StartResponse struct {
	ContainerID      string         `json:"container_id"`
	Status           Status         `json:"status"`
	ErrMessage       string         `json:"err_message"`
	ContainerName    string         `json:"container_name"`
	PublishedPorts   map[int]string `json:"published_ports"`
	AbsoluteRepoPath string         `json:"absolute_repo_path"`
	RemoteUser       string         `json:"remote_user"`
}
