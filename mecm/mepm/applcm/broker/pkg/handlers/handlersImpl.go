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
package handlers

import (
	"archive/zip"
	"broker/pkg/handlers/adapter/dbAdapter"
	"broker/pkg/handlers/adapter/pluginAdapter"
	"broker/pkg/handlers/common"
	"broker/pkg/handlers/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/ghodss/yaml"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func UploadFileHldr(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		common.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		common.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	var packageName = ""
	f := strings.Split(header.Filename, ".")
	if len(f) > 0 {
		packageName = f[0]
	}
	fmt.Println(packageName)

	pkgPath := PackageFolderPath + header.Filename
	newFile, err := os.Create(pkgPath)
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer newFile.Close()
	if _, err := newFile.Write(buf.Bytes()); err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	/* Unzip package to decode appDescriptor */
	openPackage(w, pkgPath)

	var yamlFile = PackageFolderPath + packageName + "/Definitions/" + "MainServiceTemplate.yaml"
	appPkgInfo := decodeApplicationDescriptor(w, yamlFile)
	appPkgInfo.AppPackage = header.Filename
	appPkgInfo.OnboardingState = "ONBOARDED"

	log.Println("Application package info from package")
	defer r.Body.Close()

	dbAdapter.InsertAppPackageInfo(db, appPkgInfo)

	/*http.StatusOK*/
	common.RespondJSON(w, http.StatusCreated, appPkgInfo)
}

func openPackage(w http.ResponseWriter, packagePath string) {
	zipReader, _ := zip.OpenReader(packagePath)
	for _, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err != nil {
			common.RespondError(w, http.StatusBadRequest, err.Error())
		}
		defer zippedFile.Close()

		targetDir := PackageFolderPath + "/"
		extractedFilePath := filepath.Join(
			targetDir,
			file.Name,
		)

		if file.FileInfo().IsDir() {
			os.MkdirAll(extractedFilePath, file.Mode())
		} else {
			outputFile, err := os.OpenFile(
				extractedFilePath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
				file.Mode(),
			)
			if err != nil {
				common.RespondError(w, http.StatusBadRequest, err.Error())
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				common.RespondError(w, http.StatusBadRequest, err.Error())
			}
		}
	}
}

func decodeApplicationDescriptor(w http.ResponseWriter, serviceTemplate string) model.AppPackageInfo {

	yamlFile, err := ioutil.ReadFile(serviceTemplate)
	if err != nil {
		common.RespondError(w, http.StatusBadRequest, err.Error())
	}

	jsondata, err := yaml.YAMLToJSON(yamlFile)
	if err != nil {
		common.RespondError(w, http.StatusBadRequest, err.Error())
	}

	appDId, _, _, _ := jsonparser.Get(jsondata, "topology_template", "node_templates", "face_recognition", "properties", "appDId")
	appProvider, _, _, _ := jsonparser.Get(jsondata, "topology_template", "node_templates", "face_recognition", "properties", "appProvider")
	appInfoName, _, _, _ := jsonparser.Get(jsondata, "topology_template", "node_templates", "face_recognition", "properties", "appInfoName")
	appSoftVersion, _, _, _ := jsonparser.Get(jsondata, "topology_template", "node_templates", "face_recognition", "properties", "appSoftVersion")
	appDVersion, _, _, _ := jsonparser.Get(jsondata, "topology_template", "node_templates", "face_recognition", "properties", "appDVersion")
	deployType, _, _, _ := jsonparser.Get(jsondata, "topology_template", "node_templates", "face_recognition", "properties", "type")

	appPkgInfo := model.AppPackageInfo{
		ID:                 string(appDId),
		AppDID:             string(appDId),
		AppProvider:        string(appProvider),
		AppName:            string(appInfoName),
		AppSoftwareVersion: string(appSoftVersion),
		AppDVersion:        string(appDVersion),
		DeployType:         string(deployType),
	}

	//return appPackageInfo
	return appPkgInfo
}

func QueryAppPackageInfo(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appPkgId := params["appPkgId"]
	appPkgInfo := dbAdapter.GetAppPackageInfo(db, appPkgId)
	if appPkgInfo.ID == "" {
		common.RespondJSON(w, http.StatusNotFound, "ID not exist")
		return
	}
	common.RespondJSON(w, http.StatusAccepted, json.NewEncoder(w).Encode(appPkgInfo))
}

func DeleteAppPackage(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appPkgId := params["appPkgId"]
	appPackageInfo := dbAdapter.GetAppPackageInfo(db, appPkgId)
	if appPackageInfo.ID == "" {
		common.RespondJSON(w, http.StatusNotFound, "ID not exist")
		return
	}
	dbAdapter.DeleteAppPackageInfo(db, appPkgId)

	deletePackage := PackageFolderPath + appPackageInfo.AppPackage

	/* Delete ZIP*/
	os.Remove(deletePackage)
	f := strings.Split(appPackageInfo.AppPackage, ".")
	if len(f) > 0 {
		packageName := f[0]
		/*Delete unzipped*/
		os.Remove(packageName)
	}
	common.RespondJSON(w, http.StatusAccepted, json.NewEncoder(w).Encode(""))
}

func CreateAppInstanceHldr(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var req model.CreateApplicationReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	appPkgInfo := dbAdapter.GetAppPackageInfo(db, req.AppDID)
	if appPkgInfo.ID == "" {
		common.RespondJSON(w, http.StatusNotFound, "ID not exist")
		return
	}
	fmt.Println("Query appPkg Info:", appPkgInfo)

	appInstanceId, err := uuid.NewUUID()
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
	}

	appInstanceInfo := model.AppInstanceInfo{

		ID:                     appInstanceId.String(),
		AppInstanceName:        req.AppInstancename,
		AppInstanceDescription: req.AppInstanceDescriptor,
		AppDID:                 req.AppDID,
		AppProvider:            appPkgInfo.AppProvider,
		AppName:                appPkgInfo.AppName,
		AppSoftVersion:         appPkgInfo.AppSoftwareVersion,
		AppDVersion:            appPkgInfo.AppDVersion,
		AppPkgID:               appPkgInfo.AppDID,
		InstantiationState:     "NOT_INSTANTIATED",
	}
	dbAdapter.InsertAppInstanceInfo(db, appInstanceInfo)
	fmt.Println("CreateAppInstanceHldr:", req)
	/*http.StatusOK*/
	common.RespondJSON(w, http.StatusCreated, json.NewEncoder(w).Encode(appInstanceInfo))
}

func InstantiateAppInstanceHldr(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var req model.InstantiateApplicationReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	params := mux.Vars(r)
	appInstanceId := params["appInstanceId"]

	appInstanceInfo := dbAdapter.GetAppInstanceInfo(db, appInstanceId)
	appPackageInfo := dbAdapter.GetAppPackageInfo(db, appInstanceInfo.AppDID)
	if appInstanceInfo.ID == "" || appPackageInfo.ID == "" {
		common.RespondJSON(w, http.StatusNotFound, "ID not exist")
		return
	}

	if appInstanceInfo.InstantiationState == "INSTANTIATED" {
		common.RespondError(w, http.StatusInternalServerError, "Application already instantiated")
		return
	}

	//remove extension
	var packageName = ""
	f := strings.Split(appPackageInfo.AppPackage, ".")
	if len(f) > 0 {
		packageName = f[0]
	}
	fmt.Println(packageName)

	var artifact string
	var pluginInfo string

	switch appPackageInfo.DeployType {
	case "helm":
		pkgPath := PackageFolderPath + packageName + PackageArtifactPath + "Charts"
		artifact = getDeploymentArtifact(pkgPath, ".tar")
		if artifact == "" {
			common.RespondError(w, http.StatusInternalServerError, "artifact not available in application package")
			return
		}
		pluginInfo = "helm.plugin" + ":" + os.Getenv("HELM_PLUGIN_PORT")
	case "kubernetes":
		pkgPath := PackageFolderPath + packageName + PackageArtifactPath + "Kubernetes"
		artifact = getDeploymentArtifact(pkgPath, "*.yaml")
		if artifact == "" {
			common.RespondError(w, http.StatusInternalServerError, "artifact not available in application package")
			return
		}
		pluginInfo = "kubernetes.plugin" + ":" + os.Getenv("KUBERNETES_PLUGIN_PORT")
	default:
		common.RespondError(w, http.StatusInternalServerError, "Deployment type not supported")
		return
	}
	fmt.Println("Artifact to deploy:", artifact)

	workloadId, err := pluginAdapter.Instantiate(pluginInfo, req.SelectedMECHostInfo.HostID, artifact)
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	dbAdapter.UpdateAppInstanceInfoInstStatusHostAndWorkloadId(db, appInstanceId, "INSTANTIATED", req.SelectedMECHostInfo.HostID, workloadId)

	common.RespondJSON(w, http.StatusAccepted, json.NewEncoder(w).Encode(""))
}

func getDeploymentArtifact(dir string, ext string) string {
	d, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println("Directory to read " + dir)

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ext || filepath.Ext(file.Name()) == ".gz" {
				fmt.Println(file.Name())
				fmt.Println(dir + "/" + file.Name())
				return dir + "/" + file.Name()
			}
		}
	}
	return ""
}

func QueryAppInstanceInfoHldr(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	appInstanceId := params["appInstanceId"]

	appInstanceInfo := dbAdapter.GetAppInstanceInfo(db, appInstanceId)
	appPackageInfo := dbAdapter.GetAppPackageInfo(db, appInstanceInfo.AppDID)
	if appInstanceInfo.ID == "" || appPackageInfo.ID == "" {
		common.RespondJSON(w, http.StatusNotFound, "ID not exist")
		return
	}
	var instantiatedAppState string
	if appInstanceInfo.InstantiationState == "INSTANTIATED" {

		var pluginInfo string

		switch appPackageInfo.DeployType {
		case "helm":
			pluginInfo = "helm.plugin" + ":" + os.Getenv("HELM_PLUGIN_PORT")
		case "kubernetes":
			pluginInfo = "kubernetes.plugin" + ":" + os.Getenv("KUBERNETES_PLUGIN_PORT")
		default:
			common.RespondError(w, http.StatusInternalServerError, "Deployment type not supported")
			return
		}

		state, err := pluginAdapter.Query(pluginInfo, appInstanceInfo.Host, appInstanceInfo.WorkloadID)
		if err != nil {
			common.RespondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		instantiatedAppState = state
	}
	appInstanceInfo.InstantiatedAppState = instantiatedAppState

	common.RespondJSON(w, http.StatusCreated, json.NewEncoder(w).Encode(appInstanceInfo))
}

func QueryAppLcmOperationStatusHldr(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var req model.QueryApplicationLCMOperStatusReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Fprintf(w, "QueryApplicationLCMOperStatus: %+v", req)
}

func TerminateAppInsHldr(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	log.Println("TerminateAppInsHldr...")
	params := mux.Vars(r)
	appInstanceId := params["appInstanceId"]

	appInstanceInfo := dbAdapter.GetAppInstanceInfo(db, appInstanceId)
	appPackageInfo := dbAdapter.GetAppPackageInfo(db, appInstanceInfo.AppDID)
	if appInstanceInfo.ID == "" || appPackageInfo.ID == "" {
		common.RespondJSON(w, http.StatusNotFound, "ID not exist")
		return
	}

	if appInstanceInfo.InstantiationState == "NOT_INSTANTIATED" {
		common.RespondError(w, http.StatusNotAcceptable, "instantiationState: NOT_INSTANTIATED")
		return
	}

	var pluginInfo string
	switch appPackageInfo.DeployType {
	case "helm":
		pluginInfo = "helm.plugin" + ":" + os.Getenv("HELM_PLUGIN_PORT")
	case "kubernetes":
		pluginInfo = "kubernetes.plugin" + ":" + os.Getenv("KUBERNETES_PLUGIN_PORT")
	default:
		common.RespondError(w, http.StatusInternalServerError, "Deployment type not supported")
		return
	}

	_, err := pluginAdapter.Terminate(pluginInfo, appInstanceInfo.Host, appInstanceInfo.WorkloadID)
	if err != nil {
		common.RespondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	dbAdapter.UpdateAppInstanceInfoInstStatusAndWorkload(db, appInstanceId, "NOT_INSTANTIATED", "")

	common.RespondJSON(w, http.StatusAccepted, json.NewEncoder(w).Encode(""))
}
func DeleteAppInstanceIdentifierHldr(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteAppInstanceIdentifierHldr:")
	params := mux.Vars(r)
	appInstanceId := params["appInstanceId"]

	dbAdapter.DeleteAppInstanceInfo(db, appInstanceId)
	common.RespondJSON(w, http.StatusOK, json.NewEncoder(w).Encode(""))
}
