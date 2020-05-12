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
package model

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CreateApplicationReq struct {
	AppDID                string `yaml:"appDId"`
	AppInstancename       string `yaml:"appInstancename"`
	AppInstanceDescriptor string `yaml:"appInstanceDescriptor"`
}

type CreateApplicationRsp struct {
	AppInstanceID string `yaml:"appInstanceId"`
}

type OnBoardPkgApplicationRsp struct {
	AppPkgID string `yaml:"appPkgId"`
}

type InstantiateApplicationReq struct {
	SelectedMECHostInfo struct {
		HostName string `yaml:"hostName"`
		HostID   string `yaml:"hostId"`
	} `yaml:"selectedMECHostInfo"`
}

type QueryApplicationInstanceInfoReq struct {
	Filter            string `yaml:"filter"`
	AttributeSelector string `yaml:"attributeSelector"`
}

type QueryApplicationLCMOperStatusReq struct {
	LifecycleOperationOccurrenceID string `yaml:"lifecycleOperationOccurrenceId"`
}

// User represents a user account
type AppPackageInfo struct {
	//gorm.Model
	ID                 string `gorm:"primary_key;not null;unique"`
	AppDID             string `yaml:"appDId"`
	AppProvider        string `yaml:"appProvider"`
	AppName            string `yaml:"appName"`
	AppSoftwareVersion string `yaml:"appSoftwareVersion"`
	AppDVersion        string `yaml:"appDVersion"`
	OnboardingState    string `yaml:"onboardingState"`
	DeployType         string `yaml:"deployType"`
	AppPackage         string `yaml:"appPackage"`
}

// Task represents a task for the user
type AppInstanceInfo struct {
	//gorm.Model
	ID                     string `gorm:"primary_key;not null;unique"`
	AppInstanceName        string `yaml:"appInstanceName"`
	AppInstanceDescription string `yaml:"appInstanceDescription"`
	AppDID                 string `yaml:"appDId"`
	AppProvider            string `yaml:"appProvider"`
	AppName                string `yaml:"appName"`
	AppSoftVersion         string `yaml:"appSoftVersion"`
	AppDVersion            string `yaml:"appDVersion"`
	AppPkgID               string `yaml:"appPkgId"`
	InstantiationState     string `yaml:"instantiationState"`
	Host                   string `yaml:"host"`
	WorkloadID             string `yaml:"workloadId"`
	InstantiatedAppState   string `yaml:"instantiatedAppState"`
}
