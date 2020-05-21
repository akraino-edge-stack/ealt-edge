#!/bin/bash -ex
##############################################################################
# Copyright (c) 2020 Huawei Tech and others.
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

CLUSTER_INFO=$(kubectl cluster-info)
RUNNING_PODS_COUNT=$(kubectl get pods | grep -c -e STATUS -e Running)

if [[ $CLUSTER_INFO != "" ]]; then
   kubectl create -f ~/testk8s-kubernetes.yaml
else
   echo "No kubernetes cluster present"
fi

sleep 60

echo "Kubectl deployments........................................."
kubectl get deployments

echo "Kubectl pods in default namespace............................"
kubectl get pods

echo "-------------------------------------------------------------------"
echo "-------------------------------------------------------------------"

sleep 60

echo "Checking for Pods not in running status in default namespace"

if [[ $RUNNING_PODS_COUNT > 0 ]]; then
   kubectl get pods --field-selector=status.phase!=Running
else
   echo "No Pods are presently running"
fi

echo "-------------------------------------------------------------------"
echo "-------------------------------------------------------------------"
