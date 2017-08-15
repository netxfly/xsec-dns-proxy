package models

import (
	"xsec-dns-server/settings"

	"github.com/go-xorm/xorm"

	"upper.io/db.v3"

	"log"
)

var (
	DATA_TYPE string
	DATA_NAME string
	DATA_HOST string
	DATA_PORT int
	USERNAME  string
	PASSWORD  string
	SSL_MODE  string

	Engine        *xorm.Engine
	MongodbClient db.Database
)

func init() {
	log.SetPrefix("[xsec-dns-proxy ] ")
	cfg := settings.Cfg
	sec := cfg.Section("DATA")
	DATA_TYPE = sec.Key("DATA_TYPE").MustString("sqlite")
	DATA_NAME = sec.Key("DATA_NAME").MustString("data")
	DATA_HOST = sec.Key("DATA_HOST").MustString("DATA_HOST")
	DATA_PORT = sec.Key("DATA_PORT").MustInt(3306)
	USERNAME = sec.Key("USERNAME").MustString("USERNAME")
	PASSWORD = sec.Key("PASSWORD").MustString("PASSWORD")
	SSL_MODE = sec.Key("SSL_MODE").MustString("disable")

	NewDbEngine()

	switch DATA_TYPE {
	case "mongodb":
		// MongodbClient.Ping()
	default:
		Engine.Sync2(new(DnsInfo))

	}
}
