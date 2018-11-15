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
)

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
