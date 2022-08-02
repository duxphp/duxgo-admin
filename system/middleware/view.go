package middleware

import (
	"fmt"
	duxUI "github.com/duxphp/duxgo-ui"
	"github.com/duxphp/duxgo/config"
	"github.com/duxphp/duxgo/util/function"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"net/http"
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
		if !function.IsAjax(c) {
			css := duxUI.ConfigManifest["css"].([]any)
			params := ui{
				Css:  "/" + cast.ToString(css[0]),
				Js:   "/" + cast.ToString(duxUI.ConfigManifest["file"]),
				Name: config.Get("info").GetString("info.name"),
				Logo: "/images/logo.svg",
				Login: login{
					Logo:    "/images/logo.svg",
					Title:   config.Get("info").GetString("info.name"),
					Name:    "系统登录",
					Desc:    config.Get("info").GetString("info.description"),
					Contact: config.Get("info").GetString("info.copyright"),
					Side: []string{
						"/images/login-side.png",
					},
					Foot: "/images/login-foot.png",
				},
				Socket: map[string]string{
					"api": "/admin/ws",
				},
			}

			fmt.Println("渲染页面")
			return c.Render(http.StatusOK, "admin.gohtml", params)
		}
		return next(c)
	}
}
