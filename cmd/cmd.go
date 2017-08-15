package cmd

import (
	"github.com/urfave/cli"

	"xsec-dns-server/web"
	"xsec-dns-server/util"
)

var Serve = cli.Command{
	Name:        "serve",
	Usage:       "dns proxy Server",
	Description: "Start dns proxy server",
	Action:      util.Run,
}

var RunWeb = cli.Command{
	Name:        "web",
	Usage:       "web server",
	Description: "dns log web server",
	Action:      web.RunWeb,
}
