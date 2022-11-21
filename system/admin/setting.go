package admin

import (
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo-ui/lib/form"
	"github.com/duxphp/duxgo-ui/lib/formLayout"
	"github.com/duxphp/duxgo-ui/lib/widget"
	coreConfig "github.com/duxphp/duxgo/config"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
)

func Setting(ctx echo.Context) error {
	return service.NewManageExpand().SetForm(settingForm).FormPage(ctx)
}

func SettingSave(ctx echo.Context) error {
	return service.NewManageExpand().SetForm(settingForm).FormSave(ctx)
}

func settingForm(ctx echo.Context) *form.Form {

	formUI := form.NewForm()
	formUI.SetBack(false)

	infoConfig := coreConfig.Get("info")
	appConfig := coreConfig.Get("app")
	storageConfig := coreConfig.Get("storage")

	data := map[string]any{
		"info.name":             infoConfig.Get("info.name"),
		"info.description":      infoConfig.Get("info.description"),
		"info.copyright":        infoConfig.Get("info.copyright"),
		"app.baseUrl":           appConfig.Get("app.baseUrl"),
		"logger.default.level":  appConfig.Get("logger.default.level"),
		"logger.db.level":       appConfig.Get("logger.db.level"),
		"logger.request.level":  appConfig.Get("logger.request.level"),
		"logger.request.status": appConfig.Get("logger.request.status"),

		"storage.driver.type":              storageConfig.Get("driver.type"),
		"storage.driver.qiniu.accountName": storageConfig.Get("driver.qiniu.accountName"),
		"storage.driver.qiniu.accountkey":  storageConfig.Get("driver.qiniu.accountkey"),
		"storage.driver.qiniu.bucket":      storageConfig.Get("driver.qiniu.bucket"),
		"storage.driver.qiniu.region":      storageConfig.Get("driver.qiniu.region"),
		"storage.driver.qiniu.domain":      storageConfig.Get("driver.qiniu.domain"),

		"storage.imageResize.status":  storageConfig.GetBool("imageResize.status"),
		"storage.imageResize.width":   storageConfig.GetInt("imageResize.width"),
		"storage.imageResize.height":  storageConfig.GetInt("imageResize.height"),
		"storage.imageWater.status":   storageConfig.GetBool("imageWater.status"),
		"storage.imageWater.opacity":  storageConfig.GetFloat64("imageWater.opacity"),
		"storage.imageWater.position": storageConfig.GetInt("imageWater.position"),
		"storage.imageWater.margin":   storageConfig.GetInt("imageWater.margin"),
	}

	formUI.SetData(data)
	formUI.SetDialog(false)
	formUI.SetUrl("/system/setting/save")

	formUI.AddHeader(widget.NewAlert("系统设置选项为运维人员便捷使用，非专业人士或不清楚选项请勿随意修改，更改部分配置后需重启服务", "安全提示").SetType(widget.AlertWarning))

	formUI.AddColumn(formLayout.NewTab(), func(element form.ILayout) {
		element.Column(func(formSub *form.Form) {
			formSub.AddField("系统名称", "info.name").SetUI(form.NewText())
			formSub.AddField("系统描述", "info.description").SetUI(form.NewText())
			formSub.AddField("版权信息", "info.copyright").SetUI(form.NewText())
		}, formLayout.TabArgs{
			Name: "系统信息",
		})

		element.Column(func(formSub *form.Form) {
			formSub.AddField("系统域名", "app.baseUrl").SetUI(form.NewText())
			formSub.AddField("系统日志等级", "logger.default.level").SetUI(form.NewSelect().SetOptions(map[any]any{
				"debug": "debug",
				"info":  "info",
				"warn":  "warn",
				"error": "error",
				"panic": "panic",
				"fatal": "fatal",
			}))
			formSub.AddField("数据库日志等级", "logger.db.level").SetUI(form.NewSelect().SetOptions(map[any]any{
				"debug": "debug",
				"info":  "info",
				"warn":  "warn",
				"error": "error",
				"panic": "panic",
				"fatal": "fatal",
			}))
			formSub.AddField("请求日志", "logger.request.status").SetUI(form.NewRadio().SetOptions([]form.RadioOptions{
				{
					Key:  true,
					Name: "开启",
				},
				{
					Key:  false,
					Name: "关闭",
				},
			})).SetHelp("关闭该日志或日志等级太高会影响访客统计信息")
			formSub.AddField("请求日志等级", "logger.request.level").SetUI(form.NewSelect().SetOptions(map[any]any{
				"debug": "debug",
				"info":  "info",
				"warn":  "warn",
				"error": "error",
				"panic": "panic",
				"fatal": "fatal",
			}))
		}, formLayout.TabArgs{
			Name: "运行配置",
		})

		element.Column(func(formSub *form.Form) {
			formSub.AddField("默认存储", "storage.driver.type").SetUI(form.NewSelect().SetOptions(map[any]any{
				"local": "本地存储",
				"qiniu": "七牛存储",
			}))
			formSub.AddLayout(formLayout.NewBlock("七牛配置"), func(f *form.Form) {
				f.AddField("accountName", "storage.driver.qiniu.accountName").SetUI(form.NewText())
				f.AddField("accountkey", "storage.driver.qiniu.accountkey").SetUI(form.NewText())
				f.AddField("空间名称", "storage.driver.qiniu.bucket").SetUI(form.NewText())
				f.AddField("绑定域名", "storage.driver.qiniu.domain").SetUI(form.NewText())
			})

		}, formLayout.TabArgs{
			Name: "文件存储",
		})

		element.Column(func(formSub *form.Form) {
			formSub.AddField("自动缩图", "storage.imageResize.status").SetUI(form.NewRadio().SetOptions([]form.RadioOptions{
				{
					Key:  true,
					Name: "开启",
				},
				{
					Key:  false,
					Name: "关闭",
				},
			}))
			formSub.AddColumn(formLayout.NewRow(), func(element form.ILayout) {
				element.Column(func(f *form.Form) {
					f.AddField("最大宽度", "storage.imageResize.width").SetUI(form.NewNumber())
				}, 12)

				element.Column(func(f *form.Form) {
					f.AddField("最大高度", "storage.imageResize.height").SetUI(form.NewNumber())
				}, 12)
			})
			formSub.AddField("图片水印", "storage.imageWater.status").SetUI(form.NewRadio().SetOptions([]form.RadioOptions{
				{
					Key:  true,
					Name: "开启",
				},
				{
					Key:  false,
					Name: "关闭",
				},
			}))
			formSub.AddField("水印透明度", "storage.imageWater.opacity").SetUI(form.NewNumber().SetStep(0.1, 1).SetLimit(0.1, 1.0))
			formSub.AddField("水印位置", "storage.imageWater.position").SetUI(form.NewSelect().SetOptions(map[any]any{
				0: "上居中",
				1: "左上角",
				2: "右上角",
				3: "左居中",
				4: "居中",
				5: "右居中",
				6: "下居中",
				7: "左下角",
				8: "右下角",
			}))
			formSub.AddField("水印边距", "storage.imageWater.margin").SetUI(form.NewText())
		}, formLayout.TabArgs{
			Name: "图片存储",
		})

		element.Column(func(formSub *form.Form) {
			formSub.AddField("默认短信", "message.driver.type").SetUI(form.NewSelect().SetOptions(map[any]any{
				"chuanglan": "创蓝短信",
			}))
			formSub.AddLayout(formLayout.NewBlock("验证码配置"), func(f *form.Form) {
				formSub.AddField("验证码模板", "message.tpl.code").SetUI(form.NewText())
				formSub.AddField("过期时间（秒）", "message.code.expired").SetUI(form.NewNumber().SetLimit(0, 1000))
				formSub.AddField("重试时间（秒）", "message.code.retry").SetUI(form.NewNumber().SetLimit(0, 1000))
			})

			formSub.AddLayout(formLayout.NewBlock("创蓝配置"), func(f *form.Form) {
				f.AddField("接口账号", "message.driver.chuanglan.account").SetUI(form.NewText())
				f.AddField("接口密码", "message.driver.chuanglan.password").SetUI(form.NewText())
				f.AddField("接口网址", "message.driver.chuanglan.url").SetUI(form.NewText())
			})
		}, formLayout.TabArgs{
			Name: "短信发送",
		})

	})

	formUI.SaveFn(func(data map[string]any, key uint) error {
		var err error
		infoConfig.Set("info.name", data["info.name"])
		infoConfig.Set("info.description", data["info.description"])
		infoConfig.Set("info.copyright", data["info.copyright"])
		err = infoConfig.WriteConfig()
		if err != nil {
			return err
		}
		appConfig.Set("app.baseUrl", data["app.baseUrl"])
		appConfig.Set("logger.default.level", data["logger.default.level"])
		appConfig.Set("logger.db.level", data["logger.db.level"])
		appConfig.Set("logger.request.level", data["logger.request.level"])
		appConfig.Set("logger.request.status", data["logger.request.status"])

		err = appConfig.WriteConfig()
		if err != nil {
			return err
		}
		storageConfig.Set("driver.type", data["storage.driver.type"])
		storageConfig.Set("driver.qiniu.accountName", data["storage.driver.qiniu.accountName"])
		storageConfig.Set("driver.qiniu.accountkey", data["storage.driver.qiniu.accountkey"])
		storageConfig.Set("driver.qiniu.bucket", data["storage.driver.qiniu.bucket"])
		storageConfig.Set("driver.qiniu.region", data["storage.driver.qiniu.region"])
		storageConfig.Set("driver.qiniu.domain", data["storage.driver.qiniu.domain"])
		storageConfig.Set("imageResize.status", data["storage.imageResize.status"])
		storageConfig.Set("imageResize.width", cast.ToInt(data["storage.imageResize.width"]))
		storageConfig.Set("imageResize.height", cast.ToInt(data["storage.imageResize.height"]))
		storageConfig.Set("imageWater.status", cast.ToBool(data["storage.imageWater.status"]))
		storageConfig.Set("imageWater.opacity", cast.ToFloat64(data["storage.imageWater.opacity"]))
		storageConfig.Set("imageWater.position", cast.ToInt(data["storage.imageWater.position"]))
		storageConfig.Set("imageWater.margin", cast.ToInt(data["storage.imageWater.margin"]))
		err = storageConfig.WriteConfig()
		if err != nil {
			return err
		}

		// syscall.Kill(syscall.Getpid(), syscall.SIGINT)
		return nil
	})
	return formUI
}
