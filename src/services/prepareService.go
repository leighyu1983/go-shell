package service

import (
	"fmt"
	"bytes"
	"os/exec"
)

func Prepare(ip string) (string) {

}

// for centos 7
func installExpect() {
	var stdOut, stdErr bytes.Buffer
	cmd := exec.Command("rpm -ivh ")
    cmd.Stdout = &stdOut
    cmd.Stderr = &stdErr

    if err := cmd.Run(); err != nil {
		fmt.Printf( "cmd exec failed: %s : %s", fmt.Sprint( err ), stdErr.String() )
    }

	return stdOut.String()
}
