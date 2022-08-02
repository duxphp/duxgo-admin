package admin

import (
	"embed"
	"github.com/duxphp/duxgo-admin/system"
	"github.com/duxphp/duxgo-admin/tools"
	"github.com/duxphp/duxgo/core"
	"html/template"
)

//go:embed */views/*
var ViewsFs embed.FS

func New() {
	template.Must(core.Tpl.ParseFS(ViewsFs, "*/views/*"))
	tools.App()
	system.App()
}
