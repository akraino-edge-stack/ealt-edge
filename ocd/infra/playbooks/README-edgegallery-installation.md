# EdgeGallery Ansible Offline Installation

This Guide is for EdgeGallery (EG) offline installation when there is no public network with the environment.

The same as online installation, the offline installation is also based on Ubuntu OS and Kubernetes, supports x86_64 and ARM64 as well.

## 1. The Dependencies and How to Set Nodes

  EdgeGallery supports Multi Node and All-In-One (AIO) deployment now.

### 1.1 AIO Deployment


##  Pre-requisites

   One node with below configs

  | Module     | Version | Arch            |
  |------------|---------|-----------------|
  | Ubuntu     | 18.04   | ARM 64 & X86_64 |
  | Python     | 3.6.9   | ARM 64 & X86_64 |
  | pip3       | 9.0.1   | ARM 64 & X86_64 |
  | Ansible    | 2.10.7  | ARM 64 & X86_64 |
  | sshpass    | 1.06-1  | ARM 64 & X86_64 |

  The Master Node should only install Ubuntu 18.04 and with the following hardware resources:

  - 4CPU
  - 16G RAM
  - 100G Storage
  - Single or Multi NIC

  INFO: The Ansible controller node and the Master Node could be the same node.

  Download and install the pre-requisites mentioned above

## 2. How to Config the Ansible Controller Node

  The commands in the following sections are all executed on  **Ansible controller node**  and there is  **no commands** 
  that need to be executed on any other nodes.

### 2.1 Login Ansible controller node

  The Ansible controller node should already install ubuntu 18.04, python3.6 and pip3 in advance.

### 2.2 Install Ansible: (Can be skipped If installed)

  - Ansible Online Installation (Can be skipped If installed)

      ```
      # Recommend to install Ansible with python3
      apt install -y python3-pip
      pip3 install ansible
      ```

### Set password-less ssh from Ansible controller node to other nodes

    2.1. sshpass required：

    ```
    # Install sshpass

    # Check whether sshpass installed
    sshpass -V

    ```

    2.2 There should be id_rsa and id_rsa.pub under /root/.ssh/, if not, do the following to generate them:

    ```
    ssh-keygen -t rsa
    ```

    2.3 Do the following to set the password-less ssh, execute the command several times for all master and worker nodes
        one by one where `<master-or-worker-node-ip>` is the private IP and `<master-or-worker-node-root-password>` is
        the password of root user of that node.

    ```
    sshpass -p <master-or-worker-node-root-password> ssh-copy-id -o StrictHostKeyChecking=no root@<master-or-worker-node-ip>
    ```
  3. Set hosts-aio
  Open hosts-aio in playbook directory ealt-edge/ocd/infra/playbook and provide master node ip in place of master-ip

  - AIO Inventory, replace the exactly master node IP in file `host-aio`:

    ```
    [master]
    xxx.xxx.xxx.xxx
    ```
  - If SSH port is not the default value 22, should add some more info about the ssh port

    ```
    [master]
    xxx.xxx.xxx.xxx
    [master:vars]
    ansible_ssh_port=xx

## 3. EdgeGallery Deployment
   
   ```
   # Install edgegallery
   ansible-playbook --inventory hosts-aio ealt-eg-aio-latest.yml -e "ansible_user=root" --extra-vars "operation=install"

   ```

### 3.2. How to Set the Parameters

  All parameters that user could set are in file ealtedge/ocd/infra/playbooks/var.yml.

  ```
  # Set the Password of Harbor admin account
  HARBOR_ADMIN_PASSWORD: Harbor@edge

  # ip for portals, will be set to private IP of master node default or reset it to be the public IP of master node here
  # PORTAL_IP: xxx.xxx.xxx.xxx

  # NIC name of master node
  # If master node is with single NIC, not need to set it here and will get the default NIC name during the run time
  # If master node is with multiple NICs, should set it here to be 2 different NICs
  # EG_NODE_EDGE_MP1: eth0
  # EG_NODE_EDGE_MM5: eth0
  ```

  Note: No need to modify the above file.  But credentials can be changed in var.yml

## 5. Uninstall EdgeGallery

AIO mode

```
# Uninstall AIO Deployment
cd ealt-edge/ocd/infra/playbooks
ansible-playbook --inventory hosts-aio ealt-eg-aio-unins-latest.yml -e "ansible_user=root" --extra-vars "operation=uninstall"

```