package admin

import (
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo-ui/lib/form"
	"github.com/duxphp/duxgo-ui/lib/table"
	"github.com/duxphp/duxgo-ui/lib/table/column"
	"github.com/duxphp/duxgo-ui/lib/widget"
	coreRegister "github.com/duxphp/duxgo/register"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"strings"
)

func RoleList(ctx echo.Context) error {
	return service.NewManageExpand("/admin/system/role/ajax").SetTable(roleTable).ListPage(ctx)
}

func RoleAjax(ctx echo.Context) error {
	return service.NewManageExpand("/admin/system/role/ajax").SetTable(roleTable).ListData(ctx)
}

func RolePage(ctx echo.Context) error {
	//return ctx.JSON(register.RouteAdminAuth.ParseTree("/admin"))
	return service.NewManageExpand("/admin/system/role/ajax").SetForm(roleForm).FormPage(ctx)
}

func RoleSave(ctx echo.Context) error {
	return service.NewManageExpand("/admin/system/role/ajax").SetForm(roleForm).SetTable(roleTable).FormSave(ctx)
}

func roleTable(ctx echo.Context) *table.Table {
	table := table.NewTable()
	table.SetModel(&[]model.SystemRole{}, "id")
	table.SetUrl("/admin/system/role/ajax")

	table.AddAction().SetUI(widget.NewLink("添加", "/admin/system/role/add").SetButton().SetType("dialog"))

	table.AddCol("角色", "name").SetUI(column.NewContext())

	links := column.NewLink()
	links.AddUrl("编辑", "/admin/system/role/edit", map[string]any{
		"id": "id",
	}).SetType("dialog")
	links.AddUrl("删除", "/admin/system/role/del", map[string]any{
		"id": "id",
	}).SetType("ajax")
	table.AddCol("操作", "").SetWidth(130).SetUI(links)

	return table
}

func roleForm(ctx echo.Context) *form.Form {

	formUI := form.NewForm()
	formUI.SetModel(&model.SystemRole{}, "id", cast.ToUint(ctx.QueryParam("id")))

	formUI.SetUrl("/system/role/save?id=" + ctx.QueryParam("id"))

	formUI.AddField("名称", "name").SetUI(form.NewText()).SetMust(true)
	authTree := coreRegister.AppRouter["adminAuth"].ParseTree("/admin")
	var treeData []map[string]any
	for appIndex, app := range authTree.([]any) {
		// 模块
		appData := cast.ToStringMap(app)
		if appData["permission"] != true {
			continue
		}
		appAuth := map[string]any{
			"id":   "app_" + cast.ToString(appIndex),
			"name": appData["name"],
		}
		var groupList []map[string]any
		for groupIndex, group := range appData["data"].([]any) {
			// 分组
			groupData := cast.ToStringMap(group)
			groupItem := map[string]any{
				"id":   "group_" + cast.ToString(appIndex) + "_" + cast.ToString(groupIndex),
				"name": groupData["name"],
			}
			var itemList []map[string]any
			for _, item := range groupData["data"].([]any) {
				// 功能
				itemData := cast.ToStringMap(item)
				itemAuth := map[string]any{
					"id":   itemData["path"],
					"name": itemData["name"],
				}
				itemList = append(itemList, itemAuth)
			}
			groupItem["children"] = itemList
			groupList = append(groupList, groupItem)
		}
		appAuth["children"] = groupList
		treeData = append(treeData, appAuth)
	}
	formUI.AddField("权限", "permissions").SetUI(form.NewTree().SetData(treeData)).SetHelp("全部不选择为拥有所有权限").SaveFormat(func(value any) any {
		data := cast.ToSlice(value)
		var newData []any
		for _, item := range data {
			v := cast.ToString(item)
			if strings.HasPrefix(v, "app_") || strings.HasPrefix(v, "group_") {
				continue
			}
			newData = append(newData, v)
		}
		return newData
	})

	return formUI
}
