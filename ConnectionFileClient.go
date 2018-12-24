package main

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

func (f *FileClient) ReadAll() []*ConnectionJsonModel {
	return nil
}

func (f *FileClient) AddConnection(list *ConnectionJsonModel) *ConnectionJsonModel {
	return nil
}


func (f *FileClient) UpdateConnection(conn *ConnectionJsonModel) *ConnectionJsonModel {
	return nil
}


func (f *FileClient) DeleteConnection(conn *ConnectionJsonModel) {
}