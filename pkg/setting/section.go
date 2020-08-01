package setting

import (
	"time"
)

type ServerSetting struct{
	RunMode string
    HttpPort int
    ReadTimeout time.Duration
    WriterTimeout time.Duration
}

type AppSetting struct{
	DefaultPageSize int
	MaxPageSize int
	LogSavePath string
	LogFileName string
	LogFileExt string
}

type DatabaseSetting struct{
	DBType string
    Username string 
    Password string
    Host string
    DBName string
    TablePrefix string
    Charset string
    ParseTime bool
    MaxIdleConns int
    MaxOpenConns int
}
