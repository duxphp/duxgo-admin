package service

import (
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo/core"
	"gorm.io/datatypes"
)

// ConfigGet 获取配置
func ConfigGet(hasType string, hasId uint) map[string]any {
	info := model.SystemConfig{}
	core.Db.Model(model.SystemConfig{}).Where("has_type", hasType).Where("has_id", hasId).Find(&info)
	return info.Data
}

// ConfigGetValue 获取配置值
func ConfigGetValue[T any](hasType string, hasId uint, key string) any {
	info := model.SystemConfig{}
	core.Db.Model(model.SystemConfig{}).Where("has_type", hasType).Where("has_id", hasId).Find(&info)
	return info.Data[key]
}

// ConfigSave 保存配置
func ConfigSave(hasType string, hasId uint, data map[string]any) {
	info := model.SystemConfig{}
	core.Db.Model(model.SystemConfig{}).Where("has_type", hasType).Where("has_id", hasId).FirstOrCreate(&info, model.SystemConfig{
		HasType: hasType,
		HasId:   hasId,
	})
	core.Db.Model(model.SystemConfig{}).Where("has_type", hasType).Where("has_id", hasId).Update("data", datatypes.JSONMap(data))
}

// ConfigSaveValue 保存配置值
func ConfigSaveValue(hasType string, hasId uint, key string, value any) {
	info := model.SystemConfig{}
	core.Db.Model(model.SystemConfig{}).Where("has_type", hasType).Where("has_id", hasId).FirstOrCreate(&info, model.SystemConfig{
		HasType: hasType,
		HasId:   hasId,
	})
	info.Data[key] = value
	core.Db.Model(model.SystemConfig{}).Where("has_type", hasType).Where("has_id", hasId).Update("data", datatypes.JSONMap(info.Data))
}
