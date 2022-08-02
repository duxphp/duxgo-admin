package routes

import (
	"github.com/duxphp/duxgo-admin/tools/admin"
	"github.com/duxphp/duxgo-admin/tools/service"
	"github.com/duxphp/duxgo/util"
)

func RouteAdminAuth(router *util.RouterData) {

	router.Get("/tools/area", service.AreaTree, "地区数据")

	var group *util.RouterData
	var list *util.RouterData
	group = router.Group("/tools", "工具").Permission()
	{
		list = group.Group("/queue", "队列监控")
		{
			list.Get("", admin.QueueMain, "概况")
		}

		list = group.Group("/queueList", "队列管理")
		{
			list.Get("", admin.QueueList, "列表")
			list.Get("/ajax", admin.QueueAjax, "数据")
		}

		list = group.Group("/files", "文件管理")
		{
			list.Get("", admin.FilesList, "列表")
			list.Get("/ajax", admin.FilesAjax, "数据")
			list.Get("/del", admin.FilesDel, "删除")
		}

		list = group.Group("/district", "地区管理")
		{
			list.Get("", admin.DistrictList, "列表")
			list.Get("/ajax", admin.DistrictAjax, "数据")
			list.Get("/import", admin.DistrictImport, "编辑")
			list.Post("/importSave", admin.DistrictImportSave, "保存")
		}

	}

}
