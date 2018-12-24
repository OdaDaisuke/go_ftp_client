package store

import "github.com/OdaDaisuke/go-ftp-client/lib"

var (
	FtpConnectionList = []lib.ConnectionJsonModel{}
	CurView           = -1
	IdxView           = 0
	InitY             = 7
)