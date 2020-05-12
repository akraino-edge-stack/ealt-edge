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
	"net/http"

	"github.com/apache/servicecomb-service-center/pkg/log"
	"github.com/apache/servicecomb-service-center/server/core/backend"
	"github.com/apache/servicecomb-service-center/server/core/proto"
	"github.com/apache/servicecomb-service-center/server/plugin/pkg/registry"

	"mepserver/mp1/arch/workspace"
	"mepserver/mp1/models"
)

type GetSubscribes struct {
	workspace.TaskBase
	R             *http.Request       `json:"r,in"`
	HttpErrInf    *proto.Response     `json:"httpErrInf,out"`
	Ctx           context.Context     `json:"ctx,in"`
	W             http.ResponseWriter `json:"w,in"`
	AppInstanceId string              `json:"appInstanceId,in"`
	HttpRsp       interface{}         `json:"httpRsp,out"`
}

func (t *GetSubscribes) OnRequest(data string) workspace.TaskCode {

	appInstanceId := t.AppInstanceId

	opts := []registry.PluginOp{
		registry.OpGet(registry.WithStrKey("/cse-sr/etsi/subscribe/"+appInstanceId), registry.WithPrefix()),
	}
	resp, err := backend.Registry().TxnWithCmp(context.Background(), opts, nil, nil)
	if err != nil {
		log.Errorf(err, "get subscription from etcd failed")
		t.SetFirstErrorCode(OperateDataWithEtcdErr, "get subscription from etcd failed")
		return workspace.TaskFinish
	}

	var values []string
	for _, value := range resp.Kvs {
		values = append(values, string(value.Value))
	}
	if len(values) == 0 {
		log.Errorf(err, "get subscription failed, subscription not exist")
		t.SetFirstErrorCode(SubscriptionNotFound, "get subscription failed, subscription not exist")
		return workspace.TaskFinish
	}
	subs := make([]*models.SerAvailabilityNotificationSubscription, 0, len(values))
	for _, val := range values {
		sub := &models.SerAvailabilityNotificationSubscription{}
		err := json.Unmarshal([]byte(val), sub)
		if err != nil {
			log.Warn("subscribe can not be parsed to SerAvailabilityNotificationSubscription")
			continue
		}
		subs = append(subs, sub)
	}
	t.HttpRsp = subs
	return workspace.TaskFinish
}
