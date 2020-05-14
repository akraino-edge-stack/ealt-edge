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

FROM golang:1.14.2-alpine3.11

ENV GOPROXY https://goproxy.io
ENV GO111MODULE on

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/release

ADD . .
RUN apk update && apk add --no-cache musl-dev gcc build-base

RUN GOOS=linux CGO_ENABLED=0 go build -buildmode=pie -ldflags="-s -w" -installsuffix cgo -o app main.go

RUN cp -r conf/ / && cp app start.sh /

EXPOSE 8088

CMD ["/start.sh"]
