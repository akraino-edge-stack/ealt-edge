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

import cv2


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
