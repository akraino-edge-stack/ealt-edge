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

	"github.com/apache/servicecomb-service-center/pkg/util"
	"github.com/apache/servicecomb-service-center/server/core/proto"
	svcutil "github.com/apache/servicecomb-service-center/server/service/util"

	"mepserver/mp1/arch/workspace"
	"mepserver/mp1/models"
	meputil "mepserver/mp1/util"
)

type UpdateInstance struct {
	workspace.TaskBase
	HttpErrInf *proto.Response `json:"httpErrInf,out"`
	Ctx        context.Context `json:"ctx,in"`
	ServiceId  string          `json:"serviceId,in"`
	RestBody   interface{}     `json:"restBody,in"`
	HttpRsp    interface{}     `json:"httpRsp,out"`
}

func (t *UpdateInstance) OnRequest(data string) workspace.TaskCode {
	if t.ServiceId == "" {
		t.SetFirstErrorCode(SerErrFailBase, "param is empty")
		return workspace.TaskFinish
	}
	mp1Ser, ok := t.RestBody.(*models.ServiceInfo)
	if !ok {
		t.SetFirstErrorCode(SerErrFailBase, "body invalid")
		return workspace.TaskFinish
	}

	instance, err := meputil.GetServiceInstance(t.Ctx, t.ServiceId)
	if err != nil {
		t.SetFirstErrorCode(SerInstanceNotFound, "find service failed")
		return workspace.TaskFinish
	}

	copyInstanceRef := *instance
	req := proto.RegisterInstanceRequest{
		Instance: &copyInstanceRef,
	}
	mp1Ser.ToRegisterInstance(&req)

	domainProject := util.ParseDomainProject(t.Ctx)
	centerErr := svcutil.UpdateInstance(t.Ctx, domainProject, &copyInstanceRef)
	if centerErr != nil {
		t.SetFirstErrorCode(SerErrServiceUpdFailed, "update service failed")
		return workspace.TaskFinish
	}

	err = meputil.Heartbeat(t.Ctx, mp1Ser.SerInstanceId)
	if err != nil {
		t.SetFirstErrorCode(SerErrServiceUpdFailed, "heartbeat failed")
		return workspace.TaskFinish
	}
	mp1Ser.SerInstanceId = instance.ServiceId + instance.InstanceId
	t.HttpRsp = mp1Ser
	t.HttpErrInf = proto.CreateResponse(proto.Response_SUCCESS, "Update service instance success.")
	return workspace.TaskFinish
}
