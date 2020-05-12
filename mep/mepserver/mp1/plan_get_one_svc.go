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

	"github.com/apache/servicecomb-service-center/pkg/util"
	"github.com/apache/servicecomb-service-center/server/core"
	"github.com/apache/servicecomb-service-center/server/core/proto"

	"mepserver/mp1/arch/workspace"
	"mepserver/mp1/models"
	meputil "mepserver/mp1/util"
)

type GetOneDecode struct {
	workspace.TaskBase
	R           *http.Request   `json:"r,in"`
	Ctx         context.Context `json:"ctx,out"`
	CoreRequest interface{}     `json:"coreRequest,out"`
}

func (t *GetOneDecode) OnRequest(data string) workspace.TaskCode {
	t.Ctx, t.CoreRequest = t.getFindParam(t.R)
	return workspace.TaskFinish

}

func (t *GetOneDecode) getFindParam(r *http.Request) (context.Context, *proto.GetOneInstanceRequest) {
	query, ids := meputil.GetHTTPTags(r)
	mp1SrvId := query.Get(":serviceId")
	serviceId := mp1SrvId[:len(mp1SrvId)/2]
	instanceId := mp1SrvId[len(mp1SrvId)/2:]
	req := &proto.GetOneInstanceRequest{
		ConsumerServiceId:  r.Header.Get("X-ConsumerId"),
		ProviderServiceId:  serviceId,
		ProviderInstanceId: instanceId,
		Tags:               ids,
	}

	ctx := util.SetTargetDomainProject(r.Context(), r.Header.Get("X-Domain-Name"), query.Get(":project"))
	return ctx, req
}

type GetOneInstance struct {
	workspace.TaskBase
	HttpErrInf  *proto.Response `json:"httpErrInf,out"`
	Ctx         context.Context `json:"ctx,in"`
	CoreRequest interface{}     `json:"coreRequest,in"`
	HttpRsp     interface{}     `json:"httpRsp,out"`
}

func (t *GetOneInstance) OnRequest(data string) workspace.TaskCode {
	resp, _ := core.InstanceAPI.GetOneInstance(t.Ctx, t.CoreRequest.(*proto.GetOneInstanceRequest))
	t.HttpErrInf = resp.Response
	resp.Response = nil
	mp1Rsp := &models.ServiceInfo{}
	if resp.Instance != nil {
		mp1Rsp.FromServiceInstance(resp.Instance)
	} else {
		t.SetFirstErrorCode(SerInstanceNotFound, "service instance id not found")
		return workspace.TaskFinish
	}
	t.HttpRsp = mp1Rsp

	return workspace.TaskFinish
}
