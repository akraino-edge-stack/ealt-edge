
# Pre-Setup (Before Application Integration)
Step 1: docker login to akraino dockerhub
username/pwd - ealtedge/Huawei_akraino

or 

pull images manually
docker pull ealtedge/obj-detection
docker pull ealtedge/inventory-be:v1.3
docker pull ealtedge/robo-be
docker pull ealtedge/robo


Step 2: Install Influx DB
- Install local path storage (if by default is not available on edge node)
    1. git clone https://github.com/rancher/local-path-provisioner.git
    2. cd local-path-provisioner
    3. helm install local-path --namespace kube-system ./deploy/chart/ --set storageClass.provisionerName=rancher.io/local
-path --set storageClass.defaultClass=true --set storageClass.name=local-path

- Install influx db
    1. create my-test namespace
    2. helm repo add influxdata https://influxdata.github.io/helm-charts
    3. helm upgrade -i influxdb influxdata/influxdb --set service.type=NodePort --namespace my-test

- Some Pre-Checks
    - Check whether the config file is at location /root/.kube/
    - Check that ports are not already occupied