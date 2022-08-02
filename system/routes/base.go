package routes

import (
	"github.com/duxphp/duxgo/util"
	"github.com/labstack/echo/v4"
)

func RouteBase(router *util.RouterData) {
	router.Post("/test", func(ctx echo.Context) error {

		type Params struct {
			SetToken  string `form:"setToken"`
			Threshold string `form:"threshold"`
		}

		type Person struct {
			AccessToken string   `form:"access_token"`
			Image       string   `form:"image"`
			Operation   []Params `form:"operation"`
		}
		p := new(Person)
		ctx.Bind(p)

		//fmt.Println("ddd", p)

		return nil
	}, "测试")
}
