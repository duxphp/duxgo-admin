package admin

import (
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo-admin/tools/model"
	"github.com/duxphp/duxgo-ui/lib/form"
	tableUI "github.com/duxphp/duxgo-ui/lib/table"
	"github.com/duxphp/duxgo-ui/lib/table/column"
	"github.com/duxphp/duxgo-ui/lib/widget"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/util"
	"github.com/labstack/echo/v4"
	"github.com/spf13/cast"
	"gorm.io/gorm"
)

func DistrictList(ctx echo.Context) error {
	return service.NewManageExpand("/admin/tools/district/ajax").SetTable(districtTable).ListPage(ctx)
}

func DistrictAjax(ctx echo.Context) error {
	return service.NewManageExpand("/admin/tools/district/ajax").SetTable(districtTable).ListData(ctx)
}

func DistrictImport(ctx echo.Context) error {
	return service.NewManageExpand("/admin/tools/district/ajax").SetForm(districtForm).FormPage(ctx)
}

func DistrictImportSave(ctx echo.Context) error {
	return service.NewManageExpand("/admin/member/user/ajax").SetForm(districtForm).SetTable(districtTable).FormSave(ctx)
}

func districtTable(ctx echo.Context) *tableUI.Table {
	table := tableUI.NewTable()
	table.SetUrl("/admin/tools/district/ajax")
	table.SetModel(&[]model.ToolDistrict{}, "id")

	table.AddFilter("名称", "name").SetUI(form.NewText()).SetQuick(true)

	table.AddAction().SetUI(widget.NewLink("导入", "/admin/tools/district/import").SetButton().SetType("dialog"))

	table.ModelOrder("id desc")

	table.AddFields(map[string]string{
		"path":  "path",
		"title": "title",
	})

	table.AddCol("编号", "code").SetUI(column.NewContext())
	table.AddCol("地区", "name").SetUI(column.NewContext())
	table.AddCol("类型", "level").DataFormat(func(value any, data map[string]any) any {

		switch cast.ToInt(value) {
		case 3:
			return "街道/乡镇"
		case 2:
			return "区县"
		case 1:
			return "城市"
		default:
			return "省份"
		}
	}).SetUI(column.NewContext()).SetWidth(100)

	return table
}

type DistrictRow struct {
	City         string `csv:"city"`
	District     string `csv:"district"`
	Province     string `csv:"province"`
	DistrictCode string `csv:"district_geocode"`
	CityCode     string `csv:"city_geocode"`
}

func districtForm(ctx echo.Context) *form.Form {
	formUI := form.NewForm()
	data := map[string]any{}
	formUI.SetData(data)
	formUI.SetUrl("/admin/tools/district/importSave")
	formUI.AddField("地区文件", "file").SetUI(form.NewFile().Url("/upload"))

	formUI.SaveFn(func(data map[string]any, key uint) error {

		excelData, err := util.ExcelImport(cast.ToString(data["file"]))
		if err != nil {
			return err
		}

		provinceData := map[string]string{}
		cityData := map[string]map[string]string{}
		districtData := map[string]map[string]string{}
		streetData := map[string]map[string]string{}

		for i, datum := range excelData {
			if i == 0 {
				continue
			}
			if provinceData[datum[1]] == "" {
				provinceData[datum[1]] = datum[0]
			}
			if cityData[datum[1]] == nil {
				cityData[datum[1]] = map[string]string{}
			}
			cityData[datum[1]][datum[3]] = datum[2]

			if districtData[datum[3]] == nil {
				districtData[datum[3]] = map[string]string{}
			}
			districtData[datum[3]][datum[5]] = datum[4]

			if streetData[datum[5]] == nil {
				streetData[datum[5]] = map[string]string{}
			}
			streetData[datum[5]][datum[7]] = datum[6]
		}

		stmt := &gorm.Statement{DB: core.Db}
		stmt.Parse(&model.ToolDistrict{})

		err = core.Db.Exec("TRUNCATE TABLE " + stmt.Schema.Table).Error
		if err != nil {
			return err
		}
		err = core.Db.Transaction(func(tx *gorm.DB) error {
			// 省份
			provinceDB := []*model.ToolDistrict{}
			for code, name := range provinceData {
				provinceDB = append(provinceDB, &model.ToolDistrict{
					Level:    0,
					ParentId: 0,
					Code:     code,
					Name:     name,
				})
			}
			err := tx.Model(&model.ToolDistrict{}).Create(&provinceDB).Error
			if err != nil {
				return err
			}
			// 城市
			cityDB := []*model.ToolDistrict{}
			for _, provinceItem := range provinceDB {
				for code, name := range cityData[provinceItem.Code] {
					cityDB = append(cityDB, &model.ToolDistrict{
						Level:    1,
						ParentId: provinceItem.ID,
						Name:     name,
						Code:     code,
					})
				}
			}
			err = tx.Model(&model.ToolDistrict{}).Create(&cityDB).Error
			if err != nil {
				return err
			}
			// 添加地区
			districtDB := []*model.ToolDistrict{}
			for _, cityItem := range cityDB {
				for code, name := range districtData[cityItem.Code] {
					districtDB = append(districtDB, &model.ToolDistrict{
						Level:    2,
						ParentId: cityItem.ID,
						Name:     name,
						Code:     code,
					})
				}
			}
			err = tx.Model(&model.ToolDistrict{}).Create(&districtDB).Error
			if err != nil {
				return err
			}

			// 添加街道
			for _, districtItem := range districtDB {
				streetDB := []*model.ToolDistrict{}
				for code, name := range streetData[districtItem.Code] {
					streetDB = append(streetDB, &model.ToolDistrict{
						Level:    3,
						ParentId: districtItem.ID,
						Name:     name,
						Code:     code,
					})
				}
				err = tx.Model(&model.ToolDistrict{}).Create(&streetDB).Error
				if err != nil {
					return err
				}
			}

			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})
	return formUI
}
