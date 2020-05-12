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
	"broker/pkg/handlers/model"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
	"os"
	"time"
)

const CreateAppInstance = "/ealtedge/mepm/app_lcm/v1/app_instances"
const InstantiateAppInstance = "/ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}/instantiate"
const QueryAppInstanceInfo = "/ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}"
const QueryAppLcmOperationStatus = "/ealtedge/mepm/app_lcm/v1/app_lcm_op_occs"
const TerminateAppIns = "/ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}/terminate"
const DeleteAppInstanceIdentifier = "/ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}"
const OnboardPackage = "/ealtedge/mepm/app_pkgm/v1/app_packages"
const QueryOnboardPackage = "/ealtedge/mepm/app_pkgm/v1/app_packages/{appPkgId}"

const PackageFolderPath = "/go/release/application/packages/"
const PackageArtifactPath = "/Artifacts/Deployment/"

type Handlers struct {
	Router *mux.Router
	logger *log.Logger
	db     *gorm.DB
}

const DB_NAME = "applcmDB"

// Run the app on it's router
func (hdlr *Handlers) Run(host string) {
	fmt.Println("Binding to port...: %d", host)
	log.Fatal(http.ListenAndServe(host, hdlr.Router))
}

func createDatabase() *gorm.DB {
	fmt.Println("creating Database...")

	usrpswd := os.Getenv("MYSQL_USER") + ":" + os.Getenv("MYSQL_PASSWORD")
	host := "@tcp(" + "dbhost" + ":3306)/"

	db, err := gorm.Open("mysql", usrpswd + host)
	if err != nil {
		fmt.Println("Database connect error", err.Error())
	}
//	db = db.Exec("DROP DATABASE IF EXISTS " +  DB_NAME)
//	db = db.Exec("CREATE DATABASE "+ DB_NAME)
	db.Exec("CREATE DATABASE  " + DB_NAME)
	db.Exec("USE applcmDB")

	//db.Close()
	//db, err = gorm.Open("mysql", usrpswd + host + DB_NAME + "?charset=utf8&parseTime=True")
	/*if err != nil {
		fmt.Println("Database connect error", err.Error())
	} else {
		fmt.Println("Database connected successfully")
	}*/
	gorm.DefaultCallback.Create().Remove("mysql:set_identity_insert")

	fmt.Println("Migrating models...")
	db.AutoMigrate(&model.AppPackageInfo{})
	db.AutoMigrate(&model.AppInstanceInfo{})
	//db.LogMode(true)
	return db
}

// Initialize initializes the app with predefined configuration
func (hdlr *Handlers) Initialize(logger *log.Logger) {
	hdlr.Router = mux.NewRouter()

	hdlr.logger = logger
	hdlr.setRouters()
	hdlr.db = createDatabase()
}

func (hdlr *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer hdlr.logger.Printf("request processed in %s\n", time.Now().Sub(startTime))
		next(w, r)
	}
}

// setRouters sets the all required routers
func (hdlr *Handlers) setRouters() {
	// Routing for handling the requests
	hdlr.Post(OnboardPackage, hdlr.handleRequest(UploadFileHldr))
	hdlr.Get(QueryOnboardPackage, hdlr.handleRequest(QueryAppPackageInfo))
	hdlr.Post(CreateAppInstance, hdlr.handleRequest(CreateAppInstanceHldr))
	hdlr.Delete(QueryOnboardPackage, hdlr.handleRequest(DeleteAppPackage))
	hdlr.Post(InstantiateAppInstance, hdlr.handleRequest(InstantiateAppInstanceHldr))
	hdlr.Get(QueryAppInstanceInfo, hdlr.handleRequest(QueryAppInstanceInfoHldr))
	hdlr.Get(QueryAppLcmOperationStatus, hdlr.handleRequest(QueryAppLcmOperationStatusHldr))
	hdlr.Post(TerminateAppIns, hdlr.handleRequest(TerminateAppInsHldr))
	hdlr.Delete(DeleteAppInstanceIdentifier, hdlr.handleRequest(DeleteAppInstanceIdentifierHldr))
}

// Get wraps the router for GET method
func (hdlr *Handlers) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	hdlr.Router.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (hdlr *Handlers) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	hdlr.Router.HandleFunc(path, f).Methods("POST")
}

// Put wraps the router for PUT method
func (hdlr *Handlers) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	hdlr.Router.HandleFunc(path, f).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (hdlr *Handlers) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
	hdlr.Router.HandleFunc(path, f).Methods("DELETE")
}

type RequestHandlerFunction func(db *gorm.DB, w http.ResponseWriter, r *http.Request)

func (hdlr *Handlers) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(hdlr.db, w, r)
	}
}
