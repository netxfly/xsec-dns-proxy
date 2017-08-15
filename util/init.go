package util

import (
	"xsec-dns-server/settings"
	"log"
)

var (
	LISTEN_HOST string
	LISTEN_PORT int
	DNS_SERVER  string
	DEBUG_MODE  bool
)

func init() {
	log.SetPrefix("[xsec-dns-proxy ] ")
	cfg := settings.Cfg
	sec := cfg.Section("DNS_PROXY")
	LISTEN_HOST = sec.Key("LISTEN_HOST").MustString("")
	LISTEN_PORT = sec.Key("LISTEN_PORT").MustInt(53)
	DNS_SERVER = sec.Key("DNS_SERVER").MustString("8.8.8.8:53")
	DEBUG_MODE = sec.Key("DEBUG_MODE").MustBool(false)
}
