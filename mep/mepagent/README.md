# mep_agent_lib
Agent lib for mep service regsitration API 


## Introduction

MEP agent library and sample application is provide for quickly develop applications on MEP platform.
Code is devided in 2 part.
- SampleApp
- Service registration pkg.
- config file for setting application configuration 
- docker file , build and k8s yaml files

* SampleApp 
	- sample application for service regsitration to MEP
	- contains main function and call service registration function from provided pkg
	- It has config file for setting application configuration  

* pkg for service registration
	- pkg can be used to register developer applications to MEP. 
	- support http and https connection to MEP
	- configuration can be enabled/disbaled in config files when start mep agent
	
## Configuration
mainly below configuration supported
	- MEP GW details
		- IP: IP of MEP Gateway 
		- HTTPS port:  GW HTTPS proxy port 
		- HTTP port:  GW HTTP proxy port
	- App instance ID 
	- service registration sample data as per ETSI mp1 interface.
	
## Usages
Developer who develp applications for MEP, can leverage sample application and pkg freamework to support mp1 interface for service registration.
In future this library can be extened to support all mp1 interface like discovery, service avaiibilty.
MEP support mp1 interfaces as per ETSI compliant.
	
* Steps
	- configure MEP GW IP and port in path meagent/SampleApp/conf/app_instance_info.yaml
		- based on deplyment mode(development/production) provide HTTP/HTTPS port
		- kong API GW run as K8s service, check corresponding port and config accordingly
		- Kong has admin and proxy port. 
		- use proxy port and configure
	- Build go applicaion with below cmd:
		- cd mepagent/SampleApp
		- CGO_ENABLED=0 go build -a -ldflags '-extldflags "-static"' src/main/main.go
	- Build docker and push to docker
		- sudo ./docker-build.sh
	- Deploy Application
		- In mep-k8s.yaml
		- enable/disable ssl which is env. variable in k8s yaml file
		- APP_SSL_MODE "1" to enable ssl.
		- By default app run in normal mode.
		- generate k8s secret with ca.crt file which is root CA used by MEP. 
		- MEP provide cert-manager and vault to automate it. Plz refer corresponding document.
		

