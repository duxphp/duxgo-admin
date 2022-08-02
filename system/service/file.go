package service

import (
	"github.com/duxphp/duxgo-admin/tools/model"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/exception"
	"github.com/duxphp/duxgo/pkg"
	"github.com/duxphp/duxgo/response"
	"github.com/duxphp/duxgo/util/function"
	"github.com/labstack/echo/v4"
	"github.com/panjf2000/ants/v2"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"strings"
	"time"
)

func FileManage(ctx echo.Context, HasType string) error {
	mode := ctx.QueryParam("type")
	var data any
	var err error
	switch mode {
	case "folder":
		data = getFolder(HasType)
	case "files":
		data = getFiles(HasType, ctx)
	case "files-delete":
		deleteFiles(HasType, ctx)
	case "folder-create":
		data, err = createFolder(HasType, ctx)
		if err != nil {
			return err
		}
	case "folder-delete":
		err = deleteFolder(HasType, ctx)
		if err != nil {
			return err
		}
	}

	return response.New(ctx).Send("ok", data)
}

func getFolder(HasType string) any {
	type result struct {
		Name  string `json:"name"`
		DirId int    `json:"dir_id"`
	}
	var data []result
	core.Db.Model(&model.ToolFileDir{}).Where("has_type = ?", HasType).Select("id as dir_id", "name").Find(&data)
	return data
}

func createFolder(HasType string, ctx echo.Context) (any, error) {
	name := ctx.QueryParam("name")
	if name == "" {
		return nil, exception.ParameterError("请输入目录名称")
	}
	d := model.ToolFileDir{
		Name:    name,
		HasType: HasType,
	}
	core.Db.Create(&d)
	return map[string]any{
		"id":   d.ID,
		"name": d.Name,
	}, nil
}

func deleteFolder(HasType string, ctx echo.Context) error {
	dirId := ctx.QueryParam("id")
	if dirId == "" {
		return exception.ParameterError("请选择删除目录")
	}
	var total int64
	core.Db.Model(&model.ToolFileDir{}).Where("has_type = ?", HasType).Count(&total)
	if total <= 1 {
		return exception.BusinessError("请至少保留一个目录")
	}
	var data []model.ToolFile
	core.Db.Model(&model.ToolFile{}).Where("dir_id = ?", dirId).Where("has_type = ?", HasType).Find(&data)

	var ids []int
	for _, datum := range data {
		ids = append(ids, datum.ID)
	}
	// 异步处理删除文件

	p, _ := ants.NewPool(100, ants.WithExpiryDuration(300*time.Second))
	defer p.Release()
	for _, datum := range data {
		p.Submit(func() {
			pkg.NewUpload().Remove(datum.Path, datum.Driver)
		})
	}

	// 删除数据
	core.Db.Delete(&model.ToolFile{}, ids)
	core.Db.Delete(&model.ToolFileDir{}, dirId)
	return nil
}

func getFiles(HasType string, ctx echo.Context) any {
	page := ctx.QueryParam("page")
	dirId := ctx.QueryParam("id")
	query := ctx.QueryParam("keyword")
	filter := ctx.QueryParam("filter")

	formats := map[string][]string{
		"image":    {"jpg", "png", "bmp", "jpeg", "gif"},
		"audio":    {"wav", "mp3", "acc", "ogg"},
		"video":    {"mp4", "ogv", "webm", "ogm"},
		"document": {"doc", "docx", "xls", "xlsx", "pptx", "ppt", "csv", "pdf"},
	}

	if dirId == "" {
		return map[string]any{
			"data":     []any{},
			"total":    0,
			"page":     1,
			"pageSize": 16,
		}
	}

	db := core.Db.Model(&model.ToolFile{}).Where("has_type = ?", HasType).Where("dir_id = ?", dirId)

	if query != "" {
		db.Where("title like ?", "%"+query+"%")
	}
	if filter != "all" {
		if filter == "other" {
			var notFormat []string
			for _, format := range formats {
				notFormat = append(notFormat, format...)
			}
			db.Not(map[string]any{"ext": notFormat})
		} else {
			filters := strings.Split(filter, ",")
			var format []string
			for _, item := range filters {
				format = append(format, formats[item]...)
			}
			db.Where("ext IN ?", format)
		}
	}

	var total int64
	pageSize := 16
	db.Count(&total)
	offset, _ := function.PageLimit(cast.ToInt(page), cast.ToInt(total), pageSize)
	var data []model.ToolFile
	db.Limit(pageSize).Offset(offset).Find(&data)

	newData := []map[string]any{}
	for _, item := range data {
		arr := map[string]any{
			"file_id": item.ID,
			"dir_id":  item.DirId,
			"url":     item.Url,
			"title":   item.Title,
			"ext":     item.Ext,
			"size":    function.FormatFileSize(int64(item.Size)),
			"time":    item.CreatedAt.Format("2006-01-02 15:04:05"),
			"cover":   "",
		}
		_, ok := lo.Find[string](formats["image"], func(i string) bool {
			return i == item.Ext
		})
		if ok {
			arr["cover"] = item.Url
		}
		newData = append(newData, arr)
	}

	return map[string]any{
		"data":     newData,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	}
}

func deleteFiles(HasType string, ctx echo.Context) {
	id := ctx.QueryParam("id")
	var file model.ToolFile
	core.Db.First(&file, id)
	if file.Path != "" {
		pkg.NewUpload().Remove(file.Path, file.Driver)
	}
	core.Db.Delete(file, id)
}
