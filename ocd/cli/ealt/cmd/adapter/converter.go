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
		URIString = common.AppmUri
		var packageName string
		var body []byte
		body = jsonEmptyBodyFormat()
		packageName = strings.TrimSpace(valueArgs[0])
		HttpMultiPartPostRequestBuilder(URIString, body, packageName)
		fmt.Println(packageName)

	case "NewAppDeleteCommand":
		//The Delete Application Package URI
		//ealtedge/mepm/app_pkgm/v1/app_packages/{{ID}}
		var body []byte
		URIString = common.AppmUri + strings.TrimSpace(valueArgs[0])
		body = jsonEmptyBodyFormat()
		fmt.Println(URIString)
		HttpDeleteRequestBuilder(URIString, body)

		fmt.Println(URIString)

	case "NewApplcmCreateCommand":
		//appLCM application Creation URI
		//ealtedge/mepm/app_lcm/v1/app_instances
		var body []byte

		URIString = common.ApplcmUri
		body, err := json.Marshal(map[string]string{
			"appDId":                strings.TrimSpace(valueArgs[0]),
			"appInstancename":       strings.TrimSpace(valueArgs[1]),
			"appInstanceDescriptor": strings.TrimSpace(valueArgs[2]),
		})

		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(URIString)
		HttpPostRequestBuilder(URIString, body)

	case "NewApplcmDeleteCommand":
		//appLCM Delete Application URI
		///ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}
		var body []byte
		URIString = common.ApplcmUri + strings.TrimSpace(valueArgs[0])
		body = jsonEmptyBodyFormat()
		fmt.Println(URIString)
		HttpDeleteRequestBuilder(URIString, body)

	case "NewApplcmStartCommand":
		//applcm application instantiate uri
		//ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}/instantiate
		var body []byte

		URIString = common.ApplcmUri + strings.TrimSpace(valueArgs[0]) + common.InstantiateUri
		body = jsonEmptyBodyFormat()
		fmt.Println(URIString)
		HttpPostRequestBuilder(URIString, body)

	case "NewApplcmTerminateCommand":
		//applcm application terminate uri
		//ealtedge/mepm/app_lcm/v1/app_instances/{appInstanceId}/terminate
		var body []byte
		URIString = common.ApplcmUri + strings.TrimSpace(valueArgs[0]) + common.TerminateUri
		body = jsonEmptyBodyFormat()
		fmt.Println(URIString)
		HttpPostRequestBuilder(URIString, body)

	}

	return nil
}

func jsonEmptyBodyFormat() []byte {
	var jsonstr []byte
	jsonstr = []byte(`{"":""}`)
	return jsonstr
}
