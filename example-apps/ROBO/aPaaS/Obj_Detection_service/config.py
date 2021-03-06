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
server_port = 9998
server_address = os.environ.get('LISTEN_IP')

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
obj_det = os.environ.get("OBJ_DETECTION", "objdetection")

# [Constants]
recognition_url = "http://" + api_gateway + "/" + obj_det
