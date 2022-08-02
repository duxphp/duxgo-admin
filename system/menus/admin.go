package menus

import (
	"github.com/duxphp/duxgo/util"
)

func MenuAdmin(menu *util.MenuData) {

	var app *util.MenuData
	var list *util.MenuData

	app = menu.Add(&util.MenuData{
		App:   "index",
		Name:  "控制台",
		Order: 0,
		Icon:  "<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-6 w-6\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6\" /></svg>",
		Url:   "/admin/index",
	})

	app = menu.Add(&util.MenuData{
		App:   "control",
		Name:  "运维",
		Order: 1,
		Icon:  "<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-6 w-6\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\" stroke-width=\"2\">\n  <path stroke-linecap=\"round\" stroke-linejoin=\"round\" d=\"M20 7l-8-4-8 4m16 0l-8 4m8-4v10l-8 4m0-10L4 7m8 4v10M4 7v10l8 4\" />\n</svg>",
	})
	{
		list = app.Group("监控")
		{
			list.Item("实时监控", "/admin/system/control", 0)
			list.Item("历史监控", "/admin/system/controlHistory", 1)
		}

		list = app.Group("日志")
		{
			list.Item("访客记录", "/admin/system/visitorApi", 0)
			list.Item("操作记录", "/admin/system/visitorOperate", 1)
		}
	}

	app = menu.Add(&util.MenuData{
		App:   "system",
		Name:  "设置",
		Order: 100,
		Icon:  "<svg xmlns=\"http://www.w3.org/2000/svg\" class=\"h-6 w-6\" fill=\"none\" viewBox=\"0 0 24 24\" stroke=\"currentColor\"><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M10.325 4.317c.426-1.756 2.924-1.756 3.35 0a1.724 1.724 0 002.573 1.066c1.543-.94 3.31.826 2.37 2.37a1.724 1.724 0 001.065 2.572c1.756.426 1.756 2.924 0 3.35a1.724 1.724 0 00-1.066 2.573c.94 1.543-.826 3.31-2.37 2.37a1.724 1.724 0 00-2.572 1.065c-.426 1.756-2.924 1.756-3.35 0a1.724 1.724 0 00-2.573-1.066c-1.543.94-3.31-.826-2.37-2.37a1.724 1.724 0 00-1.065-2.572c-1.756-.426-1.756-2.924 0-3.35a1.724 1.724 0 001.066-2.573c-.94-1.543.826-3.31 2.37-2.37.996.608 2.296.07 2.572-1.065z\" /><path stroke-linecap=\"round\" stroke-linejoin=\"round\" stroke-width=\"2\" d=\"M15 12a3 3 0 11-6 0 3 3 0 016 0z\" /></svg>",
	})
	{
		list = app.Group("配置管理")
		{
			list.Item("系统设置", "/admin/system/setting", 0)
			list.Item("接口授权", "/admin/system/api", 0)
		}

		list = app.Group("用户管理")
		{
			list.Item("用户管理", "/admin/system/user", 0)
			list.Item("角色管理", "/admin/system/role", 1)
		}
	}

}
