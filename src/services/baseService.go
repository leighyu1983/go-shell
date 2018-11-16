package service

import (
	/*
	    mkdir -p $GOPATH/src/golang.org/x/
        cd $GOPATH/src/golang.org/x/
        git clone https://github.com/golang/crypto.git
	*/
	"golang.org/x/crypto/ssh"
	"net"
    "fmt"
    "strings"
    "bytes"
	"os/exec"
	//"errors"
)

/*
  A
  error 猜测来自于获取上一步的返回值，$? - catch  the status of last command
  0: Success 成功
  1: Failure 错误需要分析
  2: Incorrect Usage 用法不当 
  126: Not an executable 不是可执行的
  127: Command Not Found 命令没有找到


  B
  exec包的exec.Command的第一个参数command，都源码可知，会去linux /usr/bin目录下搜索是否有这个命令并且是可以执行的。
  ssh 包的session会直接运行语句(不是cat一类的命令)

*/
func ExecCommand(command string, args ...string ) (infoMsg string, err error)  {
	var stdOut, stdErr bytes.Buffer
	cmd := exec.Command(command, args...)
    cmd.Stdout = &stdOut
    cmd.Stderr = &stdErr

    // 错误2 可以忽略，比如rpm包已经存在，重复安装
    if err := cmd.Run(); (err != nil && !strings.ContainsAny(stdErr.String(), "exit status 2")) {
		errMsg := fmt.Sprintf("cmd exec failed: %s : %s", fmt.Sprint( err ), stdErr.String())
        fmt.Println(errMsg)
        //return "", errors.New(errMsg)
        panic(err)
    }

	return stdOut.String(), nil
}

func SSHConnect( user, password, host string, port int ) ( *ssh.Session, error ) {
    var (
        auth         []ssh.AuthMethod
        addr         string
        clientConfig *ssh.ClientConfig
        client       *ssh.Client
        session      *ssh.Session
        err          error
    )
    // get auth method
    auth = make([]ssh.AuthMethod, 0)
    auth = append(auth, ssh.Password(password))

    hostKeyCallbk := func(hostname string, remote net.Addr, key ssh.PublicKey) error {
            return nil
    }

    clientConfig = &ssh.ClientConfig{
        User:               user,
        Auth:               auth,
        // Timeout:             30 * time.Second,
        HostKeyCallback:    hostKeyCallbk, 
    }

    // connet to ssh
    addr = fmt.Sprintf( "%s:%d", host, port )

    if client, err = ssh.Dial( "tcp", addr, clientConfig ); err != nil {
        return nil, err
    }

    // create session
    if session, err = client.NewSession(); err != nil {
        return nil, err
    }

    return session, nil
}
