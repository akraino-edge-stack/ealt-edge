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

# Prepare stage for multistage image build
## START OF STAGE0 ##
# FROM python:3.6-slim-stretch
# FROM python:3.9
# FROM python:3.9-slim-buster
# # CREATE APP USER ##

FROM python:2.7

# Set umask
RUN sed -i "s|umask 022|umask 027|g" /etc/profile

# Create the home directory for the new app user.
RUN mkdir -p /usr/app
RUN mkdir -p /usr/app/bin
RUN mkdir -p /usr/app/data
RUN mkdir -p /usr/app/help_utils
RUN mkdir -p /usr/app/libs
RUN mkdir -p /usr/app/output
RUN mkdir -p /usr/app/results
RUN mkdir -p /usr/app/tools

RUN mkdir -p /usr/app/input_image
RUN mkdir -p /usr/app/output_image
RUN mkdir -p /usr/app/input_image/beijinglab
RUN mkdir -p /usr/app/input_image/beijinglab/camera1
RUN mkdir -p /usr/app/input_image/beijinglab/camera2
RUN mkdir -p /usr/app/input_image/beijinglab/camera3
RUN mkdir -p /usr/app/input_image/shenzhenlab
RUN mkdir -p /usr/app/input_image/shenzhenlab/camera1
RUN mkdir -p /usr/app/input_image/shenzhenlab/camera2
RUN mkdir -p /usr/app/input_image/shenzhenlab/camera3
RUN mkdir -p /usr/app/input_image/shanghailab
RUN mkdir -p /usr/app/input_image/shanghailab/camera1
RUN mkdir -p /usr/app/input_image/shanghailab/camera2
RUN mkdir -p /usr/app/input_image/shanghailab/camera3
#output image path
RUN mkdir -p /usr/app/output_image/beijinglab
RUN mkdir -p /usr/app/output_image/beijinglab/camera1
RUN mkdir -p /usr/app/output_image/beijinglab/camera2
RUN mkdir -p /usr/app/output_image/beijinglab/camera3
RUN mkdir -p /usr/app/output_image/shenzhenlab
RUN mkdir -p /usr/app/output_image/shenzhenlab/camera1
RUN mkdir -p /usr/app/output_image/shenzhenlab/camera2
RUN mkdir -p /usr/app/output_image/shenzhenlab/camera3
RUN mkdir -p /usr/app/output_image/shanghailab
RUN mkdir -p /usr/app/output_image/shanghailab/camera1
RUN mkdir -p /usr/app/output_image/shanghailab/camera2
RUN mkdir -p /usr/app/output_image/shanghailab/camera3

# Set the home directory to our app user's home.
ENV APP_HOME=/usr/app
ENV UID=166
ENV GID=166
ENV USER_NAME=eguser
ENV GROUP_NAME=eggroup
ENV ENV="/etc/profile"
ENV PYTHONUNBUFFERED=0

# Create an app user so our program doesn't run as root.
RUN apt-get -y update &&\
    groupadd -r -g $GID $GROUP_NAME &&\
    useradd -r -u $UID -g $GID -d $APP_HOME -s /sbin/nologin -c "Docker image user" $USER_NAME

RUN apt-get install -y --fix-missing \
    build-essential \
    cmake \
    gfortran \
    git \
    wget \
    curl \
    graphicsmagick \
    libgraphicsmagick1-dev \
    libavcodec-dev \
    libavformat-dev \
    libgtk2.0-dev \
    libjpeg-dev \
    liblapack-dev \
    libswscale-dev \
    pkg-config \
    software-properties-common \
    zip \
    imagemagick \
    && apt-get clean && rm -rf /tmp/* /var/tmp/*

# Set the working directory.
WORKDIR $APP_HOME

# Copy the application & scripts
COPY requirements.txt $APP_HOME/
# Install requirements
RUN pip install -r requirements.txt

COPY data $APP_HOME/data/
COPY help_utils $APP_HOME/help_utils/
COPY libs $APP_HOME/libs/
COPY output $APP_HOME/output/
COPY tools $APP_HOME/tools/
COPY input_image/1 $APP_HOME/input_image/beijinglab/camera1
COPY input_image/2 $APP_HOME/input_image/shenzhenlab/camera1
COPY input_image/3 $APP_HOME/input_image/shanghailab/camera1
COPY input_image/4 $APP_HOME/input_image/beijinglab/camera2
COPY input_image/5 $APP_HOME/input_image/beijinglab/camera3
COPY input_image/6 $APP_HOME/input_image/shenzhenlab/camera2
COPY input_image/7 $APP_HOME/input_image/shenzhenlab/camera3
COPY input_image/8 $APP_HOME/input_image/shanghailab/camera2
COPY input_image/9 $APP_HOME/input_image/shanghailab/camera3

COPY config.py run.py $APP_HOME/

COPY configs/start.sh $APP_HOME/bin

RUN chmod 750 $APP_HOME &&\
    chmod -R 550 $APP_HOME/bin &&\
    mkdir -p -m 750 $APP_HOME/log &&\
    mkdir -p -m 700 $APP_HOME/ssl &&\
    chown -R $USER_NAME:$GROUP_NAME $APP_HOME

# Exposed port
EXPOSE 9991

# Change to the app user.
# USER $USER_NAME

# Install requirements
#RUN pip install -r requirements.txt

# Execute script & application
ENTRYPOINT ["sh", "./bin/start.sh"]
