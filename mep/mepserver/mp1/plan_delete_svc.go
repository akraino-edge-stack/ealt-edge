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

	"github.com/apache/servicecomb-service-center/pkg/log"
	"github.com/apache/servicecomb-service-center/server/core"
	"github.com/apache/servicecomb-service-center/server/core/proto"

	"mepserver/mp1/arch/workspace"
)

type DeleteService struct {
	HttpErrInf *proto.Response `json:"httpErrInf,out"`
	workspace.TaskBase
	Ctx       context.Context `json:"ctx,in"`
	ServiceId string          `json:"serviceId,in"`
}

func (t *DeleteService) OnRequest(data string) workspace.TaskCode {
	if t.ServiceId == "" {
		t.SetFirstErrorCode(SerErrServiceDelFailed, "param is empty")
		return workspace.TaskFinish
	}
	serviceID := t.ServiceId[:len(t.ServiceId)/2]
	instanceID := t.ServiceId[len(t.ServiceId)/2:]
	req := &proto.UnregisterInstanceRequest{
		ServiceId:  serviceID,
		InstanceId: instanceID,
	}
	resp, err := core.InstanceAPI.Unregister(t.Ctx, req)
	if err != nil {
		log.Errorf(err, "Service delete failed!")
	}
	t.HttpErrInf = resp.Response

	return workspace.TaskFinish
}
