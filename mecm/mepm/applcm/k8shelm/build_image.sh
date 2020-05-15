#!/bin/sh
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

DOCKER_BUILD_DIR=`pwd`
MECM_VERSION=latest
IMAGE_NAME=helmplugin
REPO_NAME=ealtedge

echo "DOCKER_BUILD_DIR=${DOCKER_BUILD_DIR}"
echo "In Build and Push K8s Help Plugin"

function build_image {
    docker build --no-cache -t ${REPO_NAME}/${IMAGE_NAME}:${MECM_VERSION} -f build/Dockerfile .
}

build_image
