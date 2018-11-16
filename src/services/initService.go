package service

import (
	"fmt"
	"utils"
)

func PrepareEnv() (infoMsg string) {
	infoMsg = installExpect() + "\n"
	infoMsg += generateSshFile()
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
func generateSshFile() (infoMsg string) {
	// doesn't work via exec.command, have to use ssh session
	//infoMsg, err := ExecCommand("ssh-keygen", "-t", "rsa", "-P", "''", "-f", "~/.ssh/id_rsa")
	msg, err := ExecSshSessionCommand("192.168.121.11", "ssh-keygen -t rsa -P '' -f ~/.ssh/id_rsa");
	if(err != nil) {
		panic(err)
	}
	return msg
}

