package routers

import (
	"gopkg.in/macaron.v1"

	"xsec-dns-server/models"
)

func Index(ctx *macaron.Context) {
	switch models.DATA_TYPE {
	case "mongodb":
		info, _ := models.MgoQuery()
		ctx.Data["info"] = info
		ctx.HTML(200, "index_m")
	default:
		info, _ := models.Query()
		ctx.Data["info"] = info
		ctx.HTML(200, "index")
	}
}
