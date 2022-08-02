package admin

import (
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo-ui/lib/form"
	"github.com/duxphp/duxgo-ui/lib/table"
	"github.com/duxphp/duxgo-ui/lib/table/column"
	"github.com/duxphp/duxgo-ui/lib/widget"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/util/function"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"strings"
)

func UserList(ctx echo.Context) error {
	return service.NewManageExpand("/admin/system/user/ajax").SetTable(userTable).ListPage(ctx)
}

func UserAjax(ctx echo.Context) error {
	return service.NewManageExpand("/admin/system/user/ajax").SetTable(userTable).ListData(ctx)
}

func UserPage(ctx echo.Context) error {
	return service.NewManageExpand("/admin/system/user/ajax").SetForm(userForm).FormPage(ctx)
}

func UserSave(ctx echo.Context) error {
	return service.NewManageExpand("/admin/system/user/ajax").SetForm(userForm).SetTable(userTable).FormSave(ctx)
}

func userTable(ctx echo.Context) *table.Table {
	table := table.NewTable()
	table.SetModel(&[]model.SystemUser{}, "id")
	table.PreloadModel("Roles")

	table.SetUrl("/admin/system/user/ajax")
	table.AddAction().SetUI(widget.NewLink("添加", "/admin/system/user/add").SetButton().SetType("dialog"))

	table.AddFilter("用户搜索", "username").SetUI(form.NewText()).SetQuick(true)

	table.AddCol("用户", "username").SetUI(column.NewContext())
	table.AddCol("角色", "roles", func(value any, items map[string]any) any {
		var names []string
		data := cast.ToSlice(value)
		for _, item := range data {
			names = append(names, cast.ToString(cast.ToStringMap(item)["name"]))
		}
		if len(names) > 0 {
			return strings.Join(names, ",")
		}
		return "-"
	}).SetUI(column.NewContext())

	links := column.NewLink()
	links.AddUrl("编辑", "/admin/system/user/edit", map[string]any{
		"id": "id",
	}).SetType("dialog")
	links.AddUrl("删除", "/admin/system/user/del", map[string]any{
		"id": "id",
	}).SetType("ajax")
	table.AddCol("操作", "").SetWidth(130).SetUI(links)

	return table
}

func userForm(ctx echo.Context) *form.Form {
	formUI := form.NewForm()
	formUI.SetModel(&model.SystemUser{}, "id", cast.ToUint(ctx.QueryParam("id")))

	formUI.SetUrl("/admin/system/user/save?id=" + ctx.QueryParam("id"))

	formUI.AddField("用户名", "username").SetUI(form.NewText())

	role := func() map[any]any {
		var results []map[string]any
		core.Db.Model(&model.SystemRole{}).Find(&results)
		return function.MapPluck(results, "name", "id")
	}
	formUI.AddField("角色", "role_id").SetUI(form.NewSelect().SetMulti().SetOptions(role())).SetHas("Roles", &[]model.SystemRole{})

	return formUI
}
