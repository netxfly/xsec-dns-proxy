package web

import (
	"xsec-dns-server/settings"
)

var (
	HTTP_HOST string
	HTTP_PORT int
)

func init() {
	cfg := settings.Cfg
	sec := cfg.Section("WEB_SERVER")
	HTTP_HOST = sec.Key("HTTP_HOST").MustString("127.0.0.1")
	HTTP_PORT = sec.Key("HTTP_PORT").MustInt(8088)
}
