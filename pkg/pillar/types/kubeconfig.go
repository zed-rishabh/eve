// Copyright (c) 2017-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package types

type KubeConfig struct {
	AppInstUUID string                 `json:"appinst-uuid"`
	Config      map[string]interface{} `json:"kubeconfig"`
}
