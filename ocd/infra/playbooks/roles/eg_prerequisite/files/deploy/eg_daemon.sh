#!/bin/bash

function _help_insecure_registry()
{
  grep  -i "insecure-registries" /etc/docker/daemon.json | grep "REGISTRIES_IP:REGISTRIES_PORT" >/dev/null 2>&1
  if [  $? != 0 ]; then
    mkdir -p /etc/docker
cat <<EOF | tee /etc/docker/daemon.json
{
    "insecure-registries" : ["REGISTRIES_IP:REGISTRIES_PORT"]
}
EOF
    service docker restart
  fi
}

##############################################################
############################################
function main(){
    _help_insecure_registry
}
#########################################
#skip main in case of source
    main $@
######################
