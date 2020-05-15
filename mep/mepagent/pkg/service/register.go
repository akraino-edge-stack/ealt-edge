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
	"encoding/json"
	"log"
	"strconv"
	"time"
)

func SvcReg(confPath string) (string, error) {
	conf, err := GetConf(confPath)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}

	appInstanceId := conf.AppInstanceId
	serviceInfos := conf.ServiceInfoPosts
	mepServerIP := conf.MepServerIP
	mepServerPORT := conf.MepServerPORT
	url := "http://" + mepServerIP + ":" + mepServerPORT + "/mep/mec_service_mgmt/v1/applications/" + appInstanceId + "/services"

    for _, serviceInfo := range serviceInfos {
		data, e := json.Marshal(serviceInfo)
		if e != nil {
			log.Println("Failed to marshal service info to object")
			continue
		}

		for i := 1; i <= 5; i++ { // if register failed, then retry five times
			_, err := RegisterToMep(string(data), url)
			if err != nil {
				log.Println("Failed to register to mep, appInstanceId is" + appInstanceId + ", serviceName is " + serviceInfo.SerName)
			} else {
				log.Println("Register mep main to mep success, appInstanceId is" + appInstanceId + ", serviceName is " + serviceInfo.SerName)
				break
			}
			log.Println("Failed to register mep main to mep, will retry 5 times, retry interval is 30 s, already retry " + strconv.Itoa(i) + " time")
			time.Sleep(30 * time.Second) // register failed , 30 seconds after try again
		}
	}

	return "", nil
}
