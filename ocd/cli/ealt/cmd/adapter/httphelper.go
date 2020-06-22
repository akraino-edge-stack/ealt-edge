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
var sslmode = os.Getenv("EALTSSLMode")

func httpEndPointBuider(uri string) string {
	localURI := strings.TrimSpace(MECMClusterIP) + ":" + strings.TrimSpace(APPLCMPort) + uri
	if sslmode == "1" {
		return "https://" + localURI
	}
	return "http://" + localURI
}

//Function to build the Get Requests for Application Package
//Management and Application Life Cycle Management.
func HttpGetRequestBuilder(uri string, body []byte) {

	uri = httpEndPointBuider(uri)
	fmt.Println("Request URL :\t" + uri)
	request, err := http.NewRequest(http.MethodGet, uri, bytes.NewBuffer(body))
	request.Header.Set(common.ContentType, common.ApplicationJson)
	if err != nil {
		log.Fatalln(err)
	}
	client := GetHttpClient()
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	output, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Response Data: \n" + string(output))
}

//HTTP DELETE Message Builder
func HttpDeleteRequestBuilder(uri string, body []byte) {

	uri = httpEndPointBuider(uri)
	fmt.Println("Request URL :\t" + uri)
	request, err := http.NewRequest(http.MethodDelete, uri, bytes.NewBuffer(body))
	request.Header.Set(common.ContentType, common.ApplicationJson)

	if err != nil {
		log.Fatalln(err)
	}
	client := GetHttpClient()
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()
	output, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Response Data: \n" + string(output))
}

func HttpPostRequestBuilder(uri string, body []byte) error {

	uri = httpEndPointBuider(uri)
	fmt.Println("Request URL :\t" + uri)
	fmt.Println("Request Body :\t" + string(body) + "\n")
	request, err := http.NewRequest(http.MethodPost, uri, bytes.NewBuffer(body))
	request.Header.Set(common.ContentType, common.ApplicationJson)

	if err != nil {
		log.Fatalln(err)
	}
	client := GetHttpClient()
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	}
	defer response.Body.Close()

	output, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Response Data: \n\n" + string(output))
	return nil
}

func HttpMultiPartPostRequestBuilder(uri string, body []byte, file string) error {

	filepath := getFilePathWithName(file)
	fmt.Println("File Path :" + filepath)
	uri = httpEndPointBuider(uri)
	fmt.Println("Request URL :\t" + uri)
	request, err := fileUploadRequest(uri, "file", filepath, file)
	if err != nil {
		log.Fatalln(err)
	}
	client := GetHttpClient()
	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
	} else {
		body := &bytes.Buffer{}
		_, err := body.ReadFrom(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		response.Body.Close()

		fmt.Println("Response Body:")

		fmt.Println(body)
		var result map[string]interface{}
		json.NewDecoder(response.Body).Decode(&result)

		fmt.Println("ID has to be send in Create Application Instance Request")
	}
	return nil
}

func getFilePathWithName(file string) string {

	return ONBOARDPACKAGEPATH + common.PATHSLASH + file
}

func fileUploadRequest(uri string, paramName, filepath, filename string) (*http.Request, error) {

	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	//close the file later
	defer file.Close()

	//Buffer to store the request body as bytes
	//var requestBody bytes.Buffer
	requestBody := &bytes.Buffer{}
	multiPartWriter := multipart.NewWriter(requestBody)

	fileWriter, err := multiPartWriter.CreateFormFile(paramName, filename)
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
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest(http.MethodPost, uri, requestBody)
	request.Header.Set(common.ContentType, multiPartWriter.FormDataContentType())
	//request.Header.Set("Content-Type", "multipart/form-data")

	if err != nil {
		log.Fatalln(err)
	}

	return request, err
}
