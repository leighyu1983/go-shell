package service

import (
	"fmt"
	"bytes"
	"strconv"
	"strings"
)


func TestSetHostname(ip string, hostname string) (string, error) {
	session, err := SSHConnect( "root", "123456", ip, 22 )
    if err != nil {
        panic(err)
    }
	defer session.Close()

	var stdOut, stdErr bytes.Buffer
    session.Stdout = &stdOut
    session.Stderr = &stdErr

	session.Run("hostnamectl set-hostname " + hostname)
	msg := strings.Replace(stdOut.String(), "\n", "", -1 )
    ret, err := strconv.Atoi(msg)
	fmt.Printf("=== finish setting %d, %s\n", ret, stdErr.String())

	return msg, err
}


func GenerateSshFile(ip string, command string) (string, error) {
	session, err := SSHConnect( "root", "123456", ip, 22 )
    if err != nil {
        panic(err)
    }
	defer session.Close()

	var stdOut, stdErr bytes.Buffer
    session.Stdout = &stdOut
    session.Stderr = &stdErr

	session.Run(command)
	msg := strings.Replace(stdOut.String(), "\n", "", -1 )
    ret, err := strconv.Atoi(msg)
	fmt.Printf("=== finish setting %d, %s\n", ret, stdErr.String())

	return msg, err
}