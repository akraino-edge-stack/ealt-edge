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
	"crypto/tls"
	"crypto/x509"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// const
var cipherSuiteMap = map[string]uint16{
	"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256": tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384": tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
}

// register to mep
func RegisterToMep(param string, url string) (string, error) {
	response, errPost := DoPost(param, url)
	if errPost != nil {
		log.Println("Failed to send request")
		return "", errPost
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return "", errors.New("request failed, status is " + strconv.Itoa(response.StatusCode))
	}
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Println("Failed to read response")
		return "", err
	}

	return string(body), nil
}

func DoPost(param string, url string) (*http.Response, error) {
	sslMode := os.Getenv("APP_SSL_MODE")

	//if ssl mode is enabled, then config tls
	if sslMode == "0" {
		response, errPost := http.Post(url, "application/json", strings.NewReader(param))
		if errPost != nil {
			log.Println("Failed to create http request")
			return nil, errPost
		}
		return response, nil
	} else {
		req, errReq := http.NewRequest("POST", url, strings.NewReader(param))
		if errReq != nil {
			log.Println("Failed to create https request")
			return nil, errReq
		}
		response, errDo := DoRegister(req)
		if errDo != nil {
			log.Println("Failed to post https request %s", errDo)
			return nil, errDo
		}
		return response, nil
	}
}

func DoRegister(req *http.Request) (*http.Response, error) {
	config, err := TlsConfig()
	if err != nil {
		log.Println("Failed to config HTTPS")
		return nil, err
	}

	trans := &http.Transport{
		TLSClientConfig: config,
	}

	client := &http.Client{Transport: trans}

	return client.Do(req)
}

func TlsConfig() (*tls.Config, error) {
	caCert, err := ioutil.ReadFile(os.Getenv("SSL_ROOT"))
	if err != nil {
		log.Println("Failed to read  cert from file")
		return nil, err
	}

	CACERT := x509.NewCertPool()
	CACERT.AppendCertsFromPEM(caCert)

	appconf, err1 := GetAppConf("./conf/app_info.yaml")
	if err1 != nil {
		log.Println("Failed to read  cipher from file")
		return nil, err1
	}

	cipherslist := appconf.SslCipherSuite
	if cipherslist == "" {
		log.Println("no cipher provided in conf")
		return nil, err
	}

    ciphermap := getcipher(cipherslist)
    if ciphermap == nil {
		return nil, err
	}

	return &tls.Config{
		RootCAs: CACERT,
		ServerName:   os.Getenv("CA_CERT_DOMAIN_NAME"),
		CipherSuites: ciphermap,
		MinVersion: tls.VersionTLS12,
	}, nil
}

func getcipher(ciphers string) []uint16 {
	ciphersmap := make([]uint16, 0)
	cipherlist := strings.Split(ciphers, ",")
	for _, ciphername := range cipherlist {
		ciphernametrim := strings.TrimSpace(ciphername)
		if len(ciphernametrim) == 0 {
			continue
		}

		ciphervalue, ok := cipherSuiteMap[ciphernametrim]
		if !ok {
			log.Println("not recommended cipher")
			return nil
		}
		ciphersmap = append(ciphersmap, ciphervalue)
	}

	if len(ciphersmap) <= 0 {
		log.Println("no cipher in list")
		return nil
	}

	return ciphersmap
}
