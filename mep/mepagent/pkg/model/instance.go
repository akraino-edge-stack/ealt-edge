/*
 *  Copyright 2020 Huawei Technologies Co., Ltd.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

// define the type information
package model

type AppInstanceInfo struct {
	AppInstanceId                            string                                    `yaml:"appInstanceId" json:"appInstanceId"`
	MepServerIP                              string                                    `yaml:"mepServerIP" json:"mepServerIP"`
	MepServerPORT                            string                                    `yaml:"mepServerPORT" json:"mepServerPORT"`
	ServiceInfoPosts                         []ServiceInfoPost                         `yaml:"serviceInfoPosts" json:"serviceInfoPosts"`
	SerAvailabilityNotificationSubscriptions []SerAvailabilityNotificationSubscription `yaml:"serAvailabilityNotificationSubscriptions" json:"serAvailabilityNotificationSubscriptions"`
}

type ServiceInfoPost struct {
	SerInstanceId     string         `yaml:"serInstanceId" json:"serInstanceId"`
	SerName           string         `yaml:"serName" json:"serName"`
	SerCategory       CategoryRef    `yaml:"serCategory" json:"serCategory"`
	Version           string         `yaml:"version" json:"version"`
	State             ServiceState   `yaml:"state" json:"state"`
	TransportId       string         `yaml:"transportId" json:"transportId"`
	TransportInfo     TransportInfo  `yaml:"transportInfo" json:"transportInfo"`
	Serializer        SerializerType `yaml:"serializer" json:"serializer"`
	ScopeOfLocality   LocalityType   `yaml:"scopeOfLocality" json:"scopeOfLocality"`
	ConsumedLocalOnly bool           `yaml:"consumedLocalOnly" json:"consumedLocalOnly"`
	IsLocal           bool           `yaml:"isLocal" json:"isLocal"`
}

type CategoryRef struct {
	Href    string `yaml:"href" json:"href"`
	Id      string `yaml:"id" json:"id"`
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
}

type ServiceState string

const (
	ACTIVE   ServiceState = "ACTIVE"
	INACTIVE ServiceState = "INACTIVE"
)

type TransportInfo struct {
	Id               string           `yaml:"id" json:"id"`
	Name             string           `yaml:"name" json:"name"`
	Description      string           `yaml:"description" json:"description"`
	TransportType    TransportType    `yaml:"type" json:"type"`
	Protocol         string           `yaml:"protocol" json:"protocol"`
	Version          string           `yaml:"version" json:"version"`
	Endpoint         EndPointInfoUris `yaml:"endpoint" json:"endpoint"`
	Security         SecurityInfo     `yaml:"security" json:"security"`
	ImplSpecificInfo ImplSpecificInfo `yaml:"implSpecificInfo" json:"implSpecificInfo"`
}

type TransportType string

const (
	REST_HTTP      TransportType = "REST_HTTP"
	MB_TOPIC_BASED TransportType = "MB_TOPIC_BASED"
	MB_ROUTING     TransportType = "MB_ROUTING"
	MB_PUBSUB      TransportType = "MB_PUBSUB"
	RPC            TransportType = "RPC"
	RPC_STREAMING  TransportType = "RPC_STREAMING"
	WEBSOCKET      TransportType = "WEBSOCKET"
)

type EndPointInfoUris struct {
	Uris []string `yaml:"uris" json:"uris"`
}

type SecurityInfo struct {
	OAuth2Info SecurityInfoOAuth2Info `yaml:"oAuth2Info" json:"oAuth2Info"`
}

type SecurityInfoOAuth2Info struct {
	GrantTypes    []SecurityInfoOAuth2InfoGrantType `yaml:"grantTypes" json:"grantTypes"`
	TokenEndpoint string                            `yaml:"tokenEndpoint" json:"tokenEndpoint"`
}

type SecurityInfoOAuth2InfoGrantType string

const (
	AUTHORIZATION_CODE SecurityInfoOAuth2InfoGrantType = "OAUTH2_AUTHORIZATION_CODE"
	IMPLICIT_GRANT     SecurityInfoOAuth2InfoGrantType = "OAUTH2_IMPLICIT_GRANT"
	RESOURCE_OWNER     SecurityInfoOAuth2InfoGrantType = "OAUTH2_RESOURCE_OWNER"
	CLIENT_CREDENTIALS SecurityInfoOAuth2InfoGrantType = "OAUTH2_CLIENT_CREDENTIALS"
)

type ImplSpecificInfo struct {
}

type SerializerType string

const (
	JSON      SerializerType = "JSON"
	XML       SerializerType = "XML"
	PROTOBUF3 SerializerType = "PROTOBUF3"
)

type LocalityType string

const (
	MEC_SYSTEM LocalityType = "MEC_SYSTEM"
	MEC_HOST   LocalityType = "MEC_HOST"
	NFVI_POP   LocalityType = "NFVI_POP"
	ZONE       LocalityType = "ZONE"
	ZONE_GROUP LocalityType = "ZONE_GROUP"
	NFVI_NODE  LocalityType = "NFVI_NODE"
)

type SerAvailabilityNotificationSubscription struct {
	SubscriptionType  string                                                   `yaml:"subscriptionType" json:"subscriptionType"`
	CallbackReference string                                                   `yaml:"callbackReference" json:"callbackReference"`
	Links             Self                                                     `yaml:"links" json:"links"`
	FilteringCriteria SerAvailabilityNotificationSubscriptionFilteringCriteria `yaml:"filteringCriteria" json:"filteringCriteria"`
}

type Self struct {
	Self LinkType `yaml:"self" json:"self"`
}

type LinkType struct {
	Href string `yaml:"href" json:"href"`
}

type SerAvailabilityNotificationSubscriptionFilteringCriteria struct {
	SerInstanceIds []string       `yaml:"serInstanceIds" json:"serInstanceIds"`
	SerNames       []string       `yaml:"serNames" json:"serNames"`
	SerCategories  []CategoryRef  `yaml:"serCategories" json:"serCategories"`
	States         []ServiceState `yaml:"states" json:"states"`
	IsLocal        bool           `yaml:"isLocal" json:"isLocal"`
}
