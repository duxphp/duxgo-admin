package tools

import (
	"github.com/duxphp/duxgo-admin/tools/menus"
	"github.com/duxphp/duxgo-admin/tools/model"
	"github.com/duxphp/duxgo-admin/tools/routes"
	"github.com/duxphp/duxgo/core"
	coreRegister "github.com/duxphp/duxgo/register"
)

var config = struct {
}{}

func App() {
	coreRegister.App(&coreRegister.AppConfig{
		Name:      "tools",
		Config:    &config,
		Model:     Model,
		Route:     Route,
		RouteAuth: RouteAuth,
		Menu:      Menu,
	})
}

func Model() {
	var err error
	dirStatus := core.Db.Migrator().HasTable("app_tools_file_dir")
	err = core.Db.AutoMigrate(&model.ToolFileDir{}, &model.ToolFile{}, &model.ToolDistrict{})
	if err != nil {
		panic("model migrate error：" + err.Error())
		return
	}
	if dirStatus == false {
		data := model.ToolFileDir{Name: "默认", HasType: "admin"}
		core.Db.Create(&data)
	}
}

func Route(router coreRegister.Router) {
	routes.RouteApi(router["api"])
	routes.RouteWeb(router["web"])
}

func RouteAuth(router coreRegister.Router) {
	routes.RouteAdminAuth(router["adminAuth"])
}

func Menu(menu coreRegister.Menu) {
	menus.MenuAdmin(menu["admin"])
}
