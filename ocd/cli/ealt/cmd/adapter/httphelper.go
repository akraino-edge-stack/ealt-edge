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
	"bytes"
	"ealt/cmd/common"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
)

var MECMClusterIP = os.Getenv("MECMClusterIP")
var APPLCMPort = os.Getenv("MECMClusterPort")
var ONBOARDPACKAGEPATH = os.Getenv("ONBOARDPACKAGEPATH")
var client = http.Client{}

func httpEndPointBuider(uri string) string {

	return "http://" + strings.TrimSpace(MECMClusterIP) + strings.TrimSpace(APPLCMPort) + uri

}

func HttpDeleteRequestBuilder(uri string, body []byte) {

	uri = httpEndPointBuider(uri)
	request, err := http.NewRequest(http.MethodDelete, uri, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	output, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(output))

}

func HttpPostRequestBuilder(uri string, body []byte) error {

	uri = httpEndPointBuider(uri)
	request, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")

	fmt.Println(request)

	if err != nil {
		log.Fatalln(err)
	}

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	output, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(string(output))
	return nil
}

func HttpMultiPartPostRequestBuilder(uri string, body []byte, file string) error {

	filepath := getFilePathWithName(file)
	uri = httpEndPointBuider(uri)

	request, err := fileUploadRequest(uri, "file", filepath)

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}

	var result map[string]interface{}
	json.NewDecoder(response.Body).Decode(&result)
	fmt.Println(result)

	return nil
}

func getFilePathWithName(file string) string {

	return ONBOARDPACKAGEPATH + common.PATHSLASH + file
}

func fileUploadRequest(uri string, paramName, filepath string) (*http.Request, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	//close the file later
	defer file.Close()

	//Buffer to store the request body as bytes
	var requestBody bytes.Buffer
	multiPartWriter := multipart.NewWriter(&requestBody)

	fileWriter, err := multiPartWriter.CreateFormFile("file_field", filepath)
	if err != nil {
		log.Fatalln(err)
	}

	//Copy the actual file contents
	_, err = io.Copy(fileWriter, file)
	if err != nil {
		log.Fatalln(err)
	}

	//Close multiwriter
	multiPartWriter.Close()

	request, err := http.NewRequest(http.MethodPost, uri, &requestBody)
	if err != nil {
		log.Fatalln(err)
	}
	request.Header.Set("Content-Type", "multipart/form-data")
	return request, err
}
