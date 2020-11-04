OFFLINE_MODE=muno
EG_NODE_DEPLOY_IP=10.10.0.202
EG_NODE_MASTER_IPS=119.8.41.103
EG_NODE_WORKER_IPS=119.8.41.103
EG_NODE_EDGE_MP1=eth1
EG_NODE_EDGE_MM5=eth2
PLATFORM_DIR=$PWD
DEVELOPER_PORT=30092
APPSTORE_PORT=30091
MECM_PORT=30093
USER_MGMT=30067
TARBALL_PATH=$PWD
#MASTER_IP=119.8.41.103
#EG_NODE_WORKER_IPS=10.10.0.116
#EG_NODE_MASTER_IPS=10.10.0.202
#EG_NODE_CONTROLLER_MASTER_IPS=10.10.0.202
#EG_NODE_CONTROLLER_WORKER_IPS=10.10.0.202
#EG_NODE_EDGE_MASTER_IPS=119.8.41.103
#EG_NODE_EDGE_WORKER_IPS=10.10.0.116
#WORKER_LIST=119.8.41.103
#EG_NODE_DEPLOY_IP=119.8.41.103


function main(){
kubectl create secret docker-registry swrregcred \   
 --docker-server=https://swr.ap-southeast-1.myhuaweicloud.com/v2/ \   
 --docker-username=ap-southeast-1@0K1RQ5EAF2QRKQWQNFY0 \   
 --docker-password=5468d8a0ebc64936a8196742d601bb95b99f1d94ad19c686488831e3dae79bb3
kubectl patch serviceaccount default -p '{"imagePullSecrets": [{"name": "swrregcred"}]}'
 _deploy_eg
_deploy_controller
_deploy_edge
}

############################################################
#if [[ $OFFLINE_MODE == "aio" ]]; then
#    CHART_PREFIX="$TARBALL_PATH/helm/helm-charts/"
#    CHART_SUFFIX="-1.0.1.tgz"
#    PROM_CHART_SUFFIX="-9.3.1.tgz"
#    GRAFANA_CHART_SUFFIX="-5.5.5.tgz"
#    REGISTRY_URL=""
#  else
#    CHART_PREFIX="/root/edgeakraino/jenkins/work/workspace/ealt-edge-deploy-virtual-daily-master/ocd/infra/playbooks/roles/edgegallery_muno/files/deploy/helm-charts/"
#    CHART_SUFFIX="-1.0.1.tgz"
    PRIVATE_REGISTRY_IP=$(echo $EG_NODE_DEPLOY_IP|cut -d "," -f1)
#    REGISTRY_URL="$PRIVATE_REGISTRY_IP:5000/"
#  fi
############################################################

#***********************************************************
#************This is for all********************************

function _deploy_eg()
{
  #password_less_ssh_check $EG_NODE_MASTER_IPS $EG_NODE_WORKER_IPS
  MASTER_IP=$(echo $EG_NODE_MASTER_IPS|cut -d "," -f1)
  setup_eg_ecosystem
    make_remote_dir $MASTER_IP $EG_NODE_WORKER_IPS
    mkdir -p $HOME/.kube
    scp root@$MASTER_IP:/root/.kube/config $HOME/.kube/
  #configure_eg_ecosystem_on_remote $MASTER_IP $EG_NODE_WORKER_IPS
  info "[Going inside deploy eg funcstion  ...]" $BLUE
  _eg_deploy all $EG_NODE_DEPLOY_IP $MASTER_IP
}

#*******************************************************************
#***********************Deploy Controller***************************

function _deploy_controller()
{
  #password_less_ssh_check $EG_NODE_CONTROLLER_MASTER_IPS $EG_NODE_CONTROLLER_WORKER_IPS
  MASTER_IP=$(echo $EG_NODE_CONTROLLER_MASTER_IPS|cut -d "," -f1)
  setup_eg_ecosystem
    make_remote_dir $MASTER_IP $EG_NODE_CONTROLLER_WORKER_IPSmkdir -p $HOME/.kube
    scp root@$MASTER_IP:/root/.kube/config $HOME/.kube/
  #configure_eg_ecosystem_on_remote $MASTER_IP $EG_NODE_CONTROLLER_WORKER_IPS
  info "[Going inside deploy controller funcstion  ...]" $BLUE
  _eg_deploy controller $EG_NODE_DEPLOY_IP $MASTER_IP
}

#************************************************************************
#*************************Deploy edge************************************

function _deploy_edge()
{
  #password_less_ssh_check $EG_NODE_EDGE_MASTER_IPS $EG_NODE_EDGE_WORKER_IPS
  MASTER_IP=$(echo $EG_NODE_EDGE_MASTER_IPS|cut -d "," -f1)
  setup_eg_ecosystem
  make_remote_dir $MASTER_IP $EG_NODE_EDGE_WORKER_IPS
    mkdir -p $HOME/.kube
    scp root@$MASTER_IP:/root/.kube/config $HOME/.kube/
  #configure_eg_ecosystem_on_remote  $MASTER_IP $EG_NODE_EDGE_WORKER_IPS
  info "[Going inside deploy edge funcstion  ...]" $BLUE
  _eg_deploy edge $EG_NODE_DEPLOY_IP $MASTER_IP
}

#***************************************************************************
#*****************************_eg_deploy called by adge,controller,all******

function _eg_deploy()
{
  FEATURE=$1
  DEPLOY_NODE_IP=$2
  MASTER_IP=$3
  info "[Going inside deploy  funcstion  ...]" $BLUE
  install_EdgeGallery $FEATURE $MASTER_IP
}

#****************************************************************************
#*******************************Install adgegallery calling by _eg_deploy****


function install_EdgeGallery ()
{
  FEATURE=$1
  NODEIP=$2
    kubectl create secret generic edgegallery-ssl-secret \
    --from-file=keystore.p12=$PLATFORM_DIR/conf/keys/keystore.p12 \
    --from-literal=keystorePassword=te9Fmv%qaq \
    --from-literal=keystoreType=PKCS12 \
    --from-literal=keyAlias=edgegallery \
    --from-file=trust.cer=$PLATFORM_DIR/conf/keys/ca.crt \
    --from-file=server.cer=$PLATFORM_DIR/conf/keys/tls.crt \
    --from-file=server_key.pem=$PLATFORM_DIR/conf/keys/encryptedtls.key \
    --from-literal=cert_pwd=te9Fmv%qaq

    info "[Going inside install edgegallery funcstion  ...]" $BLUE
    install_service-center
    install_user-mgmt
    install_mecm-meo
    install_mecm-fe
    install_appstore
    install_developer
}

#************************************************************************************
#***************************install_service-center***********************************

function install_service-center ()
{
  info "[Deploying ServiceCenter  ...]" $BLUE
  #helm install service-center-edgegallery edgegallery/servicecenter 
  #helm install service-center-edgegallery "$CHART_PREFIX"servicecenter"$CHART_SUFFIX"
  info "[Going inside install service cernter funcstion  ...]" $BLUE
  #helm install servicecenter /root/edgeakraino/jenkins/work/workspace/ealt-edge-deploy-virtual-daily-master/ocd/infra/playbooks/roles/edgegallery_muno/files/deploy/helm-charts/servicecenter-1.0.1.tgz
  helm install service-center-edgegallery "$CHART_PREFIX"edgegallery/servicecenter"$CHART_SUFFIX" \
  --set global.ssl.enabled=true \
  --set global.ssl.secretName=edgegallery-ssl-secret
  if [ $? -eq 0 ]; then
    wait "service-center" 1
    info "[Deployed ServiceCenter  ....]" $GREEN
  else
    info "[ServiceCenter Deployment Failed]" $RED
  fi
}

#******************************************************************************************
#**************************************install_user-mgmt***********************************

function install_user-mgmt ()
{
  info "[Deploying UserMgmt  ........]" $BLUE

  ## Create a jwt secret for usermgmt
  kubectl create secret generic user-mgmt-jwt-secret \
    --from-file=publicKey=$PLATFORM_DIR/conf/keys/rsa_public_key.pem \
    --from-file=encryptedPrivateKey=$PLATFORM_DIR/conf/keys/encrypted_rsa_private_key.pem \
    --from-literal=encryptPassword=te9Fmv%qaq

  #helm install user-mgmt-edgegallery "$CHART_PREFIX"usermgmt"$CHART_SUFFIX" 
info "[Helm repo command started for usermanagement]" $GREEN
  #helm install usermgmt /root/edgeakraino/jenkins/work/workspace/ealt-edge-deploy-virtual-daily-master/ocd/infra/playbooks/roles/edgegallery_muno/files/deploy/helm-charts/usermgmt-1.0.1.tgz
  helm install user-mgmt-edgegallery "$CHART_PREFIX"edgegallery/usermgmt"$CHART_SUFFIX" \
  --set global.oauth2.clients.appstore.clientUrl=https://$EG_NODE_DEPLOY_IP:$APPSTORE_PORT,\
global.oauth2.clients.developer.clientUrl=https://$EG_NODE_DEPLOY_IP:$DEVELOPER_PORT,\
global.oauth2.clients.mecm.clientUrl=https://$EG_NODE_DEPLOY_IP:$MECM_PORT, \
--set jwt.secretName=user-mgmt-jwt-secret \
--set global.ssl.enabled=true \
--set global.ssl.secretName=edgegallery-ssl-secret
  if [ $? -eq 0 ]; then
    wait "user-mgmt-redis" 1
    wait "user-mgmt-postgres" 1
    wait "user-mgmt" 3
    info "[Deployed UserMgmt  .........]" $GREEN
  else
    info "[UserMgmt Deployment Failed .]" $RED
  fi
}

#**********************************************************************
#*************************install_mecm-meo*****************************

function install_mecm-meo ()
{

  info "[Deploying MECM-MEO  ........]" $BLUE
  ## Create a keystore secret
  kubectl create secret generic mecm-ssl-secret \
  --from-file=keystore.p12=$PLATFORM_DIR/conf/keys/keystore.p12 \
  --from-file=keystore.jks=$PLATFORM_DIR/conf/keys/keystore.jks \
  --from-literal=keystorePassword=te9Fmv%qaq \
  --from-literal=keystoreType=PKCS12 \
  --from-literal=keyAlias=edgegallery \
  --from-literal=truststorePassword=te9Fmv%qaq

  ## Create a mecm-meo secret with postgres_init.sql file to create necessary db's
  info "[Deploying mecm-meo stated deployment   ...]" $BLUE
  kubectl create secret generic edgegallery-mecm-secret \
    --from-file=postgres_init.sql=$PLATFORM_DIR/conf/keys/postgres_init.sql \
    --from-literal=postgresPassword=te9Fmv%qaq \
    --from-literal=postgresApmPassword=te9Fmv%qaq \
    --from-literal=postgresAppoPassword=te9Fmv%qaq \
    --from-literal=postgresInventoryPassword=te9Fmv%qaq \
    --from-literal=edgeRepoUserName=admin	 \
    --from-literal=edgeRepoPassword=admin123

  #helm install mecm-meo-edgegallery "$CHART_PREFIX"mecm-meo"$CHART_SUFFIX" 
  info "[Deploying mecm-meo started helm calling repo  ...]" $BLUE
  #helm install mecm-meo /root/edgeakraino/jenkins/work/workspace/ealt-edge-deploy-virtual-daily-master/ocd/infra/playbooks/roles/edgegallery_muno/files/deploy/helm-charts/mecm-meo-1.0.1.tgz
helm install mecm-meo-edgegallery "$CHART_PREFIX"edgegallery/mecm-meo"$CHART_SUFFIX" \
     --set ssl.secretName=mecm-ssl-secret \
    --set mecm.secretName=edgegallery-mecm-secret
  if [ $? -eq 0 ]; then
    wait "mecm-inventory" 1
    wait "mecm-appo" 1
    wait "mecm-apm" 1
    wait "mecm-postgres" 1
    info "[Deployed MECM-MEO  .........]" $GREEN
  else
    info "[MECM-MEO Deployment Failed  ]" $RED
  fi
}


#************************************************************************************
#*******************************install_mecm-fe**************************************

function install_mecm-fe ()
{
  info "[Deploying MECM-FE  ........]" $BLUE

  #helm install mecm-fe-edgegallery "$CHART_PREFIX"mecm-fe"$CHART_SUFFIX"
 info "[Deploying start mecm-fe helm calling   ...]" $BLUE 
  #helm install mecm-fe /root/edgeakraino/jenkins/work/workspace/ealt-edge-deploy-virtual-daily-master/ocd/infra/playbooks/roles/edgegallery_muno/files/deploy/helm-charts/mecm-fe-1.0.1.tgz
  helm install mecm-fe-edgegallery "$CHART_PREFIX"edgegallery/mecm-fe"$CHART_SUFFIX" \
    --set global.oauth2.authServerAddress=https://$EG_NODE_DEPLOY_IP:$USER_MGMT \
    --set global.ssl.enabled=true \
    --set global.ssl.secretName=edgegallery-ssl-secret
  if [ $? -eq 0 ]; then
    wait "mecm-fe" 1
    info "[Deployed MECM-FE  ..........]" $GREEN
  else
    info "[MECM-FE Deployment Failed  .]" $RED
    exit 1
  fi
}

########################################################################################
#######################################             
function install_appstore ()
{
  info "[Deploying AppStore  ........]" $BLUE
  helm install appstore-edgegallery "$CHART_PREFIX"edgegallery/appstore"$CHART_SUFFIX" \
  --set global.oauth2.authServerAddress=https://$EG_NODE_DEPLOY_IP:$USER_MGMT \
  --set global.ssl.enabled=true \
  --set global.ssl.secretName=edgegallery-ssl-secret
  if [ $? -eq 0 ]; then
    wait "appstore-be" 2
    wait "appstore-fe" 1
    info "[Deployed AppStore  .........]" $GREEN
  else
    info "[AppStore Deployment Failed  ]" $RED
    exit 1
  fi
}
########################################################################################
function install_developer ()
{
  info "[Deploying Developer  .......]"  $BLUE
  helm install developer-edgegallery "$CHART_PREFIX"edgegallery/developer"$CHART_SUFFIX" \
  --set global.oauth2.authServerAddress=https://$EG_NODE_DEPLOY_IP:$USER_MGMT \
  --set global.ssl.enabled=true \
  --set global.ssl.secretName=edgegallery-ssl-secret
  if [ $? -eq 0 ]; then
    wait "developer-be" 2
    wait "developer-fe" 1
    info "[Deployed Developer .........]" $GREEN
  else
    fail "[Developer Deployment Failed ]" $RED
    exit 1
  fi
}
######################################################################################
#***************************************************************************************
#*********************************Extra dependency required*****************************
#function password_less_ssh_check() {
#  if [[ $OFFLINE_MODE == "muno" ]]; then
#    MASTER_IP=$(echo $1 |cut -d "," -f1)
#    WORKER_LIST=`echo $2 | sed -e "s/,/ /g"`
#    for node_ip in $MASTER_IP;
#    do
#      sshpass ssh -o PubkeyAuthentication=yes  -o PasswordAuthentication=no $node_ip "exit 0"
#      if [[ $? != 0 ]]; then
#        info "Error: PasswordLess SSH is not setup among hosts" $RED
#        info "Set PasswordLess SSH by running: bash eg.sh -p master-ip,worker-ip1,..worker-ipn ROOT_PASSWORD" $RED
#        exit 1
#      fi
#    done
#    for node_ip in $WORKER_LIST;
#    do
#      sshpass ssh -o PubkeyAuthentication=yes  -o PasswordAuthentication=no $node_ip "exit 0"
#      if [[ $? != 0 ]]; then
#        info "Error: PasswordLess SSH is not setup among hosts" $RED
#        info "Set PasswordLess SSH by running: bash eg.sh -p master-ip,worker-ip1,..worker-ipn ROOT_PASSWORD" $RED
#        exit 1
#      fi
#    done
#  fi
#}

#*************************************************************************************************

function make_remote_dir() {
    MASTER_IP=$(echo $1|cut -d "," -f1)
    WORKER_LIST=`echo $2 | sed -e "s/,/ /g"`
    sshpass ssh root@$MASTER_IP "mkdir -p /tmp/remote-platform/helm"
    for node_ip in $WORKER_LIST;
    do
      sshpass ssh root@$node_ip "mkdir -p /tmp/remote-platform/helm"
    done
}

#**************************************************************************************************
#function configure_eg_ecosystem_on_remote()
#{
#  MASTER_IP=$1
#  WORKER_LIST=$2
 # _setup_insecure_registry $MASTER_IP $WORKER_LIST
 # for node_ip in $MASTER_IP;
 # do
 #     sshpass ssh root@$node_ip \
 #     "helm repo remove edgegallery stable; helm repo add edgegallery http://${PRIVATE_REGISTRY_IP}:8080/edgegallery;
 #      helm repo add stable http://${PRIVATE_REGISTRY_IP}:8080/stable" < /dev/null
 # done
#}

#****************************************************************************************************
################################################################
function _setup_helm_repo()
{
  cd "$TARBALL_PATH"/helm/helm-charts/ || exit
  helm repo index edgegallery/
  helm repo index stable/
  docker run --name helm-repo -v ~/edgeakraino/jenkins/work/workspace/ealt-edge-deploy-virtual-daily-master/ocd/infra/playbooks/roles/edgegallery_muno/files/deploy/helm/helm-charts/:/usr/share/nginx/html:ro  -d -p 8080:80  nginx:stable
  helm repo remove edgegallery stable;
  sleep 3
  helm repo add edgegallery http://${PRIVATE_REGISTRY_IP}:8080/edgegallery;
  helm repo add stable http://${PRIVATE_REGISTRY_IP}:8080/stable
}

##############################################################
############################################
function setup_eg_ecosystem()
{
    _setup_helm_repo
}
#########################################
#skip main in case of source
    main $@
######################
