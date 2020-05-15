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
	"broker/pkg/handlers/model"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/ghodss/yaml"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// Handler of REST APIs
type HandlerImpl struct {
	logger     *logrus.Logger
	dbAdapter  *dbAdapter.DbAdapter
}

// Creates handler implementation
func newHandlerImpl(logger *logrus.Logger) (impl HandlerImpl) {
	impl.logger = logger
	impl.dbAdapter = dbAdapter.NewDbAdapter(logger)
	impl.dbAdapter.CreateDatabase()
	return
}

// Uploads package
func (impl *HandlerImpl) UploadPackage(w http.ResponseWriter, r *http.Request) {

	file, header, err := r.FormFile("file")
	defer file.Close()
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}

	var packageName = ""
	f := strings.Split(header.Filename, ".")
	if len(f) > 0 {
		packageName = f[0]
	}
	impl.logger.Infof(packageName)

	pkgPath := PackageFolderPath + header.Filename
	newFile, err := os.Create(pkgPath)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	defer newFile.Close()
	if _, err := newFile.Write(buf.Bytes()); err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	/* Unzip package to decode appDescriptor */
	impl.openPackage(w, pkgPath)

	var yamlFile = PackageFolderPath + packageName + "/Definitions/" + "MainServiceTemplate.yaml"
	appPkgInfo := impl.decodeApplicationDescriptor(w, yamlFile)
	appPkgInfo.AppPackage = header.Filename
	appPkgInfo.OnboardingState = "ONBOARDED"

	impl.logger.Infof("Application package info from package")
	defer r.Body.Close()

	impl.dbAdapter.InsertAppPackageInfo(appPkgInfo)

	/*http.StatusOK*/
	respondJSON(w, http.StatusCreated, appPkgInfo)
}

// Opens package
func (impl *HandlerImpl) openPackage(w http.ResponseWriter, packagePath string) {
	zipReader, _ := zip.OpenReader(packagePath)
	for _, file := range zipReader.Reader.File {

		zippedFile, err := file.Open()
		if err != nil {
			respondError(w, http.StatusBadRequest, err.Error())
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
				respondError(w, http.StatusBadRequest, err.Error())
			}
			defer outputFile.Close()

			_, err = io.Copy(outputFile, zippedFile)
			if err != nil {
				respondError(w, http.StatusBadRequest, err.Error())
			}
		}
	}
}

// Decodes application descriptor
func (impl *HandlerImpl) decodeApplicationDescriptor(w http.ResponseWriter, serviceTemplate string) model.AppPackageInfo {

	yamlFile, err := ioutil.ReadFile(serviceTemplate)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
	}

	jsondata, err := yaml.YAMLToJSON(yamlFile)
	if err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
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

// Query application package information
func (impl *HandlerImpl) QueryAppPackageInfo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appPkgId := params["appPkgId"]
	appPkgInfo := impl.dbAdapter.GetAppPackageInfo(appPkgId)
	if appPkgInfo.ID == "" {
		respondJSON(w, http.StatusNotFound, "ID not exist")
		return
	}
	respondJSON(w, http.StatusAccepted, json.NewEncoder(w).Encode(appPkgInfo))
}

// Deletes application package
func (impl *HandlerImpl) DeleteAppPackage(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	appPkgId := params["appPkgId"]
	appPackageInfo := impl.dbAdapter.GetAppPackageInfo(appPkgId)
	if appPackageInfo.ID == "" {
		respondJSON(w, http.StatusNotFound, "ID not exist")
		return
	}
	impl.dbAdapter.DeleteAppPackageInfo(appPkgId)

	deletePackage := PackageFolderPath + appPackageInfo.AppPackage

	/* Delete ZIP*/
	os.Remove(deletePackage)
	f := strings.Split(appPackageInfo.AppPackage, ".")
	if len(f) > 0 {
		packageName := f[0]
		/*Delete unzipped*/
		os.Remove(packageName)
	}
	respondJSON(w, http.StatusAccepted, json.NewEncoder(w).Encode(""))
}

// Creates application instance
func (impl *HandlerImpl) CreateAppInstance(w http.ResponseWriter, r *http.Request) {
	var req model.CreateApplicationReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	appPkgInfo := impl.dbAdapter.GetAppPackageInfo(req.AppDID)
	if appPkgInfo.ID == "" {
		respondJSON(w, http.StatusNotFound, "ID not exist")
		return
	}
	impl.logger.Infof("Query appPkg Info:", appPkgInfo)

	appInstanceId, err := uuid.NewUUID()
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
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
	impl.dbAdapter.InsertAppInstanceInfo(appInstanceInfo)
	impl.logger.Infof("CreateAppInstance:", req)
	/*http.StatusOK*/
	respondJSON(w, http.StatusCreated, json.NewEncoder(w).Encode(appInstanceInfo))
}

// Instantiates application instance
func (impl *HandlerImpl) InstantiateAppInstance(w http.ResponseWriter, r *http.Request) {
	var req model.InstantiateApplicationReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	params := mux.Vars(r)
	appInstanceId := params["appInstanceId"]

	appInstanceInfo := impl.dbAdapter.GetAppInstanceInfo(appInstanceId)
	appPackageInfo := impl.dbAdapter.GetAppPackageInfo(appInstanceInfo.AppDID)
	if appInstanceInfo.ID == "" || appPackageInfo.ID == "" {
		respondJSON(w, http.StatusNotFound, "ID not exist")
		return
	}

	if appInstanceInfo.InstantiationState == "INSTANTIATED" {
		respondError(w, http.StatusInternalServerError, "Application already instantiated")
		return
	}

	//remove extension
	var packageName = ""
	f := strings.Split(appPackageInfo.AppPackage, ".")
	if len(f) > 0 {
		packageName = f[0]
	}
	impl.logger.Infof(packageName)

	var artifact string
	var pluginInfo string

	switch appPackageInfo.DeployType {
	case "helm":
		pkgPath := PackageFolderPath + packageName + PackageArtifactPath + "Charts"
		artifact = impl.getDeploymentArtifact(pkgPath, ".tar")
		if artifact == "" {
			respondError(w, http.StatusInternalServerError, "artifact not available in application package")
			return
		}
		pluginInfo = "helmplugin" + ":" + os.Getenv("HELM_PLUGIN_PORT")
		impl.logger.Infof("Plugin Info ", pluginInfo)
	case "kubernetes":
		pkgPath := PackageFolderPath + packageName + PackageArtifactPath + "Kubernetes"
		artifact = impl.getDeploymentArtifact(pkgPath, "*.yaml")
		if artifact == "" {
			respondError(w, http.StatusInternalServerError, "artifact not available in application package")
			return
		}
		pluginInfo = "kubernetes.plugin" + ":" + os.Getenv("KUBERNETES_PLUGIN_PORT")
	default:
		respondError(w, http.StatusInternalServerError, "Deployment type not supported")
		return
	}
	impl.logger.Infof("Artifact to deploy:", artifact)

	adapter := pluginAdapter.NewPluginAdapter(pluginInfo, impl.logger)
	workloadId, err, resStatus := adapter.Instantiate(pluginInfo, req.SelectedMECHostInfo.HostID, artifact)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.InvalidArgument {
			respondError(w, http.StatusBadRequest, err.Error())
			return
		} else {
			respondError(w, http.StatusInternalServerError, err.Error())
		}
	}

	if resStatus == "Failure" {
		respondError(w, http.StatusInternalServerError, err.Error())
	}

	impl.dbAdapter.UpdateAppInstanceInfoInstStatusHostAndWorkloadId(appInstanceId, "INSTANTIATED", req.SelectedMECHostInfo.HostID, workloadId)
	respondJSON(w, http.StatusAccepted, json.NewEncoder(w).Encode(workloadId))
}

// Gets deployment artifact
func (impl *HandlerImpl) getDeploymentArtifact(dir string, ext string) string {
	d, err := os.Open(dir)
	if err != nil {
		impl.logger.Infof("Error: ", err)
		return ""
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		impl.logger.Infof("Error: ", err)
		return ""
	}

	impl.logger.Infof("Directory to read " + dir)

	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ext || filepath.Ext(file.Name()) == ".gz" {
				impl.logger.Infof(file.Name())
				impl.logger.Infof(dir + "/" + file.Name())
				return dir + "/" + file.Name()
			}
		}
	}
	return ""
}

// Queries application instance information
func (impl *HandlerImpl) QueryAppInstanceInfo(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	appInstanceId := params["appInstanceId"]

	appInstanceInfo := impl.dbAdapter.GetAppInstanceInfo(appInstanceId)
	appPackageInfo := impl.dbAdapter.GetAppPackageInfo(appInstanceInfo.AppDID)
	if appInstanceInfo.ID == "" || appPackageInfo.ID == "" {
		respondJSON(w, http.StatusNotFound, "ID not exist")
		return
	}
	var instantiatedAppState string
	if appInstanceInfo.InstantiationState == "INSTANTIATED" {

		var pluginInfo string

		switch appPackageInfo.DeployType {
		case "helm":
			pluginInfo = "helmplugin" + ":" + os.Getenv("HELM_PLUGIN_PORT")
		case "kubernetes":
			pluginInfo = "kubernetes.plugin" + ":" + os.Getenv("KUBERNETES_PLUGIN_PORT")
		default:
			respondError(w, http.StatusInternalServerError, "Deployment type not supported")
			return
		}

		adapter := pluginAdapter.NewPluginAdapter(pluginInfo, impl.logger)
		state, err := adapter.Query(pluginInfo, appInstanceInfo.Host, appInstanceInfo.WorkloadID)
		if err != nil {
			respondError(w, http.StatusInternalServerError, err.Error())
			return
		}
		instantiatedAppState = state
	}
	appInstanceInfo.InstantiatedAppState = instantiatedAppState

	respondJSON(w, http.StatusCreated, json.NewEncoder(w).Encode(appInstanceInfo))
}

// Queries application lcm operation status
func (impl *HandlerImpl) QueryAppLcmOperationStatus(w http.ResponseWriter, r *http.Request) {
	var req model.QueryApplicationLCMOperStatusReq
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Fprintf(w, "QueryApplicationLCMOperStatus: %+v", req)
}

// Terminates application instance
func (impl *HandlerImpl) TerminateAppInstance(w http.ResponseWriter, r *http.Request) {
	impl.logger.Infof("TerminateAppInstance...")
	params := mux.Vars(r)
	appInstanceId := params["appInstanceId"]

	appInstanceInfo := impl.dbAdapter.GetAppInstanceInfo(appInstanceId)
	appPackageInfo := impl.dbAdapter.GetAppPackageInfo(appInstanceInfo.AppDID)
	if appInstanceInfo.ID == "" || appPackageInfo.ID == "" {
		respondJSON(w, http.StatusNotFound, "ID not exist")
		return
	}

	if appInstanceInfo.InstantiationState == "NOT_INSTANTIATED" {
		respondError(w, http.StatusNotAcceptable, "instantiationState: NOT_INSTANTIATED")
		return
	}

	var pluginInfo string
	switch appPackageInfo.DeployType {
	case "helm":
		pluginInfo = "helmplugin" + ":" + os.Getenv("HELM_PLUGIN_PORT")
	case "kubernetes":
		pluginInfo = "kubernetes.plugin" + ":" + os.Getenv("KUBERNETES_PLUGIN_PORT")
	default:
		respondError(w, http.StatusInternalServerError, "Deployment type not supported")
		return
	}

	adapter := pluginAdapter.NewPluginAdapter(pluginInfo, impl.logger)
	_, err := adapter.Terminate(pluginInfo, appInstanceInfo.Host, appInstanceInfo.WorkloadID)
	if err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	impl.dbAdapter.UpdateAppInstanceInfoInstStatusAndWorkload(appInstanceId, "NOT_INSTANTIATED", "")

	respondJSON(w, http.StatusAccepted, json.NewEncoder(w).Encode(""))
}

// Deletes application instance identifier
func (impl *HandlerImpl) DeleteAppInstanceIdentifier(w http.ResponseWriter, r *http.Request) {
	impl.logger.Infof("DeleteAppInstanceIdentifier:")
	params := mux.Vars(r)
	appInstanceId := params["appInstanceId"]

	impl.dbAdapter.DeleteAppInstanceInfo(appInstanceId)
	respondJSON(w, http.StatusOK, json.NewEncoder(w).Encode(""))
}

// It makes the JSON
func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

// RespondError makes the error response with payload as json format
func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}
