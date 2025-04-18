/*
 * Copyright 2024 Harness, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import type { CellProps, Renderer } from 'react-table'
import type { StringsMap } from '@ar/strings/types'

export interface VersionListExpandedColumnProps {
  expandedRows: Set<string>
  setExpandedRows: React.Dispatch<React.SetStateAction<Set<string>>>
}

export enum VersionListColumnEnum {
  Name = 'Name',
  Size = 'Size',
  DownloadCount = 'DownloadCount',
  FileCount = 'FileCount',
  LastModified = 'LastModified',
  PullCommand = 'PullCommand',
  Actions = 'Actions'
}

export interface IVersionListTableColumnConfigType<T = unknown> {
  Header?: keyof StringsMap
  accessor: string
  Cell: Renderer<CellProps<{}, T>>
  hidden?: boolean
  width?: string
  disableSortBy?: boolean
}
