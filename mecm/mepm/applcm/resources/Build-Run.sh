#AppLcm broker compile and build docker image
cd ../broker
. docker-build.sh


#helmplugin compile and build docker image
cd ../k8shelm
. docker-build.sh

#Run docker images:
cd ../resources
sudo docker-compose up -d
