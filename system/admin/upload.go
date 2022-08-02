package admin

import (
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo-admin/tools/model"
	"github.com/duxphp/duxgo/core"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func Upload(ctx echo.Context) error {
	var data model.ToolFileDir
	err := core.Db.Where("has_type = ?", "admin").First(&data).Error
	if err != nil {
		return err
	}
	id := ctx.FormValue("id")
	if id == "" {
		id = cast.ToString(data.ID)
	}
	return service.Upload(ctx, "admin", cast.ToInt(id))
}
