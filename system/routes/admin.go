package routes

import (
	"github.com/duxphp/duxgo-admin/system/admin"
	"github.com/duxphp/duxgo/util"
	"github.com/duxphp/duxgo/websocket"
)

func RouteAdmin(router *util.RouterData) {
	router.Get("/map/weather", admin.Main, "天气")

	router.Post("/login", admin.Login, "账号登录")
	router.Get("/login/check", admin.LoginCheck, "登录检测")
	router.Get("/login/logout", admin.LoginLogout, "退出登录")

	router.Post("/register", admin.Register, "账号注册")

	router.Get("/ws", websocket.Socket.Handler("admin"), "socket服务")

}
