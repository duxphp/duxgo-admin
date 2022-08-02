package admin

import (
	"github.com/duxphp/duxgo/response"
	"github.com/labstack/echo/v4"
)

func ControlHistory(ctx echo.Context) error {
	assign := map[string]any{}
	return response.New(ctx).Render("adminControlHistory.gohtml", assign)
}
