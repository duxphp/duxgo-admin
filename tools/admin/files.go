package admin

import (
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo-admin/tools/model"
	"github.com/duxphp/duxgo-ui/lib/node"
	tableUI "github.com/duxphp/duxgo-ui/lib/table"
	"github.com/duxphp/duxgo-ui/lib/table/column"
	"github.com/duxphp/duxgo-ui/lib/widget"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/pkg"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"time"
)

func FilesList(ctx echo.Context) error {
	return service.NewManageExpand("/admin/tools/files/ajax").SetTable(FilesTable).ListPage(ctx)
}

func FilesAjax(ctx echo.Context) error {
	return service.NewManageExpand("/admin/tools/files/ajax").SetTable(FilesTable).ListData(ctx)
}

func FilesDel(ctx echo.Context) error {
	return service.NewManageExpand("/admin/tools/files/ajax").SetTable(FilesTable).CallDel(func(id string, tx *gorm.DB) error {
		var file model.ToolFile
		err := core.Db.First(&file, id).Error
		if err != nil {
			return err
		}
		if file.Path != "" {
			pkg.NewUpload().Remove(file.Path, file.Driver)
		}
		return nil
	}).Del(ctx, &model.ToolFile{})
}

func FilesTable(ctx echo.Context) *tableUI.Table {
	table := tableUI.NewTable()
	table.SetUrl("/admin/tools/files/ajax")
	table.SetModel(&[]model.ToolFile{}, "id")

	table.AddFilter("目录", "dir_id")

	table.PreloadModel("Dir")

	table.ModelOrder("id desc")

	table.AddFields(map[string]string{
		"path":  "path",
		"title": "title",
	})

	table.AddCol("ID", "id").SetUI(column.NewContext())
	table.AddCol("文件", "url").SetUI(column.NewContext()).SetNode(node.TNode{
		"nodeName":   "a",
		"vBind:href": "rowData.record.url",
		"target":     "_blank",
		"child":      "{{rowData.record.title}}",
	})

	table.AddCol("目录", "dir.name")
	table.AddCol("时间", "created_at").SetUI(column.NewContext()).SetWidth(250).DataFormat(func(value any, data map[string]any) any {
		duration, _ := time.ParseInLocation(time.RFC3339, cast.ToString(value), time.Local)
		return duration.Format("2006-01-02 15:04:05")
	})

	links := column.NewLink()
	links.AddUrl("删除", "/admin/tools/files/del", map[string]any{
		"id": "id",
	}).SetType("ajax")
	table.AddCol("操作", "").SetWidth(130).SetUI(links)

	var dirData []model.ToolFileDir
	core.Db.Find(&dirData)

	var treeData []map[string]any
	for _, dir := range dirData {
		treeData = append(treeData, map[string]any{
			"key":   dir.ID,
			"title": dir.Name + "(" + dir.HasType + ")",
		})
	}
	sideNode := widget.NewTreeList(ctx.QueryParam("dir_id"), "dir_id").SetData(treeData).Render()

	table.AddSide(&tableUI.Side{
		Direction: "left",
		Node:      sideNode,
		Title:     "文件目录",
		Width:     "200px",
	})

	// 排序规则

	return table
}
