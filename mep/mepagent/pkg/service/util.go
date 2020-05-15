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

package service

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/akraino-edge-stack/ealt-edge/mep/mepagent/pkg/model"
)

// get yaml and parse to struct
func GetConf(path string) (model.AppInstanceInfo, error) {
	yamlFile, err := ioutil.ReadFile(path)
	var info model.AppInstanceInfo
	if err != nil {
		return info, err
	}

	err = yaml.UnmarshalStrict(yamlFile, &info)

	if err != nil {
		return info, err
	}

	return info, nil
}

// register to mep
func RegisterToMep(param string, url string) (string, error) {
	response, err := http.Post(url, "application/json", strings.NewReader(param))
	if err != nil {
		return "", err
	}

	if response.StatusCode != http.StatusCreated {
		return "", errors.New("created failed")
	}
	defer response.Body.Close()
	body, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		return "", err2
	}

	return string(body), nil
}
