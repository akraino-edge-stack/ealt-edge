/*
 * Copyright 2020 Huawei Technologies Co., Ltd.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package util

import (
	"github.com/sirupsen/logrus"
	"os"
)

func GetLogger(logFile string, loggerLevel logrus.Level, file *os.File) *logrus.Logger {

	logger := logrus.New()
	logger.SetOutput(file)
	logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})
	logger.SetLevel(loggerLevel)
	logger.Info("logger created")
	return logger
}