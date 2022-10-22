package service

import (
	"encoding/json"
	"github.com/duxphp/duxgo-admin/tools/model"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/exception"
	"github.com/duxphp/duxgo/response"
	"github.com/duxphp/duxgo/util/function"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func AreaTree(ctx echo.Context) error {
	type Result struct {
		Label    string `json:"label"`
		Value    string `json:"value"`
		Id       string `json:"id"`
		ParentId string `json:"parent_id"`
	}
	level := cast.ToUint(ctx.QueryParam("level"))
	if level == 0 {
		level = 2
	}
	results := []Result{}
	err := core.Db.Model(&model.ToolDistrict{}).Select("name as label", "name as value", "id", "parent_id").Where("level <= ?", level).Find(&results).Error
	if err != nil {
		return exception.Internal(err)
	}
	dataJson, _ := json.Marshal(results)
	data := []map[string]any{}
	json.Unmarshal(dataJson, &data)
	return response.New(ctx).Send("ok", function.SliceToTree(data, "id", "parent_id", "children"))
}
