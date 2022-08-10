# OCD for ealt-edge
It contains ansible scripts for EALTEdge stack deployment.

This contains the config.yml which helps to configure the edge gallery deployment.   
config.yml which contains variable for which user needs to provide values for deployment of EALTEdge stack.

ealt-all.yml and ealt-all-uninstall.yml which helps to install and uninstall the platform respectively.   
In general, the above path contains config files to deploy the platform.

ealt-eg-aio-latest.yml and ealt-eg-aio-unins-latest.yml which helps to install and uninstall the edge gallery in one node respectively. 

ealt-eg-muno-controller.yml which helps to install and uninstall the controller node of edge gallery when deploying in multi nodes. 

ealt-eg-muno-edge.yml which helps to install and uninstall the edge node of edge gallery when deploying in multi nodes. 

ealt-inventory.ini for user to specify node informations.

In the directory of roles, it contains the ansible playbook roles which internally provides the functionalities of EALTEdge Blueprint.
Each role corresponds to specific functionalities of EALTEdge.

In the directory of muno-config, it contains the controller node config info and edge node config info of deploying in multi nodes.
