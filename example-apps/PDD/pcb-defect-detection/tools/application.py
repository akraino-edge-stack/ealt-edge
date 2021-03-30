
import config
from flask_sslify import SSLify
from flask import Flask, request, jsonify, Response
from flask_cors import CORS

import json
import requests
import os
import os.path
from os import path
import base64
import sys

app = Flask(__name__)
CORS(app)
sslify = SSLify(app)
app.config['JSON_AS_ASCII'] = False
app.config['INPUT_IMAGE_PATH'] = '/usr/app/input_image/'
app.config['OUTPUT_IMAGE_PATH'] = '/usr/app/output_image/'

app.config['MAX_CONTENT_LENGTH'] = 16 * 1024 * 1024
ALLOWED_EXTENSIONS = set(['png', 'jpg', 'jpeg'])
count = 0
listOfPath = []
listOfImages = []
mock_func = 0
# result_image_path = app.config['OUTPUT_IMAGE_PATH']

listOfPath = ["beijinglab/camera1", "beijinglab/camera2",
              "beijinglab/camera3",
              "shenzhenlab/camera1", "shenzhenlab/camera2",
              "shenzhenlab/camera3",
              "shanghailab/camera1", "shanghailab/camera2",
              "shanghailab/camera3"]

def path_exist(image_path):
    print('image path:',image_path)

    pathexist = 0
    for imagepath in listOfPath:
        print('image path in list:', imagepath)
        if (image_path == imagepath):
            print('image path match')
            pathexist = 1
            break

    return pathexist

def get_images(abs_image_path):
    """
    preview images

    :return: result image
    """
    del listOfImages[:]

    arr = os.listdir(abs_image_path)
    for file in arr:
        print('image name:',abs_image_path + file)

    for image in arr:
        ## as base64 string
        with open(abs_image_path + image, "rb") as img_file:
            jpeg_bin = base64.b64encode(img_file.read())
            listOfImages.append(jpeg_bin)

    response = {'image01': listOfImages[0],
                'image02': listOfImages[1],
                'image03': listOfImages[2],}
                #'image04': listOfImages[3]}
               # 'image05': listOfImages[4],}

    return response

def detection(input_image_path, output_image_path):
    """
    detection

    :return: detection status
    """
    cmd = 'cd tools' + ' && python inference.py --data_dir=' + \
          input_image_path + ' --save_dir=' + output_image_path + ' --GPU=0'
    print(cmd)

    os.system(cmd)
    return 1


@app.route('/v1/pcb/preview/<EdgeLoc>/<cameraNum>', methods=['GET'])
def preview_image(EdgeLoc, cameraNum):
    """
    preview images

    :return: result image
    """
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")

    image_path = EdgeLoc + '/' + cameraNum
    pathexist = path_exist(image_path)
    if pathexist == 0:
        response = {'image' : 'null'}
        print ('image path not exist:',image_path)
        return jsonify(response)

    abs_image_path = app.config['INPUT_IMAGE_PATH'] + image_path + '/'
    print('abs image path:',abs_image_path)
    response = get_images(abs_image_path)

    return jsonify(response)


@app.route('/v1/pcb/resultimage', methods=['GET'])
def result_image():
    """
    preview images

    :return: result image
    """
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")

    print('result image path global:',result_image_path)
    response = get_images(result_image_path)

    return jsonify(response)


@app.route('/v1/pcb/detection/<EdgeLoc>/<cameraNum>', methods=['GET'])
def detect_image(EdgeLoc, cameraNum):
    """
    detect images

    :return: success or failure
    """
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    global result_image_path

    image_path = EdgeLoc + '/' + cameraNum
    pathexist = path_exist(image_path)
    if pathexist == 0:
        response = {'responce': 'failure'}
        print ('image path not exist:',image_path)
        return jsonify(response)

    input_image_path = app.config['INPUT_IMAGE_PATH'] + image_path + '/'
    print('input image path:', input_image_path)

    output_image_path = app.config['OUTPUT_IMAGE_PATH'] + image_path + '/'
    print('out image path:', output_image_path)

    ret = detection(input_image_path, output_image_path)
    if (ret == 0):
        response = {'responce': 'failure'}
        print('detectio algo failed ')
        return jsonify(response)

    result_image_path = output_image_path
    print('result image path:', result_image_path)

    response = {'responce': 'success'}
    print('detection sucess')
    return jsonify(response)


def start_server(handler):
    app.logger.addHandler(handler)
    if config.ssl_enabled:
        context = (config.ssl_certfilepath, config.ssl_keyfilepath)
        app.run(host=config.server_address, debug=True, ssl_context=context,
                threaded=True, port=config.server_port)
    else:
        app.run(host=config.server_address, debug=True, threaded=True,
                port=config.server_port)
