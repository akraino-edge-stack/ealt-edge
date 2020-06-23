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

FROM golang:1.13.4-alpine3.10 as builder

ENV GOPROXY https://goproxy.io
ENV GO111MODULE on
ENV HOME=/go/release

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN mkdir -p $HOME

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR $HOME
RUN mkdir -p /go/release/application/packages

ADD . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o app cmd/broker/main.go


FROM golang:1.13.4-alpine3.10

RUN mkdir -p /go/release

ENV HOME=/go/release
ENV APP_HOME=/go/release
ENV UID=166
ENV GID=166
ENV USER_NAME=ealtuser
ENV GROUP_NAME=ealtgroup

RUN apk update &&\
    apk add shadow &&\
    groupadd -r -g $GID $GROUP_NAME &&\
    useradd -r -u $UID -g $GID -d $HOME -s /sbin/nologin -c "Docker image user" $USER_NAME &&\
    chown -R $USER_NAME:$GROUP_NAME $HOME

WORKDIR $APP_HOME
RUN mkdir -p /go/release/application/packages

USER $USER_NAME

COPY --chown=ealtuser:ealtgroup --from=builder $HOME/app $APP_HOME
COPY --chown=ealtuser:ealtgroup --from=builder $HOME/start.sh $APP_HOME

RUN chmod 750 $HOME &&\
    chmod 550 $HOME/app &&\
    chmod 550 $HOME/start.sh

EXPOSE 8081

ENTRYPOINT ["./app"]
