package service

import (
	"fmt"
)

// copy ssh key to target machine which is specified by parameter ip
func SyncSshKey(ip string, password string) (string) {
	command := fmt.Sprintf("autocopy.exp  root@%s %s", ip, password);
	msg, _:= ExecCommand(command)
	
	return msg;
}

func SetHostname(ip string, hostname string) (string) {
	command := fmt.Sprintf("ssh %s hostnamectl set-hostname %s", ip, hostname);
	msg, _:= ExecCommand(command)
	
	return msg;
}

func DisbleFirewall(ip string) (string) {
	commands := []string{
		"ssh %s systemctl stop firewalld", 
		"ssh %s systemctl disable firewalld"}
	
	var (
		msgs string 
	)

	for _, commandStr := range commands {
		command := fmt.Sprintf(commandStr, ip);
		msg, _:= ExecCommand(command)
		msgs += "\n" + msg
	}

	return msgs;
}


func DisbleSelinux(ip string) (string) {
	commands := []string{
		"ssh %s setenforce 0", 
		"ssh %s sed -i s/'SELINUX=enforcing'/'SELINUX=disabled'/g  /etc/selinux/config"}
	
	var (
		msgs string 
	)

	for _, commandStr := range commands {
		command := fmt.Sprintf(commandStr, ip);
		msg, _:= ExecCommand(command)
		msgs += "\n" + msg
	}

	return msgs;
}

func SetTimezone(ip string) (string) {
	command := fmt.Sprintf(
		"ssh %s \\cp -rf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime", ip);
	msg, _:= ExecCommand(command)
	
	return msg;
}