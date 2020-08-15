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
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/log"
	"github.com/apache/servicecomb-service-center/server/core/backend"
	"github.com/apache/servicecomb-service-center/server/core/proto"
	"github.com/apache/servicecomb-service-center/server/plugin/pkg/registry"
	"github.com/satori/go.uuid"

	"mepserver/mp1/arch/workspace"
	"mepserver/mp1/models"
)

type SubscribeIst struct {
	workspace.TaskBase
	R             *http.Request       `json:"r,in"`
	HttpErrInf    *proto.Response     `json:"httpErrInf,out"`
	Ctx           context.Context     `json:"ctx,in"`
	W             http.ResponseWriter `json:"w,in"`
	RestBody      interface{}         `json:"restBody,in"`
	AppInstanceId string              `json:"appInstanceId,in"`
	SubscribeId   string              `json:"subscribeId,in"`
	HttpRsp       interface{}         `json:"httpRsp,out"`
}

//service subscription request
func (t *SubscribeIst) OnRequest(data string) workspace.TaskCode {

	mp1SubscribeInfo, ok := t.RestBody.(*models.SerAvailabilityNotificationSubscription)
	if !ok {
		t.SetFirstErrorCode(SerErrFailBase, "restBody failed")
		return workspace.TaskFinish
	}

	appInstanceId := t.AppInstanceId
	subscribeId := uuid.NewV4().String()
	t.SubscribeId = subscribeId
	subscribeJSON, err := json.Marshal(mp1SubscribeInfo)
	if err != nil {
		log.Errorf(err, "can not Marshal subscribe info")
		t.SetFirstErrorCode(ParseInfoErr, "can not marshal subscribe info")
		return workspace.TaskFinish
	}
	opts := []registry.PluginOp{
		registry.OpPut(registry.WithStrKey("/cse-sr/etsi/subscribe/"+appInstanceId+"/"+subscribeId),
			           registry.WithValue(subscribeJSON)),
	}
	_, resultErr := backend.Registry().TxnWithCmp(context.Background(), opts, nil, nil)
	if resultErr != nil {
		log.Errorf(err, "subscription to etcd failed!")
		t.SetFirstErrorCode(OperateDataWithEtcdErr, "put subscription to etcd failed")
		return workspace.TaskFinish
	}

	req := &proto.WatchInstanceRequest{SelfServiceId: appInstanceId[:len(appInstanceId)/2]}
	t.R.Method = "WATCHLIST"
	WebsocketListAndWatch(t.Ctx, req, appInstanceId)
	t.buildResponse(mp1SubscribeInfo)

	return workspace.TaskFinish
}

func (t *SubscribeIst) buildResponse(sub *models.SerAvailabilityNotificationSubscription) {
	appInstanceID := t.AppInstanceId
	subscribeID := t.SubscribeId

	t.HttpRsp = sub
	location := fmt.Sprintf("/mec_service_mgmt/v1/applications/%s/subscriptions/%s/", appInstanceID, subscribeID)
	t.W.Header().Set("Location", location)
}
