package routes

import (
	"github.com/duxphp/duxgo/response"
	"github.com/duxphp/duxgo/util"
	"github.com/labstack/echo/v4"
)

func RouteWeb(router *util.RouterData) {
	router.Get("", func(ctx echo.Context) error {
		return response.New(ctx).Send("ok")
	}, "home")
}
