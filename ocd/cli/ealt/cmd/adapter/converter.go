/*
Copyright 2020 Huawei Technologies Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package adapter

import (
	"ealt/cmd/common"
	model "ealt/cmd/model"
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func BuilderRequest(valueArgs []string, command string) error {

	var URIString string

	switch command {
	case "NewAppCreateCommand":
		//Onboard Command
		//ealtedge/mepm/app_pkgm/v1/app_packages/
		//read the file from the system.
		URIString = common.AppmUriCreate
		var packageName string
		var body []byte
		body = jsonEmptyBodyFormat()
		packageName = strings.TrimSpace(valueArgs[0])
		HttpMultiPartPostRequestBuilder(URIString, body, packageName)

	case "NewAppInfoCommand":
		URIString = common.AppmUri
		var body []byte
		URIString = common.AppmUri + strings.TrimSpace(valueArgs[0])
		body = jsonEmptyBodyFormat()
		HttpGetRequestBuilder(URIString, body)

	case "NewAppDeleteCommand":
		//The Delete Application Package URI
		//ealtedge/mepm/app_pkgm/v1/app_packages/{{ID}}
		var body []byte
		URIString = common.AppmUri + strings.TrimSpace(valueArgs[0])
		body = jsonEmptyBodyFormat()
		HttpDeleteRequestBuilder(URIString, body)

	case "NewApplcmCreateCommand":
		//appLCM application Creation URI
		//ealtedge/mepm/app_lcm/v1/app_instances
		var body []byte

		URIString = common.ApplcmUriCreate
		//Assigning the AppLcm Create Command Line Flags to the Json Paylod.
		payload := model.CreateApplicationReq{AppDID: strings.TrimSpace(valueArgs[0]),
			AppInstancename:       strings.TrimSpace(valueArgs[1]),
			AppInstanceDescriptor: strings.TrimSpace(valueArgs[2])}
		body, err := json.Marshal(payload)

		if err != nil {
			log.Fatalln(err)
		}
		HttpPostRequestBuilder(URIString, body)

	case "NewApplcmInfoCommand":
		//appLCM Get Application URI
		///ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}
		var body []byte
		URIString = common.ApplcmUri + strings.TrimSpace(valueArgs[0])

		//Empty body for Delete Command.
		body = jsonEmptyBodyFormat()
		HttpGetRequestBuilder(URIString, body)

	case "NewApplcmDeleteCommand":
		//appLCM Delete Application URI
		///ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}
		var body []byte
		URIString = common.ApplcmUri + strings.TrimSpace(valueArgs[0])

		//Empty body for Delete Command.
		body = jsonEmptyBodyFormat()
		HttpDeleteRequestBuilder(URIString, body)

	case "NewApplcmStartCommand":
		//applcm application instantiate uri
		//ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}/instantiate
		var body []byte

		URIString = common.ApplcmUri + strings.TrimSpace(valueArgs[0]) + common.InstantiateUri

		selectedMECHostInfo := model.SelectedMECHostInfo{HostName: strings.TrimSpace(valueArgs[1]),
			HostId: strings.TrimSpace(valueArgs[2])}
		//Payload
		payload := model.InstantiateApplicationReq{SelectedMECHostInfo: selectedMECHostInfo}
		body, err := json.Marshal(payload)
		if err != nil {
			fmt.Println(err)
		}
		HttpPostRequestBuilder(URIString, body)

	case "NewApplcmTerminateCommand":
		//applcm application terminate uri
		//ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}/terminate
		var body []byte
		URIString = common.ApplcmUri + strings.TrimSpace(valueArgs[0]) + common.TerminateUri
		body = jsonEmptyBodyFormat()
		HttpPostRequestBuilder(URIString, body)
	}
	return nil
}

func jsonEmptyBodyFormat() []byte {
	var jsonstr []byte
	jsonstr = []byte(`{"":""}`)
	return jsonstr
}
