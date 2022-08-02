package admin

import (
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/labstack/echo/v4"
)

func FileManage(ctx echo.Context) error {
	return service.FileManage(ctx, "admin")
}
