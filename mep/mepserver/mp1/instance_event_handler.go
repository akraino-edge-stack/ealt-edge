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
	"encoding/json"
	"strings"

	"github.com/apache/servicecomb-service-center/pkg/log"
	"github.com/apache/servicecomb-service-center/pkg/util"
	apt "github.com/apache/servicecomb-service-center/server/core"
	"github.com/apache/servicecomb-service-center/server/core/backend"
	"github.com/apache/servicecomb-service-center/server/core/proto"
	"github.com/apache/servicecomb-service-center/server/notify"
	"github.com/apache/servicecomb-service-center/server/plugin/pkg/discovery"
	"github.com/apache/servicecomb-service-center/server/plugin/pkg/registry"
	"github.com/apache/servicecomb-service-center/server/service/cache"
	"github.com/apache/servicecomb-service-center/server/service/metrics"
	svcutil "github.com/apache/servicecomb-service-center/server/service/util"
	"golang.org/x/net/context"

	"mepserver/mp1/models"
)

const indexStart, indexEnd = 33, 17

type InstanceEtsiEventHandler struct {
}

func (h *InstanceEtsiEventHandler) Type() discovery.Type {
	return backend.INSTANCE
}

func (h *InstanceEtsiEventHandler) OnEvent(evt discovery.KvEvent) {
	action := evt.Type
	instance, ok := evt.KV.Value.(*proto.MicroServiceInstance)
	if !ok {
		log.Warn("cast to instance failed")
	}
	providerId, providerInstanceId, domainProject := apt.GetInfoFromInstKV(evt.KV.Key)
	idx := strings.Index(domainProject, "/")
	domainName := domainProject[:idx]
	switch action {
	case proto.EVT_INIT:
		metrics.ReportInstances(domainName, 1)
		return
	case proto.EVT_CREATE:
		metrics.ReportInstances(domainName, 1)
	case proto.EVT_DELETE:
		metrics.ReportInstances(domainName, -1)
		if !apt.IsDefaultDomainProject(domainProject) {
			projectName := domainProject[idx+1:]
			svcutil.RemandInstanceQuota(util.SetDomainProject(context.Background(), domainName, projectName))
		}
	}

	if notify.NotifyCenter().Closed() {
		log.Warnf("caught [%s] instance [%s/%s] event, endpoints %v, but notify service is closed",
			action, providerId, providerInstanceId, instance.Endpoints)
		return
	}

	ctx := context.WithValue(context.WithValue(context.Background(),
		svcutil.CTX_CACHEONLY, "1"),
		svcutil.CTX_GLOBAL, "1")
	ms, err := svcutil.GetService(ctx, domainProject, providerId)
	if ms == nil {
		log.Errorf(err, "caught [%s] instance [%s/%s] event, endpoints %v, get cached provider's file failed",
			action, providerId, providerInstanceId, instance.Endpoints)
		return
	}

	log.Infof("caught [%s] service[%s][%s/%s/%s/%s] isntance[%s] event, endpoints %v",
		action, providerId, ms.Environment, ms.AppId, ms.ServiceName, ms.Version, providerInstanceId, instance.Endpoints)

	consumerIds := getCosumerIds()

	log.Infof("there are %d consuemrIDs, %s", len(consumerIds), consumerIds)
	PublishInstanceEvent(evt, domainProject, proto.MicroServiceToKey(domainProject, ms), consumerIds)
}

func getCosumerIds() []string {
	var consumerIds []string
	opts := []registry.PluginOp{
		registry.OpGet(registry.WithStrKey("/cse-sr/inst/files"), registry.WithPrefix()),
	}
	resp, err := backend.Registry().TxnWithCmp(context.Background(), opts, nil, nil)
	if err != nil {
		log.Errorf(err, "get subscription from etcd failed")
		return consumerIds
	}

	for _, kvs := range resp.Kvs {
		key := kvs.Key
		keystring := string(key)
		value := kvs.Value

		var mp1Req models.ServiceInfo
		err = json.Unmarshal(value, &mp1Req)
		if err != nil {
			log.Errorf(err, "parse serviceInfo failed")
		}
		length := len(keystring)
		keystring = keystring[length-indexStart : length-indexEnd]
		if StringContains(consumerIds, keystring) == -1 {
			consumerIds = append(consumerIds, keystring)
		}
	}
	return consumerIds
}

func NewInstanceEtsiEventHandler() *InstanceEtsiEventHandler {
	return &InstanceEtsiEventHandler{}
}

func PublishInstanceEvent(evt discovery.KvEvent, domainProject string, serviceKey *proto.MicroServiceKey, subscribers []string) {
	defer cache.FindInstances.Remove(serviceKey)
	if len(subscribers) == 0 {
		log.Warn("the subscribers size is 0")
		return
	}

	response := &proto.WatchInstanceResponse{
		Response: proto.CreateResponse(proto.Response_SUCCESS, "Watch instance successfully."),
		Action:   string(evt.Type),
		Key:      serviceKey,
		Instance: evt.KV.Value.(*proto.MicroServiceInstance),
	}
	for _, consumerId := range subscribers {
		job := notify.NewInstanceEventWithTime(consumerId, domainProject, evt.Revision, evt.CreateAt, response)
		err := notify.NotifyCenter().Publish(job)
		if err != nil {
			log.Errorf(err, "publish failed")
		}
	}
}
