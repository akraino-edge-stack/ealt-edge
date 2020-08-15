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

package mp1

import (
	"context"
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/log"
	"github.com/apache/servicecomb-service-center/pkg/rest"
	"github.com/apache/servicecomb-service-center/pkg/util"
	"github.com/apache/servicecomb-service-center/server/core/proto"
	"github.com/apache/servicecomb-service-center/server/notify"
	v4 "github.com/apache/servicecomb-service-center/server/rest/controller/v4"
	svcutil "github.com/apache/servicecomb-service-center/server/service/util"

	"mepserver/mp1/arch/workspace"
	"mepserver/mp1/models"
)

const (
	basePath         = "/mep/mec_service_mgmt/v1"
	servicesPath     = basePath + "/services"
	appServicesPath  = basePath + "/applications/:appInstanceId" + "/services"
	appSubscribePath = basePath + "/applications/:appInstanceId/subscriptions"
)

const (
	SerErrFailBase         workspace.ErrCode = workspace.TaskFail
	SerErrServiceNotFound                    = 2
	SerInstanceNotFound                      = 3
	ParseInfoErr                             = 4
	SubscriptionNotFound                     = 5
	OperateDataWithEtcdErr                   = 6
	SerErrServiceDelFailed                   = 7
	SerErrServiceUpdFailed                   = 8
)

type APIHookFunc func() models.EndPointInfo

type APIGwHook struct {
	APIHook APIHookFunc
}

var apihook APIGwHook

func SetAPIHook(hook APIGwHook) {
	apihook = hook
}

func GetApiHook() APIGwHook {
	return apihook
}

func init() {
	initRouter()
}

func initRouter() {
	rest.
		RegisterServant(&Mp1Service{})
}

type Mp1Service struct {
	v4.MicroServiceService
}

func (m *Mp1Service) URLPatterns() []rest.Route {
	return []rest.Route{
		// appSubscriptions
		{Method: rest.HTTP_METHOD_POST, Path: appSubscribePath, Func: doAppSubscribe},
		{Method: rest.HTTP_METHOD_GET, Path: appSubscribePath, Func: getAppSubscribes},
		{Method: rest.HTTP_METHOD_GET, Path: appSubscribePath + "/:subscriptionId", Func: getOneAppSubscribe},
		{Method: rest.HTTP_METHOD_DELETE, Path: appSubscribePath + "/:subscriptionId", Func: delOneAppSubscribe},
		// appServices
		{Method: rest.HTTP_METHOD_POST, Path: appServicesPath, Func: serviceRegister},
		{Method: rest.HTTP_METHOD_GET, Path: appServicesPath, Func: serviceDiscover},
		{Method: rest.HTTP_METHOD_PUT, Path: appServicesPath + "/:serviceId", Func: serviceUpdate},
		{Method: rest.HTTP_METHOD_GET, Path: appServicesPath + "/:serviceId", Func: getOneService},
		{Method: rest.HTTP_METHOD_DELETE, Path: appServicesPath + "/:serviceId", Func: serviceDelete},
		// services
		{Method: rest.HTTP_METHOD_GET, Path: servicesPath, Func: serviceDiscover},
		{Method: rest.HTTP_METHOD_GET, Path: servicesPath + "/:serviceId", Func: getOneService},
	}
}

//application subscription
func doAppSubscribe(w http.ResponseWriter, r *http.Request) {

	workPlan := NewWorkSpace(w, r)
	workPlan.Try(
		(&DecodeRestReq{}).WithBody(&models.SerAvailabilityNotificationSubscription{}),
		&SubscribeIst{})
	workPlan.Finally(&SendHttpRspCreated{})

	workspace.WkRun(workPlan)
}

func getAppSubscribes(w http.ResponseWriter, r *http.Request) {

	workPlan := NewWorkSpace(w, r)
	workPlan.Try(
		&DecodeRestReq{},
		&GetSubscribes{})
	workPlan.Finally(&SendHttpRsp{})

	workspace.WkRun(workPlan)
}

func getOneAppSubscribe(w http.ResponseWriter, r *http.Request) {

	workPlan := NewWorkSpace(w, r)
	workPlan.Try(
		&DecodeRestReq{},
		&GetOneSubscribe{})
	workPlan.Finally(&SendHttpRsp{})

	workspace.WkRun(workPlan)
}

func delOneAppSubscribe(w http.ResponseWriter, r *http.Request) {

	workPlan := NewWorkSpace(w, r)
	workPlan.Try(
		&DecodeRestReq{},
		&DelOneSubscribe{})
	workPlan.Finally(&SendHttpRsp{})

	workspace.WkRun(workPlan)
}

//service registery request
func serviceRegister(w http.ResponseWriter, r *http.Request) {
	log.Info("Register service start...")

	workPlan := NewWorkSpace(w, r)
	workPlan.Try(
		(&DecodeRestReq{}).WithBody(&models.ServiceInfo{}),
		&RegisterServiceId{},
		&RegisterServiceInst{})
	workPlan.Finally(&SendHttpRspCreated{})

	workspace.WkRun(workPlan)
}

func serviceDiscover(w http.ResponseWriter, r *http.Request) {
	log.Info("Discover service service start...")

	workPlan := NewWorkSpace(w, r)
	workPlan.Try(
		&DiscoverDecode{},
		&DiscoverService{},
		&ToStrDiscover{},
		&RspHook{})
	workPlan.Finally(&SendHttpRsp{})

	workspace.WkRun(workPlan)
}

func serviceUpdate(w http.ResponseWriter, r *http.Request) {
	log.Info("Update a service start...")

	workPlan := NewWorkSpace(w, r)
	workPlan.Try(
		(&DecodeRestReq{}).WithBody(&models.ServiceInfo{}),
		&UpdateInstance{})
	workPlan.Finally(&SendHttpRsp{})

	workspace.WkRun(workPlan)
}

func getOneService(w http.ResponseWriter, r *http.Request) {
	log.Info("Register service start...")

	workPlan := NewWorkSpace(w, r)
	workPlan.Try(
		&GetOneDecode{},
		&GetOneInstance{})
	workPlan.Finally(&SendHttpRsp{})

	workspace.WkRun(workPlan)

}

func serviceDelete(w http.ResponseWriter, r *http.Request) {
	log.Info("Delete a service start...")

	workPlan := NewWorkSpace(w, r)
	workPlan.Try(
		&DecodeRestReq{},
		&DeleteService{})
	workPlan.Finally(&SendHttpRsp{})

	workspace.WkRun(workPlan)
}

func WebsocketListAndWatch(ctx context.Context, req *proto.WatchInstanceRequest, consumerSvcId string) {
	if req == nil || len(req.SelfServiceId) == 0 {
		log.Warn("request fomat invalid！")
		return
	}
	domainProject := util.ParseDomainProject(ctx)
	if !svcutil.ServiceExist(ctx, domainProject, req.SelfServiceId) {
		log.Warn("service does not exist!")
		return
	}
	DoWebsocketListAndWatch(ctx, req.SelfServiceId, consumerSvcId, func() ([]*proto.WatchInstanceResponse, int64) {
		return svcutil.QueryAllProvidersInstances(ctx, req.SelfServiceId)
	})
}

func DoWebsocketListAndWatch(ctx context.Context, serviceId string, consumerSvcId string, f func() ([]*proto.WatchInstanceResponse, int64)) {
	domainProject := util.ParseDomainProject(ctx)
	socket := &Websocket{
		ctx:       ctx,
		watcher:   notify.NewInstanceEventListWatcher(serviceId, domainProject, f),
		serviceID: consumerSvcId,
	}
	ProcessSocket(socket)
}

func ProcessSocket(socket *Websocket) {
	if err := socket.Init(); err != nil {
		return
	}
	socket.HandleWatchWebSocketControlMessage()
}
