package main

import (
	"os"
	"runtime"

	"xsec-dns-server/cmd"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "xsec dns proxy server"
	app.Usage = "xsec dns proxy server"
	app.Version = "0.1"
	app.Commands = []cli.Command{
		cmd.Serve,
		cmd.RunWeb,
	}
	app.Run(os.Args)
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}
