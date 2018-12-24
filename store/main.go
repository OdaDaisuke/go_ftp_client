package store

import "github.com/OdaDaisuke/go-ftp-client/lib"

var (
	DetailViewName    = "conn_detail"
	FtpConnectionList = []lib.ConnectionJsonModel{}
	CurView           = 0
	CurConnDetailIdx  = 0
	IdxView           = 0
	InitY             = 7
)