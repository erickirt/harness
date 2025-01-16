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
import vsCodeWebIcon from 'cde-gitness/assests/vsCodeWeb.svg?url'
import vsCodeIcon from 'cde-gitness/assests/VSCode.svg?url'
import intellijIcon from 'cde-gitness/assests/intellij.svg?url'
import cLionIcon from 'cde-gitness/assests/clion.svg?url'
import phpStormIcon from 'cde-gitness/assests/phpStorm.svg?url'
import pyCharmIcon from 'cde-gitness/assests/pyCharm.svg?url'
import rubyMineIcon from 'cde-gitness/assests/rubyMine.svg?url'
import webStormIcon from 'cde-gitness/assests/webStorm.svg?url'
import goLandIcon from 'cde-gitness/assests/goLand.svg?url'
import riderIcon from 'cde-gitness/assests/rider.svg?url'
import type { StringsMap } from 'framework/strings/stringTypes'

export const docLink = 'https://developer.harness.io/docs/cloud-development-environments'

export enum IDEType {
  VSCODE = 'vs_code',
  VSCODEWEB = 'vs_code_web',
  INTELLIJ = 'intellij',
  PYCHARM = 'pycharm',
  GOLAND = 'goland',
  WEBSTORM = 'webstorm',
  CLION = 'clion',
  PHPSTORM = 'phpstorm',
  RUBYMINE = 'rubymine',
  RIDER = 'rider'
}

export interface ideType {
  label: keyof StringsMap
  value: string
  icon: any
}

export const groupEnums = {
  VSCODE: 'vscode',
  JETBRAIN: 'jetbrain'
}

export const getIDETypeOptions = (getString: any) => [
  {
    label: getString('cde.ide.browser'),
    value: IDEType.VSCODEWEB,
    icon: vsCodeWebIcon,
    group: groupEnums.VSCODE,
    buttonText: getString('cde.details.openBrowser')
  },
  {
    label: getString('cde.ide.desktop'),
    value: IDEType.VSCODE,
    icon: vsCodeIcon,
    allowSSH: true,
    group: groupEnums.VSCODE,
    buttonText: getString('cde.details.openEditor')
  },
  {
    label: getString('cde.ide.clion'),
    value: IDEType.CLION,
    icon: cLionIcon,
    allowSSH: true,
    group: groupEnums.JETBRAIN,
    buttonText: getString('cde.details.openJetBrain')
  },
  {
    label: getString('cde.ide.goland'),
    value: IDEType.GOLAND,
    icon: goLandIcon,
    allowSSH: true,
    group: groupEnums.JETBRAIN,
    buttonText: getString('cde.details.openJetBrain')
  },
  {
    label: getString('cde.ide.intellij'),
    value: IDEType.INTELLIJ,
    icon: intellijIcon,
    allowSSH: true,
    group: groupEnums.JETBRAIN,
    buttonText: getString('cde.details.openJetBrain')
  },
  {
    label: getString('cde.ide.phpstorm'),
    value: IDEType.PHPSTORM,
    icon: phpStormIcon,
    allowSSH: true,
    group: groupEnums.JETBRAIN,
    buttonText: getString('cde.details.openJetBrain')
  },
  {
    label: getString('cde.ide.pycharm'),
    value: IDEType.PYCHARM,
    icon: pyCharmIcon,
    allowSSH: true,
    group: groupEnums.JETBRAIN,
    buttonText: getString('cde.details.openJetBrain')
  },
  {
    label: getString('cde.ide.rider'),
    value: IDEType.RIDER,
    icon: riderIcon,
    allowSSH: true,
    group: groupEnums.JETBRAIN,
    buttonText: getString('cde.details.openJetBrain')
  },
  {
    label: getString('cde.ide.rubymine'),
    value: IDEType.RUBYMINE,
    icon: rubyMineIcon,
    allowSSH: true,
    group: groupEnums.JETBRAIN,
    buttonText: getString('cde.details.openJetBrain')
  },
  {
    label: getString('cde.ide.webstorm'),
    value: IDEType.WEBSTORM,
    icon: webStormIcon,
    allowSSH: true,
    group: groupEnums.JETBRAIN,
    buttonText: getString('cde.details.openJetBrain')
  }
]

export enum EnumGitspaceCodeRepoType {
  GITHUB = 'github',
  GITLAB = 'gitlab',
  HARNESS_CODE = 'harness_code',
  BITBUCKET = 'bitbucket',
  UNKNOWN = 'unknown',
  GITNESS = 'gitness'
}

export enum GitspaceStatus {
  RUNNING = 'running',
  STOPPED = 'stopped',
  UNKNOWN = 'unknown',
  ERROR = 'error',
  STARTING = 'starting',
  STOPPING = 'stopping',
  UNINITIALIZED = 'uninitialized'
}

export interface GitspaceStatusTypesListItem {
  label: string
  value: GitspaceStatus
}

export const GitspaceStatusTypes = (getString: any) => [
  {
    label: getString('cde.gitspaceStatus.active'),
    value: GitspaceStatus.RUNNING
  },
  {
    label: getString('cde.gitspaceStatus.stopped'),
    value: GitspaceStatus.STOPPED
  },
  {
    label: getString('cde.gitspaceStatus.error'),
    value: GitspaceStatus.ERROR
  }
]

export const getIDEOption: any = (type = '', getString = null) => {
  let ideItem = null
  if (type && getString) {
    ideItem = getIDETypeOptions(getString).find((ide: ideType) => ide?.value === type)
  }
  return ideItem
}

export enum GitspaceOwnerType {
  SELF = 'self',
  ALL = 'all'
}

export interface GitspaceOwnerTypeListItem {
  label: string
  value: GitspaceOwnerType
}

export const GitspaceOwnerTypes = (getString: any) => [
  {
    label: getString('cde.gitspaceOwners.allGitspaces'),
    value: GitspaceOwnerType.ALL
  },
  {
    label: getString('cde.gitspaceOwners.myGitspaces'),
    value: GitspaceOwnerType.SELF
  }
]

export enum GitspaceActionType {
  START = 'start',
  STOP = 'stop'
}

export enum GitspaceRegion {
  USEast = 'us-east',
  USWest = 'us-west',
  Europe = 'Europe',
  Australia = 'Australia'
}

export enum SortByType {
  CREATED = 'created',
  LAST_USED = 'last_used',
  LAST_ACTIVATED = 'last_activated'
}

export interface SortByTypeListItem {
  label: string
  value: SortByType
}

export const SortByTypes = (getString: any) => [
  {
    label: getString('cde.created'),
    value: SortByType.CREATED
  },
  {
    label: getString('cde.lastUsed'),
    value: SortByType.LAST_USED
  },
  {
    label: getString('cde.lastStarted'),
    value: SortByType.LAST_ACTIVATED
  }
]
