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

import config
from flask_sslify import SSLify
from flask import Flask, request, jsonify, Response
from flask_cors import CORS
# from camera_driver.capture_frame import VideoCamera, VideoFile
# from capture_frame import VideoCamera, VideoFile
from influxdb import InfluxDBClient
import json
import requests
import os
import cv2
import hashlib
from wand.image import Image
import os.path
from os import path
import base64
import time
import sys

app = Flask(__name__)
CORS(app)
sslify = SSLify(app)
app.config['JSON_AS_ASCII'] = False
app.config['UPLOAD_PATH'] = '/usr/app/images_result/'
app.config['VIDEO_PATH'] = '/usr/app/test/resources/'
app.config['supports_credentials'] = True
app.config['CORS_SUPPORTS_CREDENTIALS'] = True
app.config['MAX_CONTENT_LENGTH'] = 16 * 1024 * 1024
ALLOWED_EXTENSIONS = set(['png', 'jpg', 'jpeg'])
ALLOWED_VIDEO_EXTENSIONS = {'mp4'}
count = 0
listOfMsgs = []
listOfCameras = []
listOfVideos = []
mock_func = 0


class inventory_info:
    """
    Store the data and manage multiple input video feeds
    """
    def __init__(self, current_count=0, total_count=10, time=0):
        self.type = "Shelf_INV1"
        self.labels = "Bottles"
        self.current_count = current_count
        self.total_count = total_count
        self.time = time

    def setcurrentcount(self, current_count):
        self.current_count = current_count

    def settotalcount(self, total_count):
        self.total_count = total_count

    def getcurrentcount(self):
        return self.current_count

    def gettotalcount(self):
        return self.total_count

    def setlabel(self, labels):
        self.labels = labels

    def getlabel(self):
        return self.labels

    def settime(self, time):
        self.labels = time

    def gettime(self):
        return self.time


# temporary copied capture_frame file to this due to docker issue for module
# import
class VideoCamera(object):
    """
    opneCV to capture frame from a camera
    """
    def __init__(self, url):
        self.video = cv2.VideoCapture(url)

    def delete(self):
        self.video.release()

    def get_frame(self):
        """
        get a frame from camera url
        """
        success, image = self.video.read()
        return success, image


class VideoFile(object):
    """
    opneCV to capture frame from a video stream
    """
    def __init__(self, video_name):
        self.video = cv2.VideoCapture(app.config['VIDEO_PATH'] + video_name)

    def delete(self):
        self.video.release()

    def get_frame(self):
        """
        get a frane from stream
        """
        success, image = self.video.read()
        return success, image


def mock_table(inven_info):
    current_count = 3
    labels = "Bottles"
    total_count = 6
    inven_info.setcurrentcount(current_count)
    inven_info.settotalcount(total_count)
    inven_info.setlabel(labels)
    time_sec = time.time()
    local_time = time.ctime(time_sec)
    inven_info.utime = local_time
    # store_data(inven_info)
    local_store(inven_info)


def shelf_inventory(video_capture, camera_info, true=None):
    """
    shelf_inventory
    """
    global count
    global mock_func

    labels = "Bottles"
    count_val = 'ObjCount'
    process_this_frame = 0
    i = 0
    url = config.detection_url + "detect"
    url_get = config.detection_url + "image"

    if mock_func == 1:
        inven_info = inventory_info()
        mock_table(inven_info)
    else:
        while True:
            success, frame = video_capture.get_frame()
            if not success:
                print ('read frame from file is failed')
                break
            i = i+1
            if i < 100:
                continue

            i = 0
            if process_this_frame == 0:

                # info1 = cv2.imencode(".jpg", frame)[1].tobytes()
                imencoded = cv2.imencode(".jpg", frame)[1]
                file = {'file': (
                'image.jpg', imencoded.tostring(), 'image/jpeg',
                {'Expires': '0'})}
                # data = json.loads(requests.post
                #                  (url, data=frame,
                #                  verify=config.ssl_cacertpath).text)

                res = requests.post(url, files=file)
                data = json.loads(res.text)

                ## get image
                response = requests.get(url_get)

                file = open(app.config['UPLOAD_PATH'] + "sample_image.jpg",
                            "wb")
                file.write(response.content)
                file.close()

                inven_info = inventory_info()
                current_count = data[count_val]
                total_count = inven_info.current_count + inven_info.total_count
                inven_info.setcurrentcount(current_count)
                inven_info.settotalcount(total_count)
                inven_info.setlabel(labels)
                time_sec = time.time()
                local_time = time.ctime(time_sec)
                inven_info.time = local_time
                store_info_db(inven_info)
                # local_store(inven_info)
                time.sleep(0.01)


def local_store(inven_info):
    """
    store "shelf" data to array

    :param inven_info: Inventry object
    :return: None
    """
    if len(listOfMsgs) >= 100:
        listOfMsgs.pop()

    time_sec = time.time()
    local_time = time.ctime(time_sec)
    print (local_time)
    newdict = {"shelfName": inven_info.type, "ObjType": inven_info.labels,
               "currentCount": inven_info.current_count,
               "totalCount": inven_info.total_count,
               "time": local_time}
    listOfMsgs.insert(0, newdict)


def store_info_db(inven_info):
    """
    Send "shelf" data to InfluxDB

    :param inven_info: Inventry object
    :return: None
    """
    global db_client
    json_body = [
        {
            "measurement": inven_info.type,
            "tags": {
                "object": "bottles",
            },
            "fields": {
                "time": inven_info.time,
                "Current-Count": inven_info.current_count,
                "Total-Count": inven_info.total_count,
            }
        }]
    db_client.write_points(json_body)


def retrive_info_db():
    """
    Send "shelf" data to InfluxDB

    :param inven_info: Inventry object
    :return: None
    """
    global db_client

    # get data last n data points from DB
    result = db_client.query('select * from Shelf_INV1 order by desc limit '
                             '10;')

    # Get points and iterate over each record
    points = result.get_points(tags={"object": "bottles"})

    # clear the msg list
    # listOfMsgs.clear()
    del listOfMsgs[:]

    # iterate points and fill the records and insert to list
    for point in points:
        print("Current-Count: %i, Total-Count: %i, Time: %s" % (point[
                                                             'Current-Count'],
                                          point['Total-Count'],
                                          point['time']))
        newdict = {"shelfName": 'Shelf_INV1', "ObjType": "bottles",
                   "currentCount": point['Current-Count'],
                   "totalCount": point['Total-Count'],
                   "time": point['time']}
        listOfMsgs.insert(0, newdict)


def create_database():
    """
    Connect to InfluxDB and create the database

    :return: None
    """
    global db_client
    proxy = {"http": "http://{}:{}".format(config.IPADDRESS, config.PORT)}
    db_client = InfluxDBClient(host=config.IPADDRESS, port=config.PORT,
                proxies=proxy, database=config.DATABASE_NAME)
    db_client.create_database(config.DATABASE_NAME)


@app.route('/v1/inventry/table', methods=['GET'])
def inventry_table():
    """
    return inventry table

    :return: inventry table
    """
    retrive_info_db()
    table = {"InventryData": listOfMsgs}
    return jsonify(table)


@app.route('/v1/inventry/image', methods=['GET'])
def detected_image():
    """
    detect images with imposed

    :return: result image
    """
    detected_image = app.config['UPLOAD_PATH'] + 'sample_image.jpg'
    print('file exits:', str(path.exists(detected_image)))
    status = str(path.exists(detected_image))
    if status == 'True':
        ## as base64 string
        with open(detected_image, "rb") as img_file:
            jpeg_bin = base64.b64encode(img_file.read())

        response = {'image': jpeg_bin}
        return jsonify(response)
    else:
        response = {'image' : 'null'}
        print('file not exist')
        return jsonify(response)


def allowed_videofile(filename):
    """
    File types to upload:mp4
    param: filename:
    """
    return '.' in filename and \
           filename.rsplit('.', 1)[1].lower() in ALLOWED_VIDEO_EXTENSIONS


@app.route('/v1/monitor/video', methods=['POST'])
def upload_video():
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    print ("videpath:" + app.config['VIDEO_PATH'])
    if 'file' in request.files:
        files = request.files.getlist("file")
        for file in files:
            if allowed_videofile(file.filename):
                file.save(os.path.join(app.config['VIDEO_PATH'],
                                       file.filename))
                print('file path is:', app.config['VIDEO_PATH'] +
                                       file.filename)
            else:
                raise IOError('video format error')
                msg = {"responce": "failure"}
                return jsonify(msg)
    msg = {"responce":"success"}
    return jsonify(msg)

def hash_func(camera_info):
    hash_string = camera_info["cameraNumber"] + \
                  camera_info["cameraLocation"] + \
                  camera_info["rtspUrl"]
    # readable_hash = hashlib.sha256(str(hash_string).encode(
    # 'utf-8')).hexdigest()
    readable_hash = hash(hash_string)
    if readable_hash < 0:
        readable_hash += sys.maxsize
    print (readable_hash)
    return readable_hash


@app.route('/v1/monitor/cameras', methods=['POST'])
def add_camera():
    camera_detail = request.json
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    camera_id = hash_func(camera_detail)
    for camera_info in listOfCameras:
        if camera_id == camera_info["cameraID"]:
            msg = {"responce": "failure"}
            return jsonify(msg)
            break
    camera_info = {"cameraID": camera_id,
                   "cameraNumber": camera_detail["cameraNumber"],
                   "rtspUrl": camera_detail["rtspUrl"],
                   "cameraLocation": camera_detail["cameraLocation"]}
    listOfCameras.append(camera_info)
    msg = {"responce":"success"}
    return jsonify(msg)


@app.route('/v1/monitor/cameras/<cameraID>', methods=['GET'])
def get_camera(cameraID):
    """
    register camera with location
    """
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    valid_id = 0
    for camera_info in listOfCameras:
        cameraID = int(cameraID)
        if cameraID == camera_info["cameraID"]:
            valid_id = 1
            break

    if valid_id == 0:
        app.logger.info ("camera ID is not valid")
        msg = {"responce":"failure"}
        return jsonify(msg)

    if "mp4" in camera_info["rtspUrl"]:
        video_file = VideoFile(camera_info["rtspUrl"])
        video_dict = {camera_info["cameraNumber"]: video_file}
        listOfVideos.append(video_dict)
        # return Response(shelf_inventory(video_file, camera_info[
        # "cameraNumber"]),
        #                mimetype='multipart/x-mixed-replace; boundary=frame')
        shelf_inventory(video_file, camera_info["cameraNumber"])
        app.logger.info("get_camera: Added json")
        msg = {"responce":"success"}
        return jsonify(msg)

    else:
        video_file = VideoCamera(camera_info["rtspUrl"])
        video_dict = {camera_info["cameraNumber"]: video_file}
        listOfVideos.append(video_dict)
        return Response(shelf_inventory(video_file, camera_info["cameraNumber"]),
                        mimetype='multipart/x-mixed-replace; boundary=frame')


@app.route('/v1/monitor/cameras/<camera_name>', methods=['DELETE'])
def delete_camera(camera_name):
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    for video1 in listOfVideos:
        if camera_name in video1:
            video_obj = video1[camera_name]
            video_obj.delete()
    for camera in listOfCameras:
        if camera_name == camera["cameraNumber"]:
            listOfCameras.remove(camera)
    for msg in listOfMsgs:
        if camera_name in msg["msg"]:
            listOfMsgs.remove(msg)
    return Response("success")


@app.route('/v1/monitor/cameras')
def query_cameras():
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    camera_info = {"roboCamera": listOfCameras}
    return jsonify(camera_info)


@app.route('/', methods=['GET'])
def hello_world():
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    return Response("Hello MEC Developer")


def start_server(handler):
    app.logger.addHandler(handler)
    create_database()
    if config.ssl_enabled:
        context = (config.ssl_certfilepath, config.ssl_keyfilepath)
        app.run(host=config.server_address, debug=True, ssl_context=context,
                threaded=True, port=config.server_port)
    else:
        app.run(host=config.server_address, debug=True, threaded=True,
                port=config.server_port)
