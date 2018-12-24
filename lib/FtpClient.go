package lib

import (
	"github.com/jlaffaye/ftp"
	"fmt"
)

type FTPClient struct {
	serverConn *(ftp.ServerConn)
	host string
	user string
	password string
	port int16
}

func NewFTPClient(
	host string,
	user string,
	password string,
	port int16,
) *FTPClient {
	return &FTPClient{
		serverConn: &(ftp.ServerConn{}),
		host: host,
		user: user,
		password: password,
		port: port,
	}
}

func (f *FTPClient) SetConf(
	host string,
	user string,
	password string,
	port int16,
) {
	f.host = host
	f.user = user
	f.password = password
	f.port = port
}

func (f *FTPClient) Connect() error {
	var address string = fmt.Sprintf("%s:%d", f.host, f.port)
	client, err := ftp.Dial(address)
	if err != nil {
		return err
	}
	defer client.Quit()
	f.serverConn = client

	err = f.serverConn.Login(f.user, f.password)
	if err != nil {
		return err
	}
	return nil
}

func (f *FTPClient) DisConnect() {
	err := f.serverConn.Logout()
	if err != nil {
		panic(err)
	}
}

func (f *FTPClient) Exec(cmd string) {}