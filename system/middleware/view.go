package middleware

import (
	duxUI "github.com/duxphp/duxgo-ui"
	"github.com/duxphp/duxgo/core"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"net/http"
	"strings"
)

type login struct {
	Logo       string
	Title      string
	Name       string
	Desc       string
	Contact    string
	Background string
	Side       []string
	Foot       string
}

type ui struct {
	Css    string
	Js     string
	Name   string
	Logo   string
	Login  login
	Socket map[string]string
}

func AdminViewHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if !wantsJson(c) {
			css := duxUI.ConfigManifest["css"].([]any)
			params := ui{
				Css:  "/" + cast.ToString(css[0]),
				Js:   "/" + cast.ToString(duxUI.ConfigManifest["file"]),
				Name: core.Config["info"].GetString("info.name"),
				Logo: "/images/logo.svg",
				Login: login{
					Logo:    "/images/logo.svg",
					Title:   core.Config["info"].GetString("info.name"),
					Name:    "系统登录",
					Desc:    core.Config["info"].GetString("info.description"),
					Contact: core.Config["info"].GetString("info.copyright"),
					Side: []string{
						"/images/login-side.png",
					},
					Foot: "/images/login-foot.png",
				},
				Socket: map[string]string{
					"api": "/admin/ws",
				},
			}
			return c.Render(http.StatusOK, "admin.gohtml", params)
		}
		return next(c)
	}
}

func wantsJson(ctx echo.Context) bool {
	xr := ctx.Request().Header.Get("X-Requested-With")
	if xr != "" && strings.Index(xr, "XMLHttpRequest") != -1 {
		return true
	}
	accept := ctx.Request().Header.Get("Accept")
	if strings.Index(accept, "/json") != -1 || strings.Index(accept, "/+json") != -1 {
		return true
	}
	return false
}
