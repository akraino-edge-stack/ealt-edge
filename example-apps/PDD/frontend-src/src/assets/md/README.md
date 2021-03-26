# Edge LightWeight IoT
<!-- [![Build Status](https://travis-ci.org/kubeedge/kubeedge.svg?branch=master)](https://travis-ci.org/kubeedge/kubeedge) -->
[eliot-gerrit-link](https://gerrit.akraino.org/r/gitweb?p=eliot.git;a=tree)

ELIOT is "Enterprise Edge Lightweight and IOT" project under Akraino approved blueprint family and part of Akraino Edge Stack, which intends to provide an platform  for a fully integrated edge network infrastructure and running edge computing applications, on lightweight Edge Nodes. ELIOT targets on making the edge node a lightweight software stack which can be deployed on edge nodes which have limited hardware capacity by leveraging lightweight OS, a container running environment and container orchestration applications.

In addition, ELIOT stack focuses to have an infrastructure for edge computing which will enable high performance, high availability, security and reduce latency.

IoT Gateway Blueprint will provide an infrastructure platform for an edge node to provide features of an IoT Gateway.  This is achieved by leveraging the upstream opensource application like Dockers, Kubernetes , Prometheus , cAdvisor, KubeEdge and EdgeX.  Companies having applications to be deployed on edge nodes related to IoT application , can use the IoT Gateway platform to have a functional end to end to ecosystem of IoT.

## Advantages

#### Where on the Edge

IoT gateway

Smart cities
Smart homes
Connected vehicles
Connected farming, agriculture
Logistics
Industrial, IIoT

#### Overall Architecture

ELIOT IoT Gateway Blueprint  Architecture consists of an ELIOT Manager and multiple ELIOT Nodes (edge nodes). The scale of ELIOT (edge) Nodes can range from 1 single node to 10, 100, 1K or more. For development environment we have tested with one ELIOT Manager Node and two ELIOT Worker Nodes.

ELIOT Manager is the central controller which manages the orchestration, life cycle, management and networking of edge nodes (IoT Gateway). It can be installed on a bare metal server or cloud VM with Ubuntu 16.04 /18.04 or Cent OS 7.5 version (currently supported OS in IoT Gateway Blueprint).

The node ELIOT Manager  is to facilitate auto deploy, high availability, orchestration, self-healing of the IoT Gateway Nodes.

ELIOT Nodes - IoT Gateway are Edge compute processing nodes which will provide the software stack for deploying IoT applications.  

The IoT Gateway can function as a Kubernetes worker node or a KubeEdge - Edge node when used as Kubernetes or KubeEdge respectively.

As the edge nodes can have restricted hardware resources of CPU , Memory etc, ELIOT uses ELIOT Minimal OS for the Edge nodes which is basically tailored version of Cent OS.  ELIOT Minimal OS is a lightweight OS suitable for ELIOT Edge nodes.

ELIOT blueprint will use lightweight OS, Lightweight Kubelet -- KubeEdge , lightweight container runtime interface -- like containerd. By this we can leverage the core features of OS, K8S, CRI and make the edge node lightweight and efficient edge computing node by executing edge applications on node and reducing the interaction between Central Controller (Cloud Node) and the Edge Node devices.

Many cloud native monitoring applications can be used to collect container matrix and show them in a graphical manner, like cadvisor , prometheus.

## Platform Architecture

#### ELIOT Manager

#### ELIOT Portal

ELIOT Portal is new feature added on the ELIOT Blueprint which is currently applicable for both IoT Gateway and uCPE Blueprint. ELIOT Portal main objective is  to have a dashboard for operations and management of ELIOT Cluster. It consists of ELIOT UI and ELIOT API Gateway.

In  Release 2.0 the portal is giving the following operations:-

Display ELIOT Nodes
Display ELIOT Services
Display ELIOT PODS. 
Other options like Application Deployments,   are currently in design phase, which will be developed for future releases.

#### ELIOT CLI: 

ELIOT CLI (Command Line Interface) is a new feature added in ELIOT Blueprint which is currently applicable only for IoT Gateway Blueprint.  It is providing a command line interface option to setup the ELIOT Cluster. Currently this feature is supported only for Ubuntu OS.

The ELIOT CLI application currently provides the following operations:

- ELIOT Cluster Setup.
- ELIOT Cluster Cleanup.
- ELIOT Cluster Reset.

#### Node Monitoring: 

Node Monitoring main task is to collect the container performance metrics from various IoT Gateway Nodes and display on a dashboard in graphical representation. 

#### Network Controller & Node Resource Manager: 

The core part of the ELIOT Cluster which controls the Node health, high availability , overlay networking , cluster management. To manage the same Kubernetes and KubeEdge are being used for Node Resource Management with Calico for Container Networking.

## Software Platform Architecture

#### APIs

Not Applicable

#### Hardware and Software Management


Currently in IoT Gateway Blueprint virtual machines are being used in public cloud. Currently no dedicated bare metal are being used.

Software Management:-

ELIOT gerrit - https://gerrit.akraino.org/r/#/admin/projects/eliot

The core part of the ELIOT Cluster which controls the Node health, high availability , overlay networking , cluster management. To manage the same Kubernetes and KubeEdge are being used for Node Resource Management with Calico for Container Networking.


## Documentation

The detailed documentation for ELIOT and its modules can be found at [https://wiki.akraino.org/display/AK/ELIOT%3A+Edge+Lightweight+and+IoT+Blueprint+Family](ELIOT wiki). 


## Contributing

If you're interested in being a contributor and want to get involved in
developing the ELIOT code, please see [CONTRIBUTING](CONTRIBUTING.md) for
details on submitting patches and the contribution workflow.

## License

ELIOT is under the Apache 2.0 license. See the [LICENSE](LICENSE) file for details.