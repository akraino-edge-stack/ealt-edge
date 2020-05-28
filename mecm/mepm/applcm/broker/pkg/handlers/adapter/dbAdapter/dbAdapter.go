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
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	username = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbName = os.Getenv("POSTGRES_DATABASE")
	dbHost = os.Getenv("DBHOST")

)

// Database adapter
type DbAdapter struct {
	logger *logrus.Logger
	db     *gorm.DB
}

func NewDbAdapter(logger *logrus.Logger) *DbAdapter {
	return &DbAdapter{logger: logger}
}

// Creates database
func (adapter *DbAdapter) CreateDatabase() {
	adapter.logger.Infof("creating Database...")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password) //Build connection string

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db := conn
	db.AutoMigrate(&model.AppPackageInfo{})
	db.AutoMigrate(&model.AppInstanceInfo{})
	adapter.db = db
}

func (adapter *DbAdapter) InsertAppInstanceInfo(n model.AppInstanceInfo) {
	adapter.logger.Infof("Insert App Instance Info (%v, %T)\n", n, n)
	adapter.db.Create(&model.AppInstanceInfo{ID: n.ID, AppInstanceName: n.AppInstanceName, AppInstanceDescription: n.AppInstanceDescription,
		AppDID: n.AppDID, AppProvider: n.AppProvider, AppName: n.AppName, AppSoftVersion: n.AppSoftVersion, AppDVersion: n.AppDVersion, AppPkgID: n.AppPkgID, InstantiationState: n.InstantiationState})

	adapter.logger.Infof("Inserting Done")
}

func (adapter *DbAdapter) GetAppInstanceInfo(key string) (appInstInfo model.AppInstanceInfo) {
	adapter.logger.Infof("Get App Instance Info %s", key)
	var appInstanceInfo model.AppInstanceInfo
	returnVal := adapter.db.First(&appInstanceInfo, "id=?", key).Error
	if returnVal != nil {
		return
	}

	return appInstanceInfo
}

func (adapter *DbAdapter) UpdateAppInstanceInfoInstStatusHostAndWorkloadId(id string, instantiationState string, host string, workloadId string) {
	adapter.logger.Infof("update into DB (%v, %T)\n", id, instantiationState, host, workloadId)

	var appInstInfo model.AppInstanceInfo
	adapter.db.Where("id=?", id).First(&appInstInfo).Update("instantiationState", instantiationState).Update("host", host).Update("workloadID", workloadId)
	adapter.logger.Infof("AppName: %s\nAppDID: %s\nAppInstanceDescription:%t\n\n",
		appInstInfo.AppName, appInstInfo.AppDID, appInstInfo.AppInstanceDescription)

	adapter.logger.Infof("Update Done")
}

func (adapter *DbAdapter) UpdateAppInstanceInfoInstStatusAndWorkload(id string, instantiationState string, workloadId string) {
	adapter.logger.Infof("update DB (%v, %T)\n", id, instantiationState)

	var appInstInfo model.AppInstanceInfo
	adapter.db.Where("id=?", id).First(&appInstInfo).Update("instantiationState", instantiationState).Update("workloadID", workloadId)
	adapter.logger.Infof("AppName: %s\nAppDID: %s\nAppInstanceDescription:%t\n\n",
		appInstInfo.AppName, appInstInfo.AppDID, appInstInfo.AppInstanceDescription)
	adapter.logger.Infof("Update Done")
}

func (adapter *DbAdapter) DeleteAppInstanceInfo(key string) {
	adapter.db.Where("id=?", key).Delete(&model.AppInstanceInfo{})
	adapter.logger.Infof("Delete App Instance Info: $s", key)
}

func (adapter *DbAdapter) InsertAppPackageInfo(n model.AppPackageInfo) {
	adapter.logger.Infof("Insert App Package Info (%v, %T)\n", n, n)
	adapter.db.Create(&model.AppPackageInfo{ID: n.ID, AppDID: n.AppDID, AppProvider: n.AppProvider,
		AppName: n.AppName, AppSoftwareVersion: n.AppSoftwareVersion, AppDVersion: n.AppDVersion,
		OnboardingState: n.OnboardingState, DeployType: n.DeployType, AppPackage: n.AppPackage})
	adapter.logger.Infof("Inserting done")
}

func (adapter *DbAdapter) GetAppPackageInfo(key string) (appPackageInfo model.AppPackageInfo) {
	adapter.logger.Infof("Get App Package Info: %s", key)
	var appPkgInfo model.AppPackageInfo
	err := adapter.db.First(&appPkgInfo, "id=?", key).Error
	if err != nil {
		return
	}
	return appPkgInfo
}

func (adapter *DbAdapter) DeleteAppPackageInfo(key string) {
	adapter.logger.Infof("Delete App Package Info: %s", key)
	adapter.db.Where("id=?", key).Delete(&model.AppPackageInfo{})
	adapter.logger.Infof("Delete App Package Info: $s", key)
}
