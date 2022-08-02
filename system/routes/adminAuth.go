package routes

import (
	"github.com/duxphp/duxgo-admin/system/admin"
	"github.com/duxphp/duxgo/util"
)

func RouteAdminAuth(router *util.RouterData) {

	var group *util.RouterData
	var list *util.RouterData

	router.Get("/index", admin.Main, "首页")
	router.Get("/menu", admin.Menu, "菜单")
	router.Get("/notification", admin.Notification, "消息通知")
	router.Get("/fileManage", admin.FileManage, "文件管理器")
	router.Post("/upload", admin.Upload, "上传")

	group = router.Group("/system", "系统管理").Permission()
	{
		list = group.Group("/control", "运行监控")
		{
			list.Get("", admin.Control, "运行监控")
		}

		list = group.Group("/controlHistory", "历史监控")
		{
			list.Get("", admin.ControlHistory, "历史监控")
		}

		list = group.Group("/setting", "系统设置")
		{
			list.Get("", admin.Setting, "列表")
			list.Post("/save", admin.SettingSave, "保存")
		}

		list = group.Group("/api", "接口管理")
		{
			list.Get("", admin.ApiList, "列表")
			list.Get("/ajax", admin.ApiAjax, "数据")
			list.Get("/add", admin.ApiPage, "添加")
			list.Get("/edit", admin.ApiPage, "编辑")
			list.Post("/save", admin.ApiSave, "保存")
			list.Get("/del", admin.ApiDel, "删除")
			list.Post("/status", admin.ApiStatus, "状态")
		}

		list = group.Group("/user", "账号管理")
		{
			list.Get("", admin.UserList, "列表")
			list.Get("/ajax", admin.UserAjax, "数据")
			list.Get("/add", admin.UserPage, "添加")
			list.Get("/edit", admin.UserPage, "编辑")
			list.Post("/save", admin.UserSave, "保存")
		}

		list = group.Group("/role", "角色管理")
		{
			list.Get("", admin.RoleList, "列表")
			list.Get("/ajax", admin.RoleAjax, "数据")
			list.Get("/add", admin.RolePage, "添加")
			list.Get("/edit", admin.RolePage, "编辑")
			list.Post("/save", admin.RoleSave, "保存")
		}

		list = group.Group("/visitorApi", "浏览量统计")
		{
			list.Get("", admin.VisitorApiList, "列表")
			list.Get("/ajax", admin.VisitorApiAjax, "数据")
		}

		list = group.Group("/visitorOperate", "操作统计")
		{
			list.Get("", admin.VisitorOperateList, "列表")
			list.Get("/ajax", admin.VisitorOperateAjax, "数据")
			list.Get("/info", admin.VisitorOperateInfo, "数据")
		}
	}

}
