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
from flask import Flask, request, Response
from flask_cors import CORS


app = Flask(__name__)
CORS(app)
sslify = SSLify(app)
app.config['JSON_AS_ASCII'] = False
app.config['UPLOAD_PATH'] = '/usr/app/images/'
app.config['supports_credentials'] = True
app.config['CORS_SUPPORTS_CREDENTIALS'] = True
app.config['MAX_CONTENT_LENGTH'] = 16 * 1024 * 1024
ALLOWED_EXTENSIONS = set(['png', 'jpg', 'jpeg'])
count = 0
listOfMsgs = []


class shelf_inventry():
    """
    def __init__(self, url):
        # self.video = cv2.VideoCapture(url)

    def delete(self):
        # self.video.release()
        return
    """


def store_data():
    """
    store time series data in influx db
    """
    # TODO config, schema table, DB, fill data set


def obj_detect():
    """
    detect obj and count for self
    """


@app.route('/v1/monitor/cameras', methods=['POST'])
def add_camera():
    camera_info = request.json
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    camera_info = {"name": camera_info["name"],
                   "rtspurl": camera_info["rtspurl"],
                   "location": camera_info["location"]}
    # listOfCameras.append(camera_info)
    return Response("success")


@app.route('/v1/monitor/cameras/<name>/<rtspurl>/<location>', methods=['GET'])
def get_camera(name, rtspurl, location):
    """
    register camera with location
    """
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    # camera_info = {"name": name, "rtspurl": rtspurl, "location": location}
    """
    if "mp4" in camera_info["rtspurl"]:
        # video_file = VideoFile(camera_info["rtspurl"])
        # video_dict = {camera_info["name"]:video_file}
        # listOfVideos.append(video_dict)
        # return Response(video(video_file, camera_info["name"]),
                        # mimetype='multipart/x-mixed-replace; boundary=frame')
    else:
        # video_file = VideoCamera(camera_info["rtspurl"])
        # video_dict = {camera_info["name"]: video_file}
        # listOfVideos.append(video_dict)
        # return Response(video(video_file, camera_info["name"]),
                     # mimetype='multipart/x-mixed-replace; boundary=frame')
        return Response("success")
    """


@app.route('/v1/monitor/cameras/<camera_name>', methods=['DELETE'])
def delete_camera(camera_name):
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    """
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
    """


@app.route('/v1/monitor/cameras')
def query_cameras():
    app.logger.info("Received message from ClientIP [" + request.remote_addr
                    + "] Operation [" + request.method + "]" +
                    " Resource [" + request.url + "]")
    # return jsonify(listOfCameras)
    return Response("success")


def start_server(handler):
    app.logger.addHandler(handler)
    if config.ssl_enabled:
        context = (config.ssl_certfilepath, config.ssl_keyfilepath)
        app.run(host=config.server_address, debug=True, ssl_context=context,
                threaded=True, port=config.server_port)
    else:
        app.run(host=config.server_address, debug=True, threaded=True,
                port=config.server_port)
