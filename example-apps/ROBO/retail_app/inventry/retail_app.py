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
#from capture_frame import VideoCamera, VideoFile
#from influxdb import InfluxDBClient
import json
import time
import requests
import os
import cv2


app = Flask(__name__)
CORS(app)
sslify = SSLify(app)
app.config['JSON_AS_ASCII'] = False
app.config['UPLOAD_PATH'] = '/usr/app/images/'
app.config['supports_credentials'] = True
app.config['CORS_SUPPORTS_CREDENTIALS'] = True
app.config['MAX_CONTENT_LENGTH'] = 16 * 1024 * 1024
ALLOWED_EXTENSIONS = set(['png', 'jpg', 'jpeg'])
ALLOWED_VIDEO_EXTENSIONS = {'mp4'}
count = 0
listOfMsgs = []
listOfCameras = []
listOfVideos = []
mock_func = 1

class inventory_info:
    """
    Store the data and manage multiple input video feeds
    """
    def __init__(self, current_count=0, total_count=0, time=0):
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
        self.video = cv2.VideoCapture("./test/resources/" + video_name)

    def delete(self):
        self.video.release()

    def get_frame(self):
        """
        get a frane from stream
        """
        success, image = self.video.read()
        return success, image


def store_data(inventory_info):
    """
    store time series data in influx db
    """
    # TODO config, schema table, DB, fill data set
    create_database()
    store_info_db(inventory_info)


def mock_table(inven_info):
    current_count = 3
    labels = "Bottles"
    total_count = 6
    inven_info.setcurrentcount(current_count)
    inven_info.settotalcount(total_count)
    inven_info.setlabel(labels)
    inven_info.utime = time.time()
    # store_data(inven_info)
    local_store(inven_info)


def shelf_inventory(video_capture, camera_info, true=None):
    """
    shelf_inventory
    """
    global count
    global mock_func

    labels = "bottles"
    process_this_frame = 0
    if mock_func is 1:
        inven_info = inventory_info()
        mock_table(inven_info)
    else:
        while True:
            success, frame = video_capture.get_frame()
            if not success:
                break
            if process_this_frame == 0:
                url = config.detection_url + "/v1/obj_detection/detect"
                # info1 = cv2.imencode(".jpg", rgb_small_frame)[1].tobytes()
                data = json.loads(requests.post(url, data=frame,
                                            verify=config.ssl_cacertpath).text)
        inven_info = inventory_info()
        current_count = data[count]
        labels = data[labels]
        total_count = inven_info.current_count + inven_info.total_count
        inven_info.setcurrentcount(current_count)
        inven_info.settotalcount(total_count)
        inven_info.setlabel(labels)
        inven_info.utime = time.time()
        #store_data(inven_info)
        local_store(inven_info)


def local_store(inven_info):
    """
    store "shelf" data to array

    :param inven_info: Inventry object
    :return: None
    """
    if len(listOfMsgs) >= 100:
        listOfMsgs.pop()
    newdict = {"shelfName": inven_info.type, "ObjType": inven_info.labels,
               "currentCount": inven_info.current_count,
               "totalCount": inven_info.total_count,
               "time": time.time()}
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
                "Current Count": inven_info.current_count,
                "Total Count": inven_info.total_count,
            }
        }]
    db_client.write_points(json_body)


def create_database():
    """
    Connect to InfluxDB and create the database

    :return: None
    """
    global db_client

    proxy = {"http": "http://{}:{}".format(config.IPADDRESS, config.PORT)}
#    db_client = InfluxDBClient(host=config.IPADDRESS, port=config.PORT,
 #                              proxies=proxy, database=config.DATABASE_NAME)
#    db_client.create_database(config.DATABASE_NAME)


@app.route('/v1/inventry/table', methods=['GET'])
def inventry_table():
    """
    return inventry table

    :return: inventry table
    """
    return jsonify(listOfMsgs)


@app.route('/v1/inventry/image', methods=['GET'])
def inventry_table():
    """
    return inventry table

    :return: inventry table
    """
    return jsonify(listOfMsgs)

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
    if 'file' in request.files:
        files = request.files.getlist("file")
        for file in files:
            if allowed_videofile(file.filename):
                file.save(os.path.join(app.config['VIDEO_PATH'], file.filename))
            else:
                raise IOError('video format error')
    return Response("success")


@app.route('/v1/monitor/cameras', methods=['POST'])
def add_camera():
    camera_info = request.json
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    camera_info = {"name": camera_info["name"],
                   "rtspurl": camera_info["rtspurl"],
                   "location": camera_info["location"]}
    listOfCameras.append(camera_info)
    return Response("success")


@app.route('/v1/monitor/cameras/<name>/<rtspurl>/<location>', methods=['GET'])
def get_camera(name, rtspurl, location):
    """
    register camera with location
    """
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    camera_info = {"name": name, "rtspurl": rtspurl, "location": location}
    if "mp4" in camera_info["rtspurl"]:
        video_file = VideoFile(camera_info["rtspurl"])
        video_dict = {camera_info["name"]: video_file}
        listOfVideos.append(video_dict)
        return Response(shelf_inventory(video_file, camera_info["name"]),
                        mimetype='multipart/x-mixed-replace; boundary=frame')
    else:
        video_file = VideoCamera(camera_info["rtspurl"])
        video_dict = {camera_info["name"]: video_file}
        listOfVideos.append(video_dict)
        return Response(shelf_inventory(video_file, camera_info["name"]),
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
        if camera_name == camera["name"]:
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
    return jsonify(listOfCameras)


@app.route('/', methods=['GET'])
def hello_world():
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    return Response("Hello MEC Developer")


def start_server(handler):
    app.logger.addHandler(handler)
    if config.ssl_enabled:
        context = (config.ssl_certfilepath, config.ssl_keyfilepath)
        app.run(host=config.server_address, debug=True, ssl_context=context,
                threaded=True, port=config.server_port)
    else:
        app.run(host=config.server_address, debug=True, threaded=True,
                port=config.server_port)
