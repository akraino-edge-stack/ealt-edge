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
import os
import cv2
import config
from flask_sslify import SSLify
from flask import Flask, request, jsonify, Response, send_file
from flask_cors import CORS
from werkzeug import secure_filename

app = Flask(__name__)
CORS(app)
sslify = SSLify(app)
app.config['JSON_AS_ASCII'] = False
app.config['UPLOAD_PATH'] = '/usr/app/images/input/'
app.config['supports_credentials'] = True
app.config['CORS_SUPPORTS_CREDENTIALS'] = True
app.config['MAX_CONTENT_LENGTH'] = 16 * 1024 * 1024
ALLOWED_EXTENSIONS = set(['png', 'jpg', 'jpeg'])
MODEL_PATH = '/usr/app/model/'
IMAGE_PATH = '/usr/app/images/result/'
count = 0
listOfMsgs = []


class model_info():
    def __init__(self, model_name):
        self.model_name = 'MobileNetSSD_deploy.caffemodel'
        self.prototxt = 'MobileNetSSD_deploy.prototxt'
        self.confidenceLevel = 0.2

    def get_prototxt(self):
        return self.prototxt

    def get_model_name(self):
        return self.model_name

    def set_confidence_level(self, confidenceLevel):
        self.confidenceLevel = confidenceLevel

    def get_confidence_level(self):
        return self.confidenceLevel

    def update_model(self, model_loc, prototxt, model_name):
        self.prototxt = prototxt
        self.model_name = model_name


# Labels of Network.
classNames = {0: 'background',
              1: 'aeroplane', 2: 'bicycle', 3: 'bird', 4: 'boat',
              5: 'bottle', 6: 'bus', 7: 'car', 8: 'cat', 9: 'chair',
              10: 'cow', 11: 'diningtable', 12: 'dog', 13: 'horse',
              14: 'motorbike', 15: 'person', 16: 'pottedplant',
              17: 'sheep', 18: 'sofa', 19: 'train', 20: 'tvmonitor'}


def allowed_file(filename):
    return '.' in filename and filename.rsplit('.', 1)[1].lower() \
       in ALLOWED_EXTENSIONS


# Obj-detection from input frame
def Detection(img):
    print ('inside detection func')
    modelInfo = model_info("caffe")
    ConfPercent = modelInfo.get_confidence_level()
    model_name = modelInfo.get_model_name()
    prototxt_name = modelInfo.get_prototxt()

    model = MODEL_PATH + model_name
    prototxt = MODEL_PATH + prototxt_name
    image = app.config['UPLOAD_PATH'] + img
    label = 'bottels'
    print(ConfPercent)
    print(model)
    print(prototxt)
    print("image path is" + image)

    # Load the Caffe model
    net = cv2.dnn.readNetFromCaffe(prototxt, model)
    # Load image fro
    frame = cv2.imread(image)

    frame_resized = cv2.resize(frame, (300, 300))  # resize frame for
    # prediction
    heightFactor = frame.shape[0]/300.0
    widthFactor = frame.shape[1]/300.0

    # MobileNet requires fixed dimensions for input image(s)
    # so we have to ensure that it is resized to 300x300 pixels.
    # set a scale factor to image because network the objects has
    # differents size.
    # We perform a mean subtraction (127.5, 127.5, 127.5)
    # to normalize the input;
    # after executing this command our "blob" now has the shape:
    # (1, 3, 300, 300)
    blob = cv2.dnn.blobFromImage(frame_resized, 0.007843, (300, 300),
                                 (127.5, 127.5, 127.5), False)
    # Set to network the input blob
    net.setInput(blob)
    # Prediction of network
    detections = net.forward()

    frame_copy = frame.copy()
    # Size of frame resize (300x300)
    cols = frame_resized.shape[1]
    rows = frame_resized.shape[0]

    # For get the class and location of object detected,
    # There is a fix index for class, location and confidence
    # value in @detections array .
    for i in range(detections.shape[2]):
        confidence = detections[0, 0, i, 2]  # Confidence of prediction
        if confidence > ConfPercent:  # Filter prediction
            class_id = int(detections[0, 0, i, 1])  # Class label

            # Object location
            xLeftBottom = int(detections[0, 0, i, 3] * cols)
            yLeftBottom = int(detections[0, 0, i, 4] * rows)
            xRightTop = int(detections[0, 0, i, 5] * cols)
            yRightTop = int(detections[0, 0, i, 6] * rows)

            xLeftBottom_ = int(widthFactor * xLeftBottom)
            yLeftBottom_ = int(heightFactor * yLeftBottom)
            xRightTop_ = int(widthFactor * xRightTop)
            yRightTop_ = int(heightFactor * yRightTop)
            # Draw location of object
            cv2.rectangle(frame_resized, (xLeftBottom, yLeftBottom),
                          (xRightTop, yRightTop),
                          (0, 255, 0))

            cv2.rectangle(frame_copy, (xLeftBottom_, yLeftBottom_),
                          (xRightTop_, yRightTop_),
                          (0, 255, 0), -1)
    opacity = 0.3
    cv2.addWeighted(frame_copy, opacity, frame, 1 - opacity, 0, frame)

    count = 0
    for i in range(detections.shape[2]):
        confidence = detections[0, 0, i, 2]  # Confidence of prediction
        if confidence > ConfPercent:  # Filter prediction
            class_id = int(detections[0, 0, i, 1])  # Class label

            # Object location
            xLeftBottom = int(detections[0, 0, i, 3] * cols)
            yLeftBottom = int(detections[0, 0, i, 4] * rows)
            xRightTop = int(detections[0, 0, i, 5] * cols)
            yRightTop = int(detections[0, 0, i, 6] * rows)

            xLeftBottom_ = int(widthFactor * xLeftBottom)
            yLeftBottom_ = int(heightFactor * yLeftBottom)
            xRightTop_ = int(widthFactor * xRightTop)
            yRightTop_ = int(heightFactor * yRightTop)
            cv2.rectangle(frame, (xLeftBottom_, yLeftBottom_),
                          (xRightTop_, yRightTop_),
                          (0, 0, 0), 2)

            # Draw label and confidence of prediction in frame resized
            if class_id in classNames:
                label = classNames[class_id] + ": " + str(confidence)
                labelSize, baseLine = cv2.getTextSize(label,
                                                      cv2.FONT_HERSHEY_TRIPLEX,
                                                      0.8, 1)

                yLeftBottom_ = max(yLeftBottom_, labelSize[1])
                cv2.rectangle(
                    frame,
                    (xLeftBottom_, yLeftBottom_ - labelSize[1]),
                    (xLeftBottom_ + labelSize[0], yLeftBottom_ + baseLine),
                    (255, 255, 255), cv2.FILLED)
                cv2.putText(frame, label, (xLeftBottom_, yLeftBottom_),
                            cv2.FONT_HERSHEY_TRIPLEX, 0.8, (0, 0, 0))
                print(label)
                count = count + 1

    print("total item count", count)
    # cv2.namedWindow("frame", cv2.WINDOW_NORMAL)
    print("before im write")
    cv2.imwrite(IMAGE_PATH + "result.jpeg", frame)
    # cv2.imshow("frame", frame)
    # cv2.waitKey(0)
    print("before im before destroy window")
    # cv2.destroyAllWindows()
    # Detect_result = {'ImposedImage': 'frame', 'ObjCount': count,
    # 'ObjType': type, 'Time': time}
    Detect_result = {'ObjCount': count}
    print(Detect_result)
    return Detect_result


@app.route('/mep/v1/obj_detection/uploadModel', methods=['POST'])
def uploadModel():
    """
    upload model
    :return: html file
    """
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")

    modelInfo = model_info("caffe")
    modelInfo.update_model("caffe", "mobilenet_ssd", "prototext")
    return Response("success")


@app.route('/mep/v1/obj_detection/confidencelevel', methods=['POST'])
def setConfidenceLevel():
    """
    Trigger the video_feed() function on opening "0.0.0.0:5000/video_feed" URL
    :return:
    """
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")

    confidenceLevel = 0.2
    modelInfo = model_info("caffe")
    modelInfo.set_confidence_level(confidenceLevel)
    return Response("success")


@app.route('/mep/v1/obj_detection/detect', methods=['POST'])
def Obj_Detection():
    """
    Trigger the Obj detection on input frame/image
    Input: frame/image
    :return: imposed frame, count, Obj type, time taken by detection
    """
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")

    if 'file' not in request.files:
        raise IOError('No file')

    file = request.files['file']
    if file.filename == '':
        app.logger.info('No file selected for uploading')
        raise IOError('No file')
    if file and allowed_file(file.filename):
        filename = secure_filename(file.filename)
        print('file name', filename)
        file.save(os.path.join(app.config['UPLOAD_PATH'], filename))
        app.logger.info('File successfully uploaded')
        print('file path', app.config['UPLOAD_PATH'] + filename)
        Detect_result = Detection(filename)
    else:
        app.logger.info('Allowed file types are txt, pdf, png, jpg, jpeg, gif')
        return Response("failure")
    return jsonify(Detect_result)


@app.route('/mep/v1/obj_detection/image', methods=['GET'])
def image_download():
    """
    Trigger the Obj detection on input frame/image
    Input: frame/image
    :return: imposed frame, count, Obj type, time taken by detection
    """
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")

    return send_file(IMAGE_PATH + "result.jpeg",
                     attachment_filename='result.jpeg')


def start_server(handler):
    app.logger.addHandler(handler)
    if config.ssl_enabled:
        context = (config.ssl_certfilepath, config.ssl_keyfilepath)
        app.run(host=config.server_address, debug=True, ssl_context=context,
                threaded=True, port=config.server_port)
    else:
        app.run(host=config.server_address, debug=True, threaded=True,
                port=config.server_port)
