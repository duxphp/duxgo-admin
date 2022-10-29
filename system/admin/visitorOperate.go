package admin

import (
	"encoding/json"
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo-ui/lib/form"
	"github.com/duxphp/duxgo-ui/lib/table"
	"github.com/duxphp/duxgo-ui/lib/table/column"
	"github.com/duxphp/duxgo-ui/lib/widget"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/response"
	"github.com/golang-module/carbon/v2"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func VisitorOperateList(ctx echo.Context) error {
	return service.NewManageExpand("/admin/system/visitorOperate/ajax").SetTable(visitorOperateTable).ListPage(ctx)
}

func VisitorOperateAjax(ctx echo.Context) error {
	return service.NewManageExpand("/admin/system/visitorOperate/ajax").SetTable(visitorOperateTable).ListData(ctx)
}

func VisitorOperateInfo(ctx echo.Context) error {
	id := cast.ToInt(ctx.QueryParam("id"))
	info := model.VisitorOperate{}
	err := core.Db.Model(&[]model.VisitorOperate{}).Find(&info, id).Error
	if err != nil {
		return err
	}
	data := map[string]any{}
	json.Unmarshal([]byte(info.Params), &data)
	node := map[string]any{
		"node": map[string]any{
			"nodeName": "app-dialog",
			"title":    "请求数据",
			"child": map[string]any{
				"nodeName":     "json-viewer",
				"class":        "bg-blackgray-4",
				"copyable":     "true",
				"boxed":        "true",
				"expand-depth": "5",
				"value":        data,
			},
		},
	}

	return response.New(ctx).Send("ok", node)

}

func visitorOperateTable(ctx echo.Context) *table.Table {
	table := table.NewTable()
	table.SetModel(&[]model.VisitorOperate{}, "id")
	table.ModelOrder("id desc")

	table.AddFilter("天数", "day").SetWhere(func(s string, db *gorm.DB) {
		day := cast.ToInt(s)
		stopTime := carbon.Now().ToDateString()
		startTime := carbon.Now().SubDays(day).ToDateString()
		db.Where("created_at >= ?", startTime).Where("created_at <= ?", stopTime)

	})

	table.SetUrl("/admin/system/visitorOperate/ajax")
	table.AddAction().SetUI(widget.NewLink("添加", "/admin/system/visitorOperate/add").SetButton().SetType("dialog"))

	table.AddFilter("路径", "url").SetUI(form.NewText()).SetQuick(true)

	table.AddCol("动作", "method").SetUI(column.NewContext())
	table.AddCol("路径", "url").SetUI(column.NewContext())
	table.AddCol("用户ID", "user_id").SetUI(column.NewContext())
	table.AddCol("时间", "created_at", func(val any, items map[string]any) any {
		return carbon.Parse(cast.ToString(val)).ToDateTimeString()
	}).SetUI(column.NewContext())

	links := column.NewLink()
	links.AddUrl("查看", "/admin/system/visitorOperate/info", map[string]any{
		"id": "id",
	}).SetType("dialog")
	table.AddCol("操作", "").SetWidth(130).SetUI(links)

	return table
}
