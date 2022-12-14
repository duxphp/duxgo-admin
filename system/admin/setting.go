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

	formUI.AddHeader(widget.NewAlert("??????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????????", "????????????").SetType(widget.AlertWarning))

	formUI.AddColumn(formLayout.NewTab(), func(element form.ILayout) {
		element.Column(func(formSub *form.Form) {
			formSub.AddField("????????????", "info.name").SetUI(form.NewText())
			formSub.AddField("????????????", "info.description").SetUI(form.NewText())
			formSub.AddField("????????????", "info.copyright").SetUI(form.NewText())
		}, formLayout.TabArgs{
			Name: "????????????",
		})

		element.Column(func(formSub *form.Form) {
			formSub.AddField("????????????", "app.baseUrl").SetUI(form.NewText())
			formSub.AddField("??????????????????", "logger.default.level").SetUI(form.NewSelect().SetOptions(map[any]any{
				"debug": "debug",
				"info":  "info",
				"warn":  "warn",
				"error": "error",
				"panic": "panic",
				"fatal": "fatal",
			}))
			formSub.AddField("?????????????????????", "logger.db.level").SetUI(form.NewSelect().SetOptions(map[any]any{
				"debug": "debug",
				"info":  "info",
				"warn":  "warn",
				"error": "error",
				"panic": "panic",
				"fatal": "fatal",
			}))
			formSub.AddField("????????????", "logger.request.status").SetUI(form.NewRadio().SetOptions([]form.RadioOptions{
				{
					Key:  true,
					Name: "??????",
				},
				{
					Key:  false,
					Name: "??????",
				},
			})).SetHelp("???????????????????????????????????????????????????????????????")
			formSub.AddField("??????????????????", "logger.request.level").SetUI(form.NewSelect().SetOptions(map[any]any{
				"debug": "debug",
				"info":  "info",
				"warn":  "warn",
				"error": "error",
				"panic": "panic",
				"fatal": "fatal",
			}))
		}, formLayout.TabArgs{
			Name: "????????????",
		})

		element.Column(func(formSub *form.Form) {
			formSub.AddField("????????????", "storage.driver.type").SetUI(form.NewSelect().SetOptions(map[any]any{
				"local": "????????????",
				"qiniu": "????????????",
			}))
			formSub.AddLayout(formLayout.NewBlock("????????????"), func(f *form.Form) {
				f.AddField("accountName", "storage.driver.qiniu.accountName").SetUI(form.NewText())
				f.AddField("accountkey", "storage.driver.qiniu.accountkey").SetUI(form.NewText())
				f.AddField("????????????", "storage.driver.qiniu.bucket").SetUI(form.NewText())
				f.AddField("????????????", "storage.driver.qiniu.domain").SetUI(form.NewText())
			})

		}, formLayout.TabArgs{
			Name: "????????????",
		})

		element.Column(func(formSub *form.Form) {
			formSub.AddField("????????????", "storage.imageResize.status").SetUI(form.NewRadio().SetOptions([]form.RadioOptions{
				{
					Key:  true,
					Name: "??????",
				},
				{
					Key:  false,
					Name: "??????",
				},
			}))
			formSub.AddColumn(formLayout.NewRow(), func(element form.ILayout) {
				element.Column(func(f *form.Form) {
					f.AddField("????????????", "storage.imageResize.width").SetUI(form.NewNumber())
				}, 12)

				element.Column(func(f *form.Form) {
					f.AddField("????????????", "storage.imageResize.height").SetUI(form.NewNumber())
				}, 12)
			})
			formSub.AddField("????????????", "storage.imageWater.status").SetUI(form.NewRadio().SetOptions([]form.RadioOptions{
				{
					Key:  true,
					Name: "??????",
				},
				{
					Key:  false,
					Name: "??????",
				},
			}))
			formSub.AddField("???????????????", "storage.imageWater.opacity").SetUI(form.NewNumber().SetStep(0.1, 1).SetLimit(0.1, 1.0))
			formSub.AddField("????????????", "storage.imageWater.position").SetUI(form.NewSelect().SetOptions(map[any]any{
				0: "?????????",
				1: "?????????",
				2: "?????????",
				3: "?????????",
				4: "??????",
				5: "?????????",
				6: "?????????",
				7: "?????????",
				8: "?????????",
			}))
			formSub.AddField("????????????", "storage.imageWater.margin").SetUI(form.NewText())
		}, formLayout.TabArgs{
			Name: "????????????",
		})

		element.Column(func(formSub *form.Form) {
			formSub.AddField("????????????", "message.driver.type").SetUI(form.NewSelect().SetOptions(map[any]any{
				"chuanglan": "????????????",
			}))
			formSub.AddLayout(formLayout.NewBlock("???????????????"), func(f *form.Form) {
				formSub.AddField("???????????????", "message.tpl.code").SetUI(form.NewText())
				formSub.AddField("?????????????????????", "message.code.expired").SetUI(form.NewNumber().SetLimit(0, 1000))
				formSub.AddField("?????????????????????", "message.code.retry").SetUI(form.NewNumber().SetLimit(0, 1000))
			})

			formSub.AddLayout(formLayout.NewBlock("????????????"), func(f *form.Form) {
				f.AddField("????????????", "message.driver.chuanglan.account").SetUI(form.NewText())
				f.AddField("????????????", "message.driver.chuanglan.password").SetUI(form.NewText())
				f.AddField("????????????", "message.driver.chuanglan.url").SetUI(form.NewText())
			})
		}, formLayout.TabArgs{
			Name: "????????????",
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
