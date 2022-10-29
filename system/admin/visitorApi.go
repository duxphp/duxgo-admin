package admin

import (
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo-ui/lib/form"
	"github.com/duxphp/duxgo-ui/lib/table"
	"github.com/duxphp/duxgo-ui/lib/table/column"
	"github.com/duxphp/duxgo-ui/lib/widget"
	"github.com/golang-module/carbon/v2"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func VisitorApiList(ctx echo.Context) error {
	return service.NewManageExpand("/admin/system/visitorApi/ajax").SetTable(visitorApiTable).ListPage(ctx)
}

func VisitorApiAjax(ctx echo.Context) error {
	return service.NewManageExpand("/admin/system/visitorApi/ajax").SetTable(visitorApiTable).ListData(ctx)
}

func visitorApiTable(ctx echo.Context) *table.Table {
	table := table.NewTable()
	table.SetModel(&[]model.VisitorApi{}, "id")

	table.SetUrl("/admin/system/visitorApi/ajax")
	table.AddAction().SetUI(widget.NewLink("添加", "/admin/system/visitorApi/add").SetButton().SetType("dialog"))

	table.AddFilter("用户搜索", "username").SetUI(form.NewText()).SetQuick(true)

	table.AddFilter("天数", "day").SetWhere(func(s string, db *gorm.DB) {
		day := cast.ToInt(s)
		stopTime := carbon.Now().ToDateString()
		startTime := carbon.Now().SubDays(day).ToDateString()
		db.Where("date >= ?", startTime).Where("date <= ?", stopTime)
	})

	table.AddCol("动作", "method").SetUI(column.NewContext())
	table.AddCol("路径", "url").SetUI(column.NewContext())
	table.AddCol("路径", "uv").SetUI(column.NewContext()).SetSort()
	table.AddCol("路径", "pv").SetUI(column.NewContext()).SetSort()
	table.AddCol("最大延迟", "max_time").SetUI(column.NewContext()).SetSort()
	table.AddCol("最小延迟", "min_time").SetUI(column.NewContext()).SetSort()
	table.AddCol("日期", "date", func(val any, items map[string]any) any {
		return carbon.Parse(cast.ToString(val)).ToDateString()
	}).SetUI(column.NewContext())

	return table
}
