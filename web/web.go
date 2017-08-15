package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-macaron/cache"
	"github.com/go-macaron/csrf"
	"github.com/go-macaron/session"
	"github.com/urfave/cli"
	"gopkg.in/macaron.v1"

	"xsec-dns-server/web/routers"
)

func RunWeb(ctx *cli.Context) (err error) {
	m := macaron.Classic()
	m.Use(macaron.Renderer())
	m.Use(session.Sessioner())
	m.Use(csrf.Csrfer())
	m.Use(cache.Cacher())

	m.Get("/", routers.Index)
	m.Get("/admin/index/", routers.Index)

	log.Printf("run server on %v\n", fmt.Sprintf("%v:%v", HTTP_HOST, HTTP_PORT))
	err = http.ListenAndServe(fmt.Sprintf("%v:%v", HTTP_HOST, HTTP_PORT), m)

	return err
}
