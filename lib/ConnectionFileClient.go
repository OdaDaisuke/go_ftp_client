package lib

import (
	"github.com/OdaDaisuke/go-ftp-client/configs"
	"io/ioutil"
	"github.com/labstack/gommon/log"
	"encoding/json"
)

type FileClient struct {
}

type FileClientActions interface {
	ReadAll() []*ConnectionJsonModel
	AddConnection(list *ConnectionJsonModel) *ConnectionJsonModel
	UpdateConnection(conn *ConnectionJsonModel) *ConnectionJsonModel
	DeleteConnection(conn *ConnectionJsonModel)
}

type ConnectionJsonModel struct {
	Name      string `json:"name"`
	Host      string `json:"host"`
	IpAddress string `json:"ip_address"`
	Port      int16  `json:port`
	User      string `json:user`
	Password  string `json:password`
}

func NewFileClient() *FileClient {
	return &FileClient{}
}

func (f *FileClient) ReadAll() []ConnectionJsonModel {
	bytes, err := ioutil.ReadFile(configs.CONNECTIONS_FILE_NAME)
	if err != nil {
		log.Fatal(err)
	}
	var connList []ConnectionJsonModel
	if err := json.Unmarshal(bytes, &connList); err != nil {
		log.Fatal(err)
	}
	return connList
}

func (f *FileClient) AddConnection(list *ConnectionJsonModel) *ConnectionJsonModel {
	return nil
}


func (f *FileClient) UpdateConnection(conn *ConnectionJsonModel) *ConnectionJsonModel {
	return nil
}

func (f *FileClient) DeleteConnection(conn *ConnectionJsonModel) {
}