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
 ip_var="$1"
    # checks if variable is unset
 if [ -z "$ip_var" ] ; then
    echo "ip is not set"
    return 1
 fi

 if ! echo "$ip_var" | grep -qE '^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.)' ; then
   return 1
 fi
 return 0
}

validate_name() {
  hostname="$1"
  len="${#hostname}"
  if [ "${len}" -gt "64" ]; then
    return 1
  fi
  if ! echo "$hostname" | grep -qE '^[a-zA-Z0-9]*$|^[a-zA-Z0-9][a-zA-Z0-9_\-]*[a-zA-Z0-9]$'; then
    return 1
  fi
  return 0
}

echo "Running pcb defect detection App"
umask 0027
cd /usr/app || exit
python -u run.py
