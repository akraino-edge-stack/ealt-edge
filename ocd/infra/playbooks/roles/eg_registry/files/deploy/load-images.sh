#!/bin/bash
TARBALL_PATH=/tmp/eg_registry/deploy/

function _load_and_run_docker_registry()
{
  
    docker ps | grep registry >/dev/null
    if [ $? != 0 ]; then
      cd "$TARBALL_PATH"/registry
      docker load --input registry-2.tar.gz
      docker run -d -p 5000:5000 --restart=always --name registry registry:2
    fi
}

function _load_swr_images_and_push_to_private_registry()
{
  IP=REGISTRIES_IP
  PORT="REGISTRIES_PORT"
  cd "$TARBALL_PATH"/eg_swr_images

  for f in *.tar.gz;
  do
    cat $f | docker load
      IMAGE_NAME=`echo $f|rev|cut -c8-|rev|sed -e "s/\#/:/g" | sed -e "s/\@/\//g"`;
      docker image tag $IMAGE_NAME $IP:$PORT/$IMAGE_NAME
      docker push $IP:$PORT/$IMAGE_NAME
  done
}

##############################################################
############################################
function main(){
    _load_and_run_docker_registry
  _load_swr_images_and_push_to_private_registry
}
#########################################
#skip main in case of source
    main $@
######################
