#!/bin/bash
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

# Command to update Kong


curl --location --request DELETE 'http://159.138.22.15:30012/routes/mp1'
curl --location --request DELETE 'http://159.138.22.15:30012/services/http-mp1'
curl --location --request POST 'http://159.138.22.15:30012/services' --header 'Content-Type: application/json' --data '{"url": "https://mep-service:8088","name": "http-mp1"}'
curl --location --request POST 'http://159.138.22.15:30012/services/http-mp1/routes' --header 'Content-Type: application/json' --data '{"paths": ["/mp1"], "name": "mp1"}'
