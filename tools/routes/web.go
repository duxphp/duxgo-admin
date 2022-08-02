package routes

import "C"
import (
	"github.com/duxphp/duxgo/util"
	"github.com/labstack/echo/v4"
)

func RouteWeb(router *util.RouterData) {

	router.Get("tools/test", func(ctx echo.Context) error {

		//coll := core.Mgo.Collection("notice_high")
		//
		//_, _ = coll.RemoveAll(core.Ctx, bson.M{"guard": "guest"})

		return nil

	}, "地区数据")
}
