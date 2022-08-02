package system

import (
	"github.com/duxphp/duxgo-admin/system/menus"
	"github.com/duxphp/duxgo-admin/system/middleware"
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo-admin/system/routes"
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo-admin/system/task"
	"github.com/duxphp/duxgo-admin/system/websocket"
	"github.com/duxphp/duxgo/core"
	coreMiddleware "github.com/duxphp/duxgo/middleware"
	coreRegister "github.com/duxphp/duxgo/register"
	coreTask "github.com/duxphp/duxgo/task"
	"github.com/duxphp/duxgo/util"
	coreWebsocket "github.com/duxphp/duxgo/websocket"
	"github.com/hibiken/asynq"
	"github.com/labstack/echo/v4"
)

var config = struct{}{}

func App() {
	// 注册应用
	coreRegister.App(&coreRegister.AppConfig{
		Name:         "system",
		Config:       &config,
		Register:     Register,
		AppRoute:     AppRoute,
		AppRouteAuth: AppRouteAuth,
		AppMenu:      AppMenu,
		Model:        Model,
		Route:        Route,
		RouteAuth:    RouteAuth,
		Menu:         Menu,
		Scheduler:    Scheduler,
		Queue:        Queue,
		Websocket:    Websocket,
	})
}

func AppRoute(router coreRegister.Router, app *echo.Echo) {
	router["web"] = util.NewRouter(app.Group("/"))
	router["admin"] = util.NewRouter(app.Group("/admin"))
	router["api"] = util.NewRouter(app.Group("/api", middleware.ApiHandler))

}

func AppRouteAuth(router coreRegister.Router, app *echo.Echo) {
	router["adminAuth"] = util.NewRouter(app.Group("/admin", middleware.AdminViewHandler, coreMiddleware.AuthJwt("admin"), middleware.Operate("admin")))
}

func AppMenu(menus coreRegister.Menu) {
	menus["admin"] = util.NewMenu()
}

func Model() {
	var err error
	err = core.Db.AutoMigrate(&model.SystemUser{}, &model.SystemRole{}, &model.SystemApi{}, &model.SystemConfig{}, &model.VisitorApi{}, &model.VisitorOperate{})
	if err != nil {
		panic("model migrate error：" + err.Error())
		return
	}
}

func Route(router coreRegister.Router) {
	routes.RouteAdmin(router["admin"])
	routes.RouteWeb(router["web"])
}

func RouteAuth(router coreRegister.Router) {
	routes.RouteAdminAuth(router["adminAuth"])
}

func Menu(menu coreRegister.Menu) {
	menus.MenuAdmin(menu["admin"])
}

func Register(app *echo.Echo) {
	// 注册接口服务
	service.InitApi()
}

func Queue(queue *asynq.ServeMux) {
	task.ControlInit()
	queue.HandleFunc("system.visitor", task.Visitor)
	queue.HandleFunc("system.control", task.Control)
}

func Scheduler(scheduler *asynq.Scheduler) {
	// 服务器监控
	if core.Config["app"].GetBool("logger.service.status") {
		coreTask.RegScheduler("*/1 * * * *", "system.control", map[string]any{}, coreTask.PRIORITY_LOW)
	}
	// 保存访客日志
	if core.Config["app"].GetBool("logger.request.status") {
		coreTask.RegScheduler("*/30 * * * *", "system.visitor", map[string]any{}, coreTask.PRIORITY_LOW)
	}
}

func Websocket() {
	// 实时数据
	coreWebsocket.Event("control", websocket.ControlData)
	// 历史数据
	coreWebsocket.Event("controlLog", websocket.ControlLog)
}
