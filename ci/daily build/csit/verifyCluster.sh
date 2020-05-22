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

count=0
total=2
result_nginx="failure"
result_pods="failure"

if [[ $CLUSTER_INFO != "" ]]; then
   kubectl create -f ~/testk8s-kubernetes.yaml
   count=$((count+1))
else
   echo "No kubernetes cluster present"
fi

sleep 20

echo "Kubectl deployments........................................."
kubectl get deployments

echo "Test Case: Nginx-Deployment started"

DEPLOY_CONDIT=$(kubectl get pods \
              --field-selector=status.phase==Running \
              | grep kubernets-deployment \
              | grep -c Running)

if [[ $DEPLOY_CONDIT == 1 ]]; then
   result_nginx="success";
   echo $result_nginx
fi

echo "Kubectl pods in default namespace................................."
kubectl get pods

echo "-------------------------------------------------------------------"
echo "-------------------------------------------------------------------"

sleep 20

echo "Test Case: Pods status check started"

PODS_NOT_RUN_COUNT=$(kubectl get pods \
                    --field-selector=status.phase!=Running \
                    | grep -c STATUS)

if [[ $PODS_NOT_RUN_COUNT > 0 ]]; then > /dev/null 2>&1
   result_pods="failure";
   count=$((count+1))
else
   count=$((count+1))
   result_pods="success";
   echo $result_pods
fi


echo "-------------------------------------------------------------------"
echo "|                        Total CSIT Tests: $count                      |"
echo "|-----------------------------------------------------------------|"
echo "|           TEST CASE NAME          |           RESULT            |"
echo "|-----------------------------------------------------------------|"
echo "|                                   |                             |"
echo "|          Nginx-Deployment         |           $result_nginx           |"
echo "|                                   |                             |"
echo "|          Pods status check        |           $result_pods           |"
echo "|                                   |                             |"
echo "|-----------------------------------------------------------------|"
echo "|              Executed Total CSIT Tests: $count                       |"
echo "-------------------------------------------------------------------"
