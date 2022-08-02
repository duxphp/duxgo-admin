package menus

import (
	"github.com/duxphp/duxgo/util"
)

func MenuAdmin(menu *util.MenuData) {

	var app *util.MenuData
	var list *util.MenuData

	app = menu.Add(&util.MenuData{
		App:   "tools",
		Name:  "工具",
		Order: 90,
		Icon:  "<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-6 w-6\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\">\n  <path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M13 10V3L4 14h7v7l9-11h-7z\" />\n</svg>",
	})
	{
		list = app.Group("存储")
		{
			list.Item("文件管理", "/admin/tools/files", 0)
		}
		list = app.Group("地区")
		{
			list.Item("地区管理", "/admin/tools/district", 0)
		}
	}

	app = menu.Push("control")
	{
		list = app.Group("队列")
		{
			list.Item("队列监控", "/admin/tools/queue", 0)
			list.Item("队列管理", "/admin/tools/queueList", 0)
		}
	}

}
