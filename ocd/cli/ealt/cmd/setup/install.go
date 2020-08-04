/*
Copyright 2020 Huawei Technologies Co., Ltd.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package setup

import (
	"fmt"
)

//Function : Commands for all installation components.
//Depending on the option respective command will be executed.
func EaltInstall(component string) error {
	var strEaltSetup string
	switch component {
	case "all":
		strEaltSetup = fmt.Sprintf("ansible-playbook ealt-all.yml -i ealt-inventory.ini --extra-vars \"operation=install mode=dev\"")
	// Production Mode : SSL Mode Installation Command.
	case "sslall":
		strEaltSetup = fmt.Sprintf("ansible-playbook ealt-all.yml -i ealt-inventory.ini --extra-vars \"operation=install mode=prod\"")
	case "infra":
		strEaltSetup = fmt.Sprintf("ansible-playbook ealt-all.yml -i ealt-inventory.ini --tags \"infra\" --extra-vars \"operation=install\"")
	case "mecm":
		strEaltSetup = fmt.Sprintf("ansible-playbook ealt-all.yml -i ealt-inventory.ini --tags \"mecm\" --extra-vars \"operation=install mode=dev\"")
	case "sslmecm":
		strEaltSetup = fmt.Sprintf("ansible-playbook ealt-all.yml -i ealt-inventory.ini --tags \"mecm\" --extra-vars \"operation=install mode=prod\"")
	case "edge":
		strEaltSetup = fmt.Sprintf("ansible-playbook ealt-all.yml -i ealt-inventory.ini --tags \"mep\" --extra-vars \"operation=install mode=dev\"")
	case "ssledge":
		strEaltSetup = fmt.Sprintf("ansible-playbook ealt-all.yml -i ealt-inventory.ini --tags \"mep\" --extra-vars \"operation=install mode=prod\"")
	case "k8s":
		strEaltSetup = fmt.Sprintf("ansible-playbook ealt-all.yml -i ealt-inventory.ini --tags \"k8s\" --extra-vars \"operation=install\"")
	case "k3s":
		strEaltSetup = fmt.Sprintf("ansible-playbook ealt-all.yml -i ealt-inventory.ini --tags \"k3s\" --extra-vars \"operation=install\"")
	default:
		fmt.Println("Provide subcommand for ealt init [all|infra|manager|edge|k8s|k3s]")
	}

	stdout, err := runCommandAtShell(strEaltSetup)
	if err != nil {
		return err
	}
	fmt.Println(stdout)
	return nil
}

func EaltReset(component string) error {
	var strEaltReset string

	switch component {
	case "all":
		strEaltReset = fmt.Sprintf("ansible-playbook ealt-all-uninstall.yml -i ealt-inventory.ini --extra-vars \"operation=uninstall\"")
	case "infra":
		strEaltReset = fmt.Sprintf("ansible-playbook ealt-all-uninstall.yml -i ealt-inventory.ini --tags \"infra\" --extra-vars \"operation=uninstall\"")
	case "manager":
		strEaltReset = fmt.Sprintf("ansible-playbook ealt-all-uninstall.yml -i ealt-inventory.ini --tags \"mecm\" --extra-vars \"operation=uninstall\"")
	case "edge":
		strEaltReset = fmt.Sprintf("ansible-playbook ealt-all-uninstall.yml -i ealt-inventory.ini --tags \"mep\" --extra-vars \"operation=uninstall\"")
	case "k8s":
		strEaltReset = fmt.Sprintf("cd ~/kubespray && yes | ansible-playbook -i inventory/mycluster/hosts.yaml --user root reset.yml")
	case "k3s":
		strEaltReset = fmt.Sprintf("ansible-playbook ealt-all-uninstall.yml -i ealt-inventory.ini --tags \"k3s\" --extra-vars \"operation=uninstall\"")
	default:
		fmt.Println("Provide subcommand for ealt clean [all|infra|manager|edge|k8s|k3s]")
	}

	stdout, err := runCommandAtShell(strEaltReset)
	if err != nil {
		return err
	}
	fmt.Println(stdout)
	return nil
}
