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
	"net/url"

	"github.com/apache/servicecomb-service-center/server/core"
	"github.com/apache/servicecomb-service-center/server/core/proto"

	"mepserver/mp1/arch/workspace"
	"mepserver/mp1/models"
	meputil "mepserver/mp1/util"
)

type DiscoverDecode struct {
	workspace.TaskBase
	R           *http.Request   `json:"r,in"`
	Ctx         context.Context `json:"ctx,out"`
	QueryParam  url.Values      `json:"queryParam,out"`
	CoreRequest interface{}     `json:"coreRequest,out"`
}

func (t *DiscoverDecode) OnRequest(data string) workspace.TaskCode {
	t.Ctx, t.CoreRequest, t.QueryParam = meputil.GetFindParam(t.R)
	return workspace.TaskFinish
}

type DiscoverService struct {
	workspace.TaskBase
	Ctx         context.Context `json:"ctx,in"`
	QueryParam  url.Values      `json:"queryParam,in"`
	CoreRequest interface{}     `json:"coreRequest,in"`
	CoreRsp     interface{}     `json:"coreRsp,out"`
}

func (t *DiscoverService) checkInstanceId(req *proto.FindInstancesRequest) bool {
	instanceId := req.AppId
	if instanceId != "default" {
		instances := t.CoreRsp.(*proto.FindInstancesResponse).Instances
		for _, val := range instances {
			if val.ServiceId+val.InstanceId == instanceId {
				return true
			}
		}
		return false
	}
	return true
}

func (t *DiscoverService) OnRequest(data string) workspace.TaskCode {
	req, ok := t.CoreRequest.(*proto.FindInstancesRequest)
	if !ok {
		t.SetFirstErrorCode(SerErrServiceNotFound, "cast to request fail")
		return workspace.TaskFinish
	}
	if req.ServiceName == "" {
		var errFindByKey error
		t.CoreRsp, errFindByKey = meputil.FindInstanceByKey(t.QueryParam)
		if errFindByKey != nil || t.CoreRsp == nil {
			t.SetFirstErrorCode(SerErrServiceNotFound, errFindByKey.Error())
			return workspace.TaskFinish
		}
		if !t.checkInstanceId(req) {
			t.SetFirstErrorCode(SerErrServiceNotFound, "instance id not found")
		}
		return workspace.TaskFinish
	}

	findInstance, err := core.InstanceAPI.Find(t.Ctx, req)
	if err != nil {
		t.SetFirstErrorCode(SerErrServiceNotFound, err.Error())
		return workspace.TaskFinish
	}
	if findInstance == nil || len(findInstance.Instances) == 0 {
		t.SetFirstErrorCode(SerErrServiceNotFound, "service not found")
		return workspace.TaskFinish
	}

	t.CoreRsp = findInstance
	return workspace.TaskFinish
}

type ToStrDiscover struct {
	HttpErrInf *proto.Response `json:"httpErrInf,out"`
	workspace.TaskBase
	CoreRsp interface{} `json:"coreRsp,in"`
	HttpRsp interface{} `json:"httpRsp,out"`
}

func (t *ToStrDiscover) OnRequest(data string) workspace.TaskCode {
	t.HttpErrInf, t.HttpRsp = mp1CvtSrvDiscover(t.CoreRsp.(*proto.FindInstancesResponse))
	return workspace.TaskFinish
}

type RspHook struct {
	R *http.Request `json:"r,in"`
	workspace.TaskBase
	Ctx     context.Context `json:"ctx,in"`
	HttpRsp interface{}     `json:"httpRsp,in"`
	HookRsp interface{}     `json:"hookRsp,out"`
}

func (t *RspHook) OnRequest(data string) workspace.TaskCode {
	t.HookRsp = instanceHook(t.Ctx, t.R, t.HttpRsp)
	return workspace.TaskFinish
}

func instanceHook(ctx context.Context, r *http.Request, rspData interface{}) interface{} {
	rspBody, ok := rspData.([]*models.ServiceInfo)
	if !ok {
		return rspData
	}

	if len(rspBody) == 0 {
		return rspBody
	}
	consumerName := r.Header.Get("X-ConsumerName")
	if consumerName == "APIGW" {
		return rspBody
	}

	for _, v := range rspBody {
		if apihook.APIHook != nil {
			info := apihook.APIHook()
			if len(info.Addresses) == 0 && len(info.Uris) == 0 {
				return rspBody
			}
			v.TransportInfo.Endpoint = info
		}
	}
	return rspBody
}

type SendHttpRsp struct {
	HttpErrInf *proto.Response `json:"httpErrInf,in"`
	workspace.TaskBase
	W       http.ResponseWriter `json:"w,in"`
	HttpRsp interface{}         `json:"httpRsp,in"`
}

func (t *SendHttpRsp) OnRequest(data string) workspace.TaskCode {
	errInfo := t.GetSerErrInfo()
	if errInfo.ErrCode >= int(workspace.TaskFail) {
		statusCode, httpBody := t.cvtHttpErrInfo(errInfo)
		meputil.HttpErrResponse(t.W, statusCode, httpBody)

		return workspace.TaskFinish
	}
	meputil.WriteResponse(t.W, t.HttpErrInf, t.HttpRsp)
	return workspace.TaskFinish
}

func (t *SendHttpRsp) cvtHttpErrInfo(errInfo *workspace.SerErrInfo) (int, interface{}) {
	statusCode := http.StatusBadRequest
	var httpBody interface{}
	switch workspace.ErrCode(errInfo.ErrCode) {
	case SerErrServiceNotFound:
		{
			//status should return bad request
			body := &models.ProblemDetails{
				Title:  "Can not found resource",
				Status: uint32(errInfo.ErrCode),
				Detail: errInfo.Message,
			}
			httpBody = body
		}
	case SerInstanceNotFound:
		{
			statusCode = http.StatusNotFound
			body := &models.ProblemDetails{
				Title:  "Can not found resource",
				Status: uint32(errInfo.ErrCode),
				Detail: errInfo.Message,
			}
			httpBody = body
		}
	}

	return statusCode, httpBody
}

func mp1CvtSrvDiscover(findInsResp *proto.FindInstancesResponse) (*proto.Response, []*models.ServiceInfo) {
	resp := findInsResp.Response
	if resp != nil && resp.GetCode() != proto.Response_SUCCESS {
		return resp, nil
	}
	serviceInfos := make([]*models.ServiceInfo, 0, len(findInsResp.Instances))
	for _, ins := range findInsResp.Instances {
		serviceInfo := &models.ServiceInfo{}
		serviceInfo.FromServiceInstance(ins)
		serviceInfos = append(serviceInfos, serviceInfo)
	}
	return resp, serviceInfos

}
