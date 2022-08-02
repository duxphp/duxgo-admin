package api

import (
	"github.com/duxphp/duxgo-admin/tools/model"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/response"
	"github.com/duxphp/duxgo/util"
	"github.com/labstack/echo/v4"
)

func Area(ctx echo.Context) error {
	var err error
	var params struct {
		Province string `json:"province"`
		City     string `json:"city"`
		Region   string `json:"region"`
	}
	if err = util.RequestParser(ctx, &params); err != nil {
		return err
	}

	level := -1
	name := ""
	if params.Province != "" {
		level = 0
		name = params.Province
	}
	if params.City != "" {
		level = 1
		name = params.City
	}
	if params.Region != "" {
		level = 2
		name = params.Region
	}

	var pid = 0
	if level >= 0 {
		var info model.ToolDistrict
		err = core.Db.Model(&model.ToolDistrict{}).Where("level = ?", level).Where("name = ?", name).Find(&info).Error
		if err != nil {
			return err
		}
		pid = info.ID
	}
	data := []model.ToolDistrict{}
	err = core.Db.Model(&model.ToolDistrict{}).Where("parent_id = ?", pid).Find(&data).Error
	if err != nil {
		return err
	}

	return response.New(ctx).Send("ok", data)
}
