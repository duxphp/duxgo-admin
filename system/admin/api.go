package admin

import (
	"fmt"
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo-admin/system/service"
	coreService "github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo-ui/lib/form"
	"github.com/duxphp/duxgo-ui/lib/table"
	"github.com/duxphp/duxgo-ui/lib/table/column"
	"github.com/duxphp/duxgo-ui/lib/widget"
	"github.com/duxphp/duxgo/util/function"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"math/rand"
	"time"
)

func ApiList(ctx echo.Context) error {
	return coreService.NewManageExpand("/admin/system/api/ajax").SetTable(apiTable).ListPage(ctx)
}

func ApiAjax(ctx echo.Context) error {
	return coreService.NewManageExpand("/admin/system/api/ajax").SetTable(apiTable).ListData(ctx)
}

func ApiPage(ctx echo.Context) error {
	return coreService.NewManageExpand("/admin/system/api/ajax").SetForm(apiForm).FormPage(ctx)
}

func ApiSave(ctx echo.Context) error {
	return coreService.NewManageExpand("/admin/system/api/ajax").SetForm(apiForm).SetTable(apiTable).FormSave(ctx, "end")
}

func ApiDel(ctx echo.Context) error {
	msg := coreService.NewManageExpand("/admin/system/api/ajax").SetForm(apiForm).SetTable(apiTable).Del(ctx, &model.SystemApi{})
	service.InitApi()
	return msg
}

func ApiStatus(ctx echo.Context) error {
	return coreService.NewManageExpand("/admin/system/api/ajax").SetForm(apiForm).SetTable(apiTable).Status(ctx, &model.SystemApi{})
}

func apiTable(ctx echo.Context) *table.Table {
	table := table.NewTable()
	table.SetModel(&[]model.SystemApi{}, "id")
	table.SetUrl("/admin/system/api/ajax")

	table.AddAction().SetUI(widget.NewLink("添加", "/admin/system/api/add").SetButton().SetType("dialog"))

	table.AddCol("接口名称", "name").SetUI(column.NewContext())
	table.AddCol("SecretID", "secret_id").SetUI(column.NewContext())
	table.AddCol("SecretKey", "secret_key").SetUI(column.NewContext())
	table.AddCol("状态", "status").SetUI(column.NewToggle("status").SetUrl("/admin/system/api/status", map[string]any{
		"id": "id",
	}))

	links := column.NewLink()
	links.AddUrl("编辑", "/admin/system/api/edit", map[string]any{
		"id": "id",
	}).SetType("dialog")
	links.AddUrl("删除", "/admin/system/api/del", map[string]any{
		"id": "id",
	}).SetType("ajax")
	table.AddCol("操作", "").SetWidth(130).SetUI(links)

	return table
}

func apiForm(ctx echo.Context) *form.Form {

	formUI := form.NewForm()
	formUI.SetModel(&model.SystemApi{}, "id", cast.ToUint(ctx.QueryParam("id")))

	formUI.SetUrl("/admin/system/api/save?id=" + ctx.QueryParam("id"))

	formUI.AddField("接口名称", "name").SetUI(form.NewText()).SetMust(true)

	formUI.SaveBefore(func(data map[string]any, postData map[string]any, update bool, db *gorm.DB) error {
		if !update {
			data["status"] = true
			data["secretId"] = fmt.Sprintf("%06v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(100000000))
			data["secretKey"] = function.Md5(function.RandString(32))
		}
		return nil
	})

	formUI.SaveAfter(func(data map[string]any, info any, update bool, db *gorm.DB) error {
		service.InitApi()
		return nil
	})

	return formUI
}
