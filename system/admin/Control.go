package admin

import (
	"fmt"
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo/response"
	"github.com/labstack/echo/v4"
)

func Control(ctx echo.Context) error {
	info := service.GetMonitorInfo()
	assign := map[string]any{}

	fmt.Println(info)

	assign["info"] = []map[string]any{
		//{
		//	"name":  "启动时间",
		//	"color": "arcoblue",
		//	"value": info.BootTime,
		//},
		//{
		//	"name":  "操作系统",
		//	"color": "red",
		//	"value": info.OsName,
		//},
	}

	assign["dir"] = []map[string]any{
		//{
		//	"dir":      "/logs",
		//	"name":     "日志",
		//	"size":     info.LogSize,
		//	"size_str": info.LogSizeF,
		//},
		//{
		//	"dir":      "/uploads",
		//	"name":     "上传",
		//	"size":     info.UploadSize,
		//	"size_str": info.UploadSizeF,
		//},
		//{
		//	"dir":      "/tmp",
		//	"name":     "缓存",
		//	"size":     info.UploadSize,
		//	"size_str": info.UploadSizeF,
		//},
	}

	return response.New(ctx).Render("adminControl.gohtml", assign)
}
