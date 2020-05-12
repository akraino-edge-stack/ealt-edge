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
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/apache/servicecomb-service-center/pkg/log"
	"github.com/apache/servicecomb-service-center/server/core/backend"
	"github.com/apache/servicecomb-service-center/server/core/proto"
	"github.com/apache/servicecomb-service-center/server/notify"
	"github.com/apache/servicecomb-service-center/server/plugin/pkg/registry"

	"mepserver/mp1/models"
)

type Websocket struct {
	watcher         *notify.InstanceEventListWatcher
	ticker          *time.Ticker
	ctx             context.Context
	needPingWatcher bool
	free            chan struct{}
	closed          chan struct{}
	serviceID       string
}

func (ws *Websocket) Init() error {
	ws.ticker = time.NewTicker(notify.HeartbeatInterval)
	ws.needPingWatcher = true
	ws.free = make(chan struct{}, 1)
	ws.closed = make(chan struct{})
	ws.SetReady()
	if err := notify.NotifyCenter().AddSubscriber(ws.watcher); err != nil {
		return err
	}
	publisher.Accept(ws)
	return nil
}

func (ws *Websocket) ReadTimeout() time.Duration {
	return notify.ReadTimeout
}

func (ws *Websocket) SendTimeout() time.Duration {
	return notify.SendTimeout
}

func (ws *Websocket) HandleWatchWebSocketControlMessage() {

}

func (ws *Websocket) HandleWatchWebSocketJob(payload interface{}) {
	defer ws.SetReady()
	var (
		job *notify.InstanceEvent
	)
	switch v := payload.(type) {
	case error:
		err := payload.(error)
		log.Errorf(err, "watcher catch an error, subject: %s, group: %s", ws.watcher.Subject(), ws.watcher.Group())
	case time.Time:
		return
	case *notify.InstanceEvent:
		serviceID := ws.serviceID
		job = payload.(*notify.InstanceEvent)
		resp := job.Response
		SendMsgToApp(resp, serviceID)
	default:
		log.Errorf(nil, "watcher unknown input type %T, subject: %s, group: %s", v, ws.watcher.Subject(), ws.watcher.Group())
		return
	}

	select {
	case _, ok := <-ws.closed:
		if !ok {
			log.Warn("websocket channel closed")
		}
		return
	default:
	}
}

func (ws *Websocket) SetReady() {
	select {
	case ws.free <- struct{}{}:
	default:
	}

}

func (ws *Websocket) Pick() interface{} {
	select {
	case _, ok := <-ws.Ready():
		if !ok {
			log.Warn("websocket ready channel closed")
		}
		if ws.watcher.Err() != nil {
			return ws.watcher.Err()
		}

		select {
		case t, ok := <-ws.ticker.C:
			if !ok {
				log.Warn("websocket ticker C channel closed")
			}
			return t
		case j, ok := <-ws.watcher.Job:
			if !ok {
				log.Warn("websocket watcher job channel closed")
			}
			if j == nil {
				err := fmt.Errorf("server shutdown")
				log.Error("server shutdown", err)
			}
			return j
		default:
			ws.SetReady()
		}
	default:
	}
	return nil
}

func (ws *Websocket) Ready() chan struct{} {
	return ws.free
}

func (ws *Websocket) Stop() {
	close(ws.closed)
}

func getCallBackUris(serviceID string, instanceID string, serName string) []string {
	var callbackUris []string
	opts := []registry.PluginOp{
		registry.OpGet(registry.WithStrKey("/cse-sr/etsi/subscribe/"+serviceID), registry.WithPrefix()),
	}
	resp, err := backend.Registry().TxnWithCmp(context.Background(), opts, nil, nil)
	if err != nil {
		log.Errorf(err, "get subcription from etcd failed!")
		return callbackUris
	}
	for _, v := range resp.Kvs {
		var notifyInfo models.SerAvailabilityNotificationSubscription
		if v.Value == nil {
			log.Warn("the value is nil in etcd")
			continue
		}
		err = json.Unmarshal(v.Value, &notifyInfo)
		if err != nil {
			log.Warn("notify json can not be parsed to notifyInfo")
			continue
		}
		callbackURI := notifyInfo.CallbackReference
		filter := notifyInfo.FilteringCriteria

		if (len(filter.SerInstanceIds) != 0 && StringContains(filter.SerInstanceIds, instanceID) != -1) ||
			(len(filter.SerNames) != 0 && StringContains(filter.SerNames, serviceID) != -1) {
			callbackUris = append(callbackUris, callbackURI)
		}
	}
	log.Infof("send to consumerIds: %s", callbackUris)
	return callbackUris
}

func StringContains(arr []string, val string) (index int) {
	index = -1
	for i := 0; i < len(arr); i++ {
		if arr[i] == val {
			index = i
			return
		}
	}
	return
}

func SendMsgToApp(data *proto.WatchInstanceResponse, serviceID string) {
	// transfer data to instanceInfo, and get instaceid, serviceName
	instanceID := data.Instance.ServiceId + data.Instance.InstanceId
	serName := data.Instance.Properties["serName"]
	action := data.Action
	instanceInfo := data.Instance
	instanceStr, err := json.Marshal(instanceInfo)
	if err != nil {
		log.Errorf(err, "parse instanceInfo failed!")
		return
	}

	callbackUris := getCallBackUris(serviceID, instanceID, serName)
	body := strings.NewReader(string(instanceStr))
	doSend(action, body, callbackUris)
}

func doSend(action string, body io.Reader, callbackUris []string) {
	for _, callbackURI := range callbackUris {
		log.Debugf("action: %s with callbackURI:%s", action, callbackURI)
		if !strings.HasPrefix(callbackURI, "http") {
			callbackURI = "http://" + callbackURI
		}
		client := http.Client{}

		if action == "CREATE" {
			contentType := "application/x-www-form-urlencoded"
			_, err := http.Post(callbackURI, contentType, body)
			if err != nil {
				log.Warn("the consumer handle post action failed!")
			}
		} else if action == "DELETE" {
			req, err := http.NewRequest("delete", callbackURI, body)
			if err != nil {
				_, err := client.Do(req)
				if err != nil {
					log.Warn("the consumer handle delete action failed!")
				}

			} else {
				log.Errorf(err, "crate request failed!")
			}
		} else if action == "UPDATE" {
			req, err := http.NewRequest("put", callbackURI, body)
			if err != nil {
				_, err := client.Do(req)
				if err != nil {
					log.Warn("the consumer handle update action failed!")
				}

			} else {
				log.Errorf(err, "crate request failed!")
			}
		}
	}
}

func DoWebSocketListAndWatchV2(ctx context.Context, id string, id2 string, f func() ([]*proto.WatchInstanceResponse, int64)) {
	//TBD
}
