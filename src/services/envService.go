package service

import (
	"fmt"
	"bytes"
	"strconv"
	str "strings"
)

func SetHostname(ip string, hostname string) (string, error) {
	session, err := SSHConnect( "username", "passworld", ip, 22 )
    if err != nil {
        panic(err)
    }
	defer session.Close()

	var stdOut, stdErr bytes.Buffer
    session.Stdout = &stdOut
    session.Stderr = &stdErr

	session.Run("hostnamectl set-hostname " + hostname)
	msg := str.Replace(stdOut.String(), "\n", "", -1 )
    ret, err := strconv.Atoi(msg)
	fmt.Printf("%d, %s\n", ret, stdErr.String())

	return msg, err
}