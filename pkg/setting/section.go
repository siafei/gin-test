package setting

import (
	"time"
)

type ServerSettingS struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize int
	MaxPageSize     int
	LogSavePath     string
	LogFileName     string
	LogFileExt      string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type DatabaseSettingS2 struct {
	TestDBType       string
	TestUserName     string
	TestPassword     string
	TestHost         string
	TestDBName       string
	TestTablePrefix  string
	TestCharset      string
	TestParseTime    bool
	TestMaxIdleConns int
	TestMaxOpenConns int
}

type RedisSettings struct {
	Username string
	Password string
	Host     string
	Port     string
	DB       int
}
