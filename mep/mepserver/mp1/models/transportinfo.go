/*
 * Copyright 2020 Huawei Technologies Co., Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package models

// This type represents the general information of a MEC service.
type TransportInfo struct {
	ID          string         `json:"id,omitempty"`
	Name        string         `json:"name,omitempty"`
	Description string         `json:"description,omitempty"`
	TransType   TransportTypes `json:"type,omitempty"`
	Protocol    string         `json:"protocol,omitempty"`
	Version     string         `json:"version,omitempty"`
	// This type represents information about a transport endpoint
	Endpoint         EndPointInfo `json:"endpoint,omitempty"`
	Security         SecurityInfo `json:"security,omitempty"`
	ImplSpecificInfo interface{}  `json:"implSpecificInfo,omitempty"`
}

type TransportTypes string
