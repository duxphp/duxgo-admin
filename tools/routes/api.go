package routes

import (
	"github.com/duxphp/duxgo-admin/tools/api"
	"github.com/duxphp/duxgo/util"
)

func RouteApi(router *util.RouterData) {
	router.Post("/tools/area", api.Area, "地区数据")
}
