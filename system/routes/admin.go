package routes

import (
	"github.com/duxphp/duxgo-admin/system/admin"
	"github.com/duxphp/duxgo/util"
	"github.com/duxphp/duxgo/websocket"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func RouteAdmin(router *util.RouterData) {
	router.Get("/map/weather", admin.Main, "天气")

	router.Post("/login", admin.Login, "账号登录")
	router.Get("/login/check", admin.LoginCheck, "登录检测")
	router.Get("/login/logout", admin.LoginLogout, "退出登录")

	router.Post("/register", admin.Register, "账号注册")

	router.Get("/ws", func(ctx echo.Context) error {
		id := cast.ToString(ctx.Get("authID"))
		return websocket.Socket.Handler("admin", id)(ctx)
	}, "socket服务")

}
