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

package util

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/apache/servicecomb-service-center/pkg/log"
	"github.com/apache/servicecomb-service-center/pkg/rest"
	"github.com/apache/servicecomb-service-center/pkg/util"
	"github.com/apache/servicecomb-service-center/server/core"
	"github.com/apache/servicecomb-service-center/server/core/backend"
	"github.com/apache/servicecomb-service-center/server/core/proto"
	svcerror "github.com/apache/servicecomb-service-center/server/error"
	"github.com/apache/servicecomb-service-center/server/plugin/pkg/registry"
	"github.com/apache/servicecomb-service-center/server/rest/controller"
	svcutil "github.com/apache/servicecomb-service-center/server/service/util"
)

func InfoToProperties(properties map[string]string, key string, value string) {
	if value != "" {
		properties[key] = value
	}
}

func JsonTextToObj(jsonText string) (interface{}, error) {
	data := []byte(jsonText)
	var jsonMap interface{}
	decoder := json.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(&jsonMap)
	if err != nil {
		return nil, err
	}
	return jsonMap, nil
}

func GetHostPort(uri string) (string, int) {
	idx := strings.LastIndex(uri, ":")
	domain := uri
	port := 0
	var err error
	if idx > 0 {
		port, err = strconv.Atoi(uri[idx+1:])
		if err != nil {
			port = 0
		}
		domain = uri[:idx]
	}
	return domain, port
}

func GetHTTPTags(r *http.Request) (url.Values, []string) {
	var ids []string
	query := r.URL.Query()
	keys := query.Get("tags")
	if len(keys) > 0 {
		ids = strings.Split(keys, ",")
	}

	return query, ids
}

func GetFindParam(r *http.Request) (context.Context, *proto.FindInstancesRequest, url.Values) {

	query, ids := GetHTTPTags(r)

	req := &proto.FindInstancesRequest{
		ConsumerServiceId: r.Header.Get("X-ConsumerId"),
		AppId:             query.Get("instance_id"),
		ServiceName:       query.Get("ser_name"),
		VersionRule:       query.Get("version"),
		Environment:       query.Get("env"),
		Tags:              ids,
	}

	if req.AppId == "" {
		req.AppId = "default"
	}
	if req.VersionRule == "" {
		req.VersionRule = "latest"
	}
	ctx := util.SetTargetDomainProject(r.Context(), r.Header.Get("X-Domain-Name"), query.Get(":project"))
	return ctx, req, query
}

//send http response
func WriteHTTPResponse(w http.ResponseWriter, resp *proto.Response, obj interface{}) {
	if resp != nil && resp.GetCode() != proto.Response_SUCCESS {
		controller.WriteError(w, resp.GetCode(), resp.GetMessage())
		return
	}
	if obj == nil {
		w.Header().Set(rest.HEADER_RESPONSE_STATUS, strconv.Itoa(http.StatusOK))
		w.Header().Set(rest.HEADER_CONTENT_TYPE, rest.CONTENT_TYPE_TEXT)
		w.WriteHeader(http.StatusOK)
		return
	}

	objJSON, err := json.Marshal(obj)
	if err != nil {
		controller.WriteError(w, svcerror.ErrInternal, err.Error())
		return
	}
	w.Header().Set(rest.HEADER_RESPONSE_STATUS, strconv.Itoa(http.StatusOK))
	w.Header().Set(rest.HEADER_CONTENT_TYPE, rest.CONTENT_TYPE_JSON)
	w.WriteHeader(http.StatusCreated)
	_, err = fmt.Fprintln(w, util.BytesToStringWithNoCopy(objJSON))
	if err != nil {
		return
	}
}

func WriteResponse(w http.ResponseWriter, resp *proto.Response, obj interface{}) {
	if resp != nil && resp.GetCode() != proto.Response_SUCCESS {
		controller.WriteError(w, resp.GetCode(), resp.GetMessage())
		return
	}
	if obj == nil {
		w.Header().Set(rest.HEADER_RESPONSE_STATUS, strconv.Itoa(http.StatusOK))
		w.Header().Set(rest.HEADER_CONTENT_TYPE, rest.CONTENT_TYPE_TEXT)
		w.WriteHeader(http.StatusOK)
		return
	}

	objJSON, err := json.Marshal(obj)
	if err != nil {
		controller.WriteError(w, svcerror.ErrInternal, err.Error())
		return
	}
	w.Header().Set(rest.HEADER_RESPONSE_STATUS, strconv.Itoa(http.StatusOK))
	w.Header().Set(rest.HEADER_CONTENT_TYPE, rest.CONTENT_TYPE_JSON)
	w.WriteHeader(http.StatusOK)
	_, err = fmt.Fprintln(w, util.BytesToStringWithNoCopy(objJSON))
	if err != nil {
		return
	}
}

func HttpErrResponse(w http.ResponseWriter, statusCode int, obj interface{}) {
	if obj == nil {
		w.Header().Set(rest.HEADER_RESPONSE_STATUS, strconv.Itoa(statusCode))
		w.Header().Set(rest.HEADER_CONTENT_TYPE, rest.CONTENT_TYPE_TEXT)
		w.WriteHeader(statusCode)
		return
	}

	objJSON, err := json.Marshal(obj)
	if err != nil {
		log.Errorf(err, "json marshal object fail")
		return
	}
	w.Header().Set(rest.HEADER_RESPONSE_STATUS, strconv.Itoa(http.StatusOK))
	w.Header().Set(rest.HEADER_CONTENT_TYPE, rest.CONTENT_TYPE_JSON)
	w.WriteHeader(statusCode)
	_, err = fmt.Fprintln(w, util.BytesToStringWithNoCopy(objJSON))
	if err != nil {
		log.Errorf(err, "send http response fail")
	}
}

// heartbeat use put to update a service register info
func Heartbeat(ctx context.Context, mp1SvcId string) error {
	serviceID := mp1SvcId[:len(mp1SvcId)/2]
	instanceID := mp1SvcId[len(mp1SvcId)/2:]
	req := &proto.HeartbeatRequest{
		ServiceId:  serviceID,
		InstanceId: instanceID,
	}
	_, err := core.InstanceAPI.Heartbeat(ctx, req)
	return err
}

func GetServiceInstance(ctx context.Context, serviceId string) (*proto.MicroServiceInstance, error) {
	domainProjet := util.ParseDomainProject(ctx)
	serviceID := serviceId[:len(serviceId)/2]
	instanceID := serviceId[len(serviceId)/2:]
	instance, err := svcutil.GetInstance(ctx, domainProjet, serviceID, instanceID)
	if err != nil {
		return nil, err
	}
	if instance == nil {
		err = fmt.Errorf("domainProjet %s sservice Id %s not exist", domainProjet, serviceID)
	}
	return instance, err
}

func FindInstanceByKey(result url.Values) (*proto.FindInstancesResponse, error) {
	serCategoryId := result.Get("ser_category_id")
	scopeOfLocality := result.Get("scope_of_locality")
	consumedLocalOnly := result.Get("consumed_local_only")
	isLocal := result.Get("is_local")
	isQueryAllSvc := serCategoryId == "" && scopeOfLocality == "" && consumedLocalOnly == "" && isLocal == ""

	opts := []registry.PluginOp{
		registry.OpGet(registry.WithStrKey("/cse-sr/inst/files///"), registry.WithPrefix()),
	}
	resp, err := backend.Registry().TxnWithCmp(context.Background(), opts, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("query from etch error")
	}
	var findResp []*proto.MicroServiceInstance
	for _, value := range resp.Kvs {
		var instance map[string]interface{}
		err = json.Unmarshal(value.Value, &instance)
		if err != nil {
			return nil, fmt.Errorf("string convert to instance failed")
		}
		dci := &proto.DataCenterInfo{Name: "", Region: "", AvailableZone: ""}
		instance["datacenterinfo"] = dci
		var message []byte
		message, err = json.Marshal(&instance)
		if err != nil {
			log.Errorf(err, "Instance convert to string failed!")
		}
		var ins *proto.MicroServiceInstance
		err = json.Unmarshal(message, &ins)
		if err != nil {
			log.Errorf(err, "String convert to MicroServiceInstance failed!")
		}
		property := ins.Properties
		if isQueryAllSvc && property != nil {
			findResp = append(findResp, ins)
		} else if strings.EqualFold(property["serCategory/id"], serCategoryId) ||
			strings.EqualFold(property["ConsumedLocalOnly"], consumedLocalOnly) ||
			strings.EqualFold(property["ScopeOfLocality"], scopeOfLocality) ||
			strings.EqualFold(property["IsLocal"], isLocal) {
			findResp = append(findResp, ins)
		}
	}
	if len(findResp) == 0 {
		return nil, fmt.Errorf("service not found")
	}
	response := &proto.Response{Code: 0, Message: ""}
	ret := &proto.FindInstancesResponse{Response: response, Instances: findResp}
	return ret, nil
}

func SetMapValue(theMap map[string]interface{}, key string, val interface{}) {
	mapValue, ok := theMap[key]
	if !ok || mapValue == nil {
		theMap[key] = val
	}
}

