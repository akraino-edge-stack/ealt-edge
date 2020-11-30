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

# Validates if ip is valid
validate_ip()
{


}

validate_name() {

}

# validates whether file exist
validate_file_exists() {

}

validate_ip "$LISTEN_IP"
valid_listen_ip="$?"
if [ ! "$valid_listen_ip" -eq "0" ]; then
  echo "invalid ip address for listen ip"
  exit 1
fi

if [ ! -z "$SERVER_NAME" ]; then
  validate_name "$SERVER_NAME"
  valid_name="$?"
  if [ ! "$valid_name" -eq "0" ]; then
    echo "invalid ssl server name"
    exit 1
  fi
fi

# ssl parameters validation
validate_file_exists "/usr/app/ssl/server_tls.crt"
valid_ssl_server_cert="$?"
if [ ! "$valid_ssl_server_cert" -eq "0" ]; then
  echo "invalid ssl server certificate"
  exit 1
fi

# ssl parameters validation
validate_file_exists "/usr/app/ssl/server_tls.key"
valid_ssl_server_key="$?"
if [ ! "$valid_ssl_server_key" -eq "0" ]; then
  echo "invalid ssl server key"
  exit 1
fi

# ssl parameters validation
validate_file_exists "/usr/app/ssl/ca.crt"
valid_ssl_ca_crt="$?"
if [ ! "$valid_ssl_ca_crt" -eq "0" ]; then
  echo "invalid ssl ca cert"
  exit 1
fi

echo "Running Monitoring Service"
umask 0027
cd /usr/app || exit
python run.py
