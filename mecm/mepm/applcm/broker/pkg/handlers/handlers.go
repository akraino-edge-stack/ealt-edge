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
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

// URLS
const (
	CreateAppInstance = "/ealtedge/mepm/app_lcm/v1/app_instances"
    InstantiateAppInstance = "/ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}/instantiate"
    QueryAppInstanceInfo = "/ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}"
    QueryAppLcmOperationStatus = "/ealtedge/mepm/app_lcm/v1/app_lcm_op_occs"
    TerminateAppIns = "/ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}/terminate"
    DeleteAppInstanceIdentifier = "/ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}"
    OnboardPackage = "/ealtedge/mepm/app_pkgm/v1/app_packages"
    QueryOnboardPackage = "/ealtedge/mepm/app_pkgm/v1/app_packages/{appPkgId}"
)

// Package paths, to be created in deployment file (docker-compose/k8s yaml/helm)
const (
	PackageFolderPath = "/go/release/application/packages/"
	PackageArtifactPath = "/Artifacts/Deployment/"
)

// Handler of REST APIs
type Handlers struct {
	router *mux.Router
	logger *logrus.Logger
	impl   HandlerImpl
}

// Initialize initializes the handler
func (hdlr *Handlers) Initialize(logger *logrus.Logger) {
	hdlr.router = mux.NewRouter()
	hdlr.logger = logger
	hdlr.setRouters()
	hdlr.impl = newHandlerImpl(hdlr.logger)
}

// Run on it's router
func (hdlr *Handlers) Run(host string) {
	hdlr.logger.Info("Server is running on port %s", host)
	err := http.ListenAndServe(host, hdlr.router)
	if err != nil {
		hdlr.logger.Fatalf("Server couldn't run on port %s", host)
	}
}

// SetRouters sets the all required routers
func (hdlr *Handlers) setRouters() {
	// Routing for handling the requests
	hdlr.Post(OnboardPackage, hdlr.handleRequest(hdlr.impl.UploadPackage))
	hdlr.Get(QueryOnboardPackage, hdlr.handleRequest(hdlr.impl.QueryAppPackageInfo))
	hdlr.Post(CreateAppInstance, hdlr.handleRequest(hdlr.impl.CreateAppInstance))
	hdlr.Delete(QueryOnboardPackage, hdlr.handleRequest(hdlr.impl.DeleteAppPackage))
	hdlr.Post(InstantiateAppInstance, hdlr.handleRequest(hdlr.impl.InstantiateAppInstance))
	hdlr.Get(QueryAppInstanceInfo, hdlr.handleRequest(hdlr.impl.QueryAppInstanceInfo))
	hdlr.Get(QueryAppLcmOperationStatus, hdlr.handleRequest(hdlr.impl.QueryAppLcmOperationStatus))
	hdlr.Post(TerminateAppIns, hdlr.handleRequest(hdlr.impl.TerminateAppInstance))
	hdlr.Delete(DeleteAppInstanceIdentifier, hdlr.handleRequest(hdlr.impl.DeleteAppInstanceIdentifier))
}

// Get wraps the router for GET method
func (hdlr *Handlers) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	hdlr.router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (hdlr *Handlers) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	hdlr.router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (hdlr *Handlers) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	hdlr.router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (hdlr *Handlers) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	hdlr.router.HandleFunc(path, f).Methods("DELETE")
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (hdlr *Handlers) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
