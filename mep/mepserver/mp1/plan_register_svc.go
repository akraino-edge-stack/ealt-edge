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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/log"
	"github.com/apache/servicecomb-service-center/pkg/util"
	"github.com/apache/servicecomb-service-center/server/core"
	"github.com/apache/servicecomb-service-center/server/core/proto"
	svcerr "github.com/apache/servicecomb-service-center/server/error"

	"mepserver/mp1/arch/workspace"
	"mepserver/mp1/models"
	meputil "mepserver/mp1/util"
)

type DecodeRestReq struct {
	workspace.TaskBase
	R             *http.Request   `json:"r,in"`
	Ctx           context.Context `json:"ctx,out"`
	AppInstanceId string          `json:"appInstanceId,out"`
	SubscribeId   string          `json:"subscribeId,out"`
	ServiceId     string          `json:"serviceId,out"`
	RestBody      interface{}     `json:"restBody,out"`
}

func (t *DecodeRestReq) OnRequest(data string) workspace.TaskCode {
	t.GetParam(t.R)
	err := t.ParseBody(t.R)
	if err != nil {
		log.Error("parse rest body failed", err)
	}
	return workspace.TaskFinish
}

func (t *DecodeRestReq) ParseBody(r *http.Request) error {
	if t.RestBody == nil {
		return nil
	}
	msg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("read body failed", err)
		t.SetFirstErrorCode(SerErrFailBase, err.Error())
		return err
	}

	newMsg, err := t.checkParam(msg)
	if err != nil {
		log.Error("check Param failed", err)
		t.SetFirstErrorCode(SerErrFailBase, err.Error())
		return err
	}

	err = json.Unmarshal(newMsg, t.RestBody)
	if err != nil {
		log.Errorf(err, "invalid json: %s", util.BytesToStringWithNoCopy(newMsg))
		t.SetFirstErrorCode(SerErrFailBase, err.Error())
		return err
	}
	return nil

}

func (t *DecodeRestReq) checkParam(msg []byte) ([]byte, error) {

	var temp map[string]interface{}
	err := json.Unmarshal(msg, &temp)
	if err != nil {
		log.Errorf(err, "invalid json to map: %s", util.BytesToStringWithNoCopy(msg))
		t.SetFirstErrorCode(SerErrFailBase, err.Error())
		return nil, err
	}

	meputil.SetMapValue(temp, "consumedLocalOnly", true)
	meputil.SetMapValue(temp, "isLocal", true)
	meputil.SetMapValue(temp, "scopeOfLocality", "MEC_HOST")

	msg, err = json.Marshal(&temp)
	if err != nil {
		log.Errorf(err, "invalid map to json")
		t.SetFirstErrorCode(SerErrFailBase, err.Error())
		return nil, err
	}

	return msg, nil
}

func (t *DecodeRestReq) WithBody(body interface{}) *DecodeRestReq {
	t.RestBody = body
	return t
}

func (t *DecodeRestReq) GetParam(r *http.Request) {
	query, _ := meputil.GetHTTPTags(r)
	t.AppInstanceId = query.Get(":appInstanceId")
	t.SubscribeId = query.Get(":subscriptionId")
	t.ServiceId = query.Get(":serviceId")
	t.Ctx = util.SetTargetDomainProject(r.Context(), r.Header.Get("X-Domain-Name"), query.Get(":project"))
}

type RegisterServiceId struct {
	HttpErrInf *proto.Response `json:"httpErrInf,out"`
	workspace.TaskBase
	Ctx       context.Context `json:"ctx,in"`
	ServiceId string          `json:"serviceId,out"`
	RestBody  interface{}     `json:"restBody,in"`
}

func (t *RegisterServiceId) OnRequest(data string) workspace.TaskCode {

	serviceInfo, ok := t.RestBody.(*models.ServiceInfo)
	if !ok {
		t.SetFirstErrorCode(1, "restbody failed")
		return workspace.TaskFinish
	}
	req := &proto.CreateServiceRequest{}
	serviceInfo.ToServiceRequest(req)
	resp, err := core.ServiceAPI.Create(t.Ctx, req)
	if err != nil {
		log.Errorf(err, "Service Center ServiceAPI.Create fail: %s!", err.Error())
		t.SetFirstErrorCode(1, err.Error())
		return workspace.TaskFinish
	}

	if resp.ServiceId == "" {
		t.HttpErrInf = resp.Response
		log.Warn("Service id empty.")
	}
	t.ServiceId = resp.ServiceId
	return workspace.TaskFinish
}

type RegisterServiceInst struct {
	HttpErrInf *proto.Response `json:"httpErrInf,out"`
	workspace.TaskBase
	W             http.ResponseWriter `json:"w,in"`
	Ctx           context.Context     `json:"ctx,in"`
	AppInstanceId string              `json:"appInstanceId,in"`
	ServiceId     string              `json:"serviceId,in"`
	InstanceId    string              `json:"instanceId,out"`
	RestBody      interface{}         `json:"restBody,in"`
	HttpRsp       interface{}         `json:"httpRsp,out"`
}

func (t *RegisterServiceInst) OnRequest(data string) workspace.TaskCode {
	serviceInfo, ok := t.RestBody.(*models.ServiceInfo)
	if !ok {
		t.SetFirstErrorCode(1, "restbody failed")
		return workspace.TaskFinish
	}
	req := &proto.RegisterInstanceRequest{}
	serviceInfo.ToRegisterInstance(req)
	req.Instance.ServiceId = t.ServiceId
	req.Instance.Properties["appInstanceId"] = t.AppInstanceId
	resp, err := core.InstanceAPI.Register(t.Ctx, req)
	if err != nil {
		log.Errorf(err, "RegisterInstance fail: %s", t.ServiceId)
		t.HttpErrInf = &proto.Response{}
		t.HttpErrInf.Code = svcerr.ErrForbidden
		t.HttpErrInf.Message = err.Error()
		return workspace.TaskFinish
	}
	t.InstanceId = resp.InstanceId

	//build response serviceComb use serviceId + InstanceId to mark a service instance
	mp1SerId := t.ServiceId + t.InstanceId
	serviceInfo.SerInstanceId = mp1SerId
	t.HttpRsp = serviceInfo

	location := fmt.Sprintf("/mep/mp1/v1/services/%s", mp1SerId)
	t.W.Header().Set("Location", location)
	return workspace.TaskFinish
}
