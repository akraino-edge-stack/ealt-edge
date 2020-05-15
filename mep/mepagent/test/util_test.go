/*
 *  Copyright 2020 Huawei Technologies Co., Ltd.
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package test

import (
	"encoding/json"
	"pkg/service"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetConfSuccess(t *testing.T) {
	_, err := service.GetConf("../../conf/app_instance_info.yaml")
	if err != nil {
		t.Error("Read conf file failed")
	}
}

func TestGetConfFail(t *testing.T) {
	_, err := service.GetConf("../conf/app_instance_info.yaml")
	if err == nil {
		t.Error("Read conf file failed")
	}
}

func TestRegisterToMepSuccess(t *testing.T) {
	httpResponse := "response body"
	var httpResponseBytes, err1 = json.Marshal(httpResponse)
	if err1 != nil {
		t.Error("Marshal http Response Error")
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		_, err2 := w.Write(httpResponseBytes)

		if err2 != nil {
			t.Error("Write Response Error")
		}
		if r.Method != "POST" {
			t.Error("UnExcepted Method")
		}
	}))

	defer ts.Close()
	api := ts.URL

	_, err := service.RegisterToMep("param", api)
	if err != nil {
		t.Error("error")
	}
}

func TestRegisterToMepFail1(t *testing.T) {
	httpResponse := "response body"
	var httpResponseBytes, err1 = json.Marshal(httpResponse)
	if err1 != nil {
		t.Error("Marshal http Response Error")
	}

	ts2 := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		_, err2 := w.Write(httpResponseBytes)

		if err2 != nil {
			t.Error("Write Response Error")
		}
		if r.Method != "POST" {
			t.Error("UnExcepted Method")
		}
	}))

	defer ts2.Close()
	api := ts2.URL

	_, err := service.RegisterToMep("param", api)
	if err == nil {
		t.Error("error")
	}
}

