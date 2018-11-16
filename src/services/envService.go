package service

import (
	"bytes"
)

func Terminal(ip string, command string) (string) {
	session, err := SSHConnect( "root", "123456", ip, 22 )
    if err != nil {
        panic(err)
    }
	defer session.Close()

	var stdOut, stdErr bytes.Buffer
    session.Stdout = &stdOut
    session.Stderr = &stdErr

	session.Run(command)
	return stdOut.String()
}

