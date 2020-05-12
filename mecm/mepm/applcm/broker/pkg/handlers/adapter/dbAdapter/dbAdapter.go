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
package dbAdapter

import (
	"broker/pkg/handlers/model"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
)

func InsertAppInstanceInfo(db *gorm.DB, n model.AppInstanceInfo) {
	fmt.Printf("Insert App Instance Info (%v, %T)\n", n, n)
	db.Create(&model.AppInstanceInfo{ID: n.ID, AppInstanceName: n.AppInstanceName, AppInstanceDescription: n.AppInstanceDescription,
		AppDID: n.AppDID, AppProvider: n.AppProvider, AppName: n.AppName, AppSoftVersion: n.AppSoftVersion, AppDVersion: n.AppDVersion, AppPkgID: n.AppPkgID, InstantiationState: n.InstantiationState})

	fmt.Printf("Inserting Done")
}

func GetAppInstanceInfo(db *gorm.DB, key string) (appInstInfo model.AppInstanceInfo) {
	fmt.Printf("Get App Instance Info %s", key)
	var appInstanceInfo model.AppInstanceInfo
	returnVal := db.First(&appInstanceInfo, "id=?", key).Error
	if returnVal !=  nil {
		return
	}

	return appInstanceInfo
}

func UpdateAppInstanceInfoInstStatusHostAndWorkloadId(db *gorm.DB, id string, instantiationState string, host string, workloadId string) {
	fmt.Printf("update into DB (%v, %T)\n", id, instantiationState, host, workloadId)

	var appInstInfo model.AppInstanceInfo
	db.Where("id=?", id).First(&appInstInfo).Update("instantiationState", instantiationState).Update("host", host).Update("workloadID", workloadId)
	fmt.Printf("AppName: %s\nAppDID: %s\nAppInstanceDescription:%t\n\n",
		appInstInfo.AppName, appInstInfo.AppDID, appInstInfo.AppInstanceDescription)

	fmt.Printf("Update Done")
}

func UpdateAppInstanceInfoInstStatusAndWorkload(db *gorm.DB, id string, instantiationState string, workloadId string) {
	fmt.Printf("update DB (%v, %T)\n", id, instantiationState)

	var appInstInfo model.AppInstanceInfo
	db.Where("id=?", id).First(&appInstInfo).Update("instantiationState", instantiationState).Update("workloadID", workloadId)
	fmt.Printf("AppName: %s\nAppDID: %s\nAppInstanceDescription:%t\n\n",
		appInstInfo.AppName, appInstInfo.AppDID, appInstInfo.AppInstanceDescription)
	fmt.Printf("Update Done")
}

func UpdateAppInstanceInfoHost(db *gorm.DB, w http.ResponseWriter, id string, host string) {
}

func DeleteAppInstanceInfo(db *gorm.DB, key string) {

	db.Where("id=?", key).Delete(&model.AppInstanceInfo{})

	fmt.Println("Delete App Instance Info: $s", key)
}

func InsertAppPackageInfo(db *gorm.DB, n model.AppPackageInfo) {
	fmt.Printf("Insert App Package Info (%v, %T)\n", n, n)
	db.Create(&model.AppPackageInfo{ID: n.ID, AppDID: n.AppDID, AppProvider: n.AppProvider,
		AppName: n.AppName, AppSoftwareVersion: n.AppSoftwareVersion, AppDVersion: n.AppDVersion,
		OnboardingState: n.OnboardingState, DeployType: n.DeployType, AppPackage: n.AppPackage})

	fmt.Printf("Inserting done")
}

func GetAppPackageInfo(db *gorm.DB, key string) (appPackageInfo model.AppPackageInfo) {
	fmt.Printf("Get App Package Info: %s", key)
	var appPkgInfo model.AppPackageInfo
	err := db.First(&appPkgInfo, "id=?", key).Error
	if err !=  nil {
		return
	}
	return appPkgInfo
}

func DeleteAppPackageInfo(db *gorm.DB, key string) {
	fmt.Printf("Delete App Package Info: %s", key)
	db.Where("id=?", key).Delete(&model.AppPackageInfo{})
	fmt.Println("Delete App Package Info: $s", key)
}
