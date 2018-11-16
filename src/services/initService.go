package service

import (
	"fmt"
	"utils"
)

func PrepareEnv(ip string) (infoMsg string) {
	infoMsg = installExpect() + "\n"
	infoMsg += generateSSHFile()
	return infoMsg;
}

// Install auto input shell plugin for centos 7
func installExpect() (infoMsg string) {
	entity, _ := util.GetConfig()
	command := fmt.Sprintf("%s/scripts/rpms/expect/*.rpm", entity.GoPath)
	infoMsg, err := ExecCommand("rpm", "-ivh", command)
	if(err != nil) {
		panic(err)
	}
	return infoMsg
}

// Generate ssh file without password
func generateSSHFile() (infoMsg string) {
	infoMsg, err := ExecCommand("ssh-keygen", "-t", "rsa", "-P", "", "-f", "~/.ssh/id_rsa")
	if(err != nil) {
		panic(err)
	}
	return infoMsg
}