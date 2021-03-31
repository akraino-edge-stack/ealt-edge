
# Pre-Setup (Before Application Integration)
Step 1: docker login to akraino dockerhub
username/pwd - ealtedge/Huawei_akraino

or 

pull images manually
docker pull ealtedge/pcbdefdet-be
docker pull ealtedge/pcbui


Step 2: pretrained model for defect detection
- copy pretrained model to Host machine where PCB containers are running
    1. cloen model from gitlab git clone https://gitlab.com/gauravagrawal/ealtedge/-/tree/master/PCB-Def-Det-Models
    2. create folder /root/pcb-def-det/model in host machine.
    3. scp all files in models cloned to above path. 

Step 3: iput images for defect detection 
- image are copied to container images currently. if want to use different image, need to follow below steps
    1. download application code and go example-apps/PDD/pcb-defect-detection/input_image 
    2. copy the input image to one of folder.
    3. build the docker "docker build -t ealtedge/pcbdefdet-be ." and push to docker hub
    4. now new images will be avaible in container for detection
    5. in future this can be enhance to feed from UI

- Some Pre-Checks
    - Check that ports are not already occupied
    - UI running on "30002" node port of VM where containers are running.
