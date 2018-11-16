package service

import (
	
)

func PrepareEnv(ip string) (infoMsg string) {
	infoMsg = installExpect() + "\n"
	infoMsg += generateSSHFile()
	return infoMsg;
}

// Install auto input shell plugin for centos 7
func installExpect() (infoMsg string) {
	infoMsg, err := ExecCommand("rpm -ivh ./rpms/expect/*")
	if(err != nil) {
		panic(err)
	}
	return infoMsg
}

// Generate ssh file without password
func generateSSHFile() (infoMsg string) {
	infoMsg, err := ExecCommand("ssh-keygen -t rsa -P '' -f ~/.ssh/id_rsa")
	if(err != nil) {
		panic(err)
	}
	return infoMsg
}