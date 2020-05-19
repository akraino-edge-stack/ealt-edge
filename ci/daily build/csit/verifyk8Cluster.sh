#!/bin/bash -ex
##############################################################################
# Copyright (c) 2019 Huawei Tech and others.
#
# All rights reserved. This program and the accompanying materials
# are made available under the terms of the Apache License, Version 2.0
# which accompanies this distribution, and is available at
# http://www.apache.org/licenses/LICENSE-2.0
##############################################################################

KUBERNETES=~/testk8s-kubernetes.yaml

cat <<EOF > "${KUBERNETES}"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kubernets-deployment
  labels:
    app: nginx
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:1.15.12
          ports:
            - containerPort: 80
              hostPort: 80
EOF

#check if kubernetes is already running
if ! kubectl cluster-info ; then
  kubectl create -f ~/testk8s-kubernetes.yaml
fi

#To check whether the kubernetes has started successfully
retry=10
while [ $retry -gt 0 ]
do
  if [ 2 == "$(kubectl get pods | grep -c -e STATUS -e Running)" ]; then
    break
  fi
  ((retry-=1)) 
  sleep 10
done
[ $retry -gt 0 ] || exit 1
