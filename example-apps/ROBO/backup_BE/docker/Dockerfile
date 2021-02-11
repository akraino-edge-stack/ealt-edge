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

# FROM swr.ap-southeast-1.myhuaweicloud.com/edgegallery/openjdk:8u201-jdk-alpine

# currently this pod is running as root user due to this need to communicate with host shell.config
# in future this can be further checked with adding some level restrictions to user
FROM openjdk:11-jre-slim

ENV JAVA_HOME /usr/lib/jvm/java-11-openjdk-amd64/bin/java

ENV APP_FILE backup.jar
ENV APP_HOME /usr/app

# Create the home directory for the new app user.
RUN mkdir -p /usr/app
RUN mkdir -p /usr/app/bin

# Create an app user so our program doesn't run as root.
RUN apt-get update && apt-get install -y sshpass

WORKDIR $APP_HOME

RUN chmod 750 $APP_HOME &&\
    chmod -R 550 $APP_HOME/bin &&\
    mkdir -p -m 750 $APP_HOME/config &&\
    mkdir -p -m 750 $APP_HOME/log &&\
    mkdir -p -m 750 $APP_HOME/packages &&\
    mkdir -p -m 700 $APP_HOME/ssl &&\
    mkdir -p -m 700 $APP_HOME/.kube

COPY /target/*.jar $APP_HOME/bin

EXPOSE 8091

# Change to the app user.

# Execute script & application
ENTRYPOINT ["sh", "-c"]
CMD ["exec java -jar ./bin/backup.jar"]