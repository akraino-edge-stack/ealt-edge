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

import io
import sys
import os
import numpy as np
import cv2 
import logging
import config
from flask_sslify import SSLify
from flask import Flask, request, jsonify, Response, make_response
from flask_cors import CORS
from werkzeug import secure_filename


class model_info():
    def __init__ (self, model_name):
        self.model = 'model_info/MobileNetSSD_deploy.caffemodel'
        self.model_name = model_name
        self.prototxt = 'model_info/MobileNetSSD_deploy.prototxt'
        self.confidenceLevel = 80

    def get_model(self):
        return self.model

    def get_prototxt(self):
        return self.prototxt

    def get_model_name(self):
        return self.model_name

    def set_confidence_level(self, confidenceLevel):
        self.confidenceLevel = confidenceLevel
    
    def get_confidence_level(self):
        return self.confidenceLevel

    def update_model(self, model, prototxt, model_name):
        self.prototxt = prototxt
        self.model = model
        self.model_name = model_name

# Labels of Network.
classNames = { 0: 'background',
    1: 'aeroplane', 2: 'bicycle', 3: 'bird', 4: 'boat',
    5: 'bottle', 6: 'bus', 7: 'car', 8: 'cat', 9: 'chair',
    10: 'cow', 11: 'diningtable', 12: 'dog', 13: 'horse',
    14: 'motorbike', 15: 'person', 16: 'pottedplant',
    17: 'sheep', 18: 'sofa', 19: 'train', 20: 'tvmonitor' }

app = Flask(__name__)
CORS(app)
sslify = SSLify(app)
app.config['JSON_AS_ASCII'] = False
#app.config['UPLOAD_PATH'] = '/usr/app/images/'
app.config['UPLOAD_PATH'] = '/home/root1/My_Work/Akraino/MEC_BP/Rel4/Retail-apps/aPaaS/src/Obj_Detection_service/images'
#app.config['UPLOAD_FOLDER']
app.config['supports_credentials'] = True
app.config['CORS_SUPPORTS_CREDENTIALS'] = True
app.config['MAX_CONTENT_LENGTH'] = 16 * 1024 * 1024
ALLOWED_EXTENSIONS = set(['png', 'jpg', 'jpeg'])
count = 0
listOfMsgs = []

def allowed_file(filename):
	return '.' in filename and filename.rsplit('.', 1)[1].lower() in ALLOWED_EXTENSIONS



@app.route('/mep/v1/obj_detection/uploadModel', methods=['POST'])
def uploadModel():
    """
    upload model
    :return: html file
    """

    return Response("success")


@app.route('/mep/v1/obj_detection/confidencelevel', methods=['POST'])
def setConfidenceLevel():
    """
    Trigger the video_feed() function on opening "0.0.0.0:5000/video_feed" URL
    :return:
    """

    return Response("success")

@app.route('/mep/v1/obj_detection/detect', methods=['GET'])
def Obj_Detection():
    """
    Trigger the Obj detection on input frame/image
    Input: frame/image
    :return: imposed frame, count, Obj type, time taken by detection
    """
    return Response("success")

def start_server(handler):
    app.logger.addHandler(handler)
    if config.ssl_enabled:
        context = (config.ssl_certfilepath, config.ssl_keyfilepath)
        app.run(host=config.server_address, debug=True, ssl_context=context, threaded=True, port=config.server_port)
    else:
        app.run(host=config.server_address, debug=True, threaded=True, port=config.server_port)
