#
# Copyright 2020 Huawei Technologies Co., Ltd.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

import os

# [Server Configurations]
server_port = 9995
# server_address = os.environ.get('LISTEN_IP')
server_address = "0.0.0.0"

# [InfluxDB config]
IPADDRESS = os.environ.get('INFLUXDB_IP')
PORT = os.environ.get('INFLUXDB_PORT')
DATABASE_NAME = "Shelf_Inventry"

# [SSL Configurations]
ssl_enabled = False
ssl_protocol = "TLSv1.2"
ssl_ciphers = ["TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
               "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
               "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
               "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384"]
ssl_certfilepath = "/usr/app/ssl/server_tls.crt"
ssl_keyfilepath = "/usr/app/ssl/server_tls.key"
ssl_cacertpath = "/usr/app/ssl/ca.crt"
ssl_server_name = os.environ.get('SERVER_NAME', "ealtedge")

# [Service Configurations]
api_gateway = os.environ.get("API_GATEWAY", "apigw.mep.org")
# Obj_Det_SER_IP = os.environ.get("OBJ_DET_SER_IP", "localhost")

Obj_Det_SER_IP = os.environ.get("OBJ_DET_SER_IP")
# Obj_Det_SER_IP = os.environ.get("OBJ_DET_SER_IP")
Obj_Det_SER_PORT = os.environ.get("OBJ_DET_SER_PORT")
# Obj_Det_SER_PORT = os.environ.get("OBJ_DET_SER_PORT")
Obj_Det = os.environ.get("OBJ_DETECTION", "mep/v1/obj_detection")

# [Constants]
# detection_url = "http://" + api_gateway + "/" + Obj_Det
detection_url = "http://" + Obj_Det_SER_IP + ":" + Obj_Det_SER_PORT + "/" + \
                                           Obj_Det + "/"
