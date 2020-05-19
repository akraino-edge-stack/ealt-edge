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

#To check the pods which is not running
if kubectl get pods | grep -c -e STATUS -e Running > 0; then
PODS_NOT_RUNNING=$(kubectl get pods --field-selector=status.phase!=Running)
fi

echo "-------------------------------------------------------------------"
echo
echo $PODS_NOT_RUNNING > /dev/null 2>&1
echo
echo "-------------------------------------------------------------------"
