package settings

import (
	"gopkg.in/ini.v1"
	"log"
)

var (
	Cfg *ini.File
)

func init() {
	log.SetPrefix("[xsec-dns-proxy ] ")
	var err error
	source := "conf/app.ini"
	Cfg, err = ini.Load(source)

	if err != nil {
		log.Panicln(err)
	}
}
