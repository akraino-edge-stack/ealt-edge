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

CURRENT_DIR=`pwd`

echo "DOCKER_BUILD_DIR=${CURRENT_DIR}"
echo "Build and Push APP LCM"

# Build and push broker
cd ${CURRENT_DIR}
cd ../../broker/
. build_push_image.sh

# Build and push k8s helm plugin
cd ${CURRENT_DIR}
cd ../../k8shelm/
. build_push_image.sh

