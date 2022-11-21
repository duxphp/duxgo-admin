package service

import (
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo/core"
)

// ConfigGet 获取配置
func ConfigGet(hasType string, hasId string) map[string]any {
	info := model.SystemConfig{}
	core.Db.Model(model.SystemConfig{}).Where("has_type", hasType).Where("has_id", hasId).Find(&info)
	return info.Data
}

// ConfigGetValue 获取配置值
func ConfigGetValue[T any](hasType string, hasId string, key string) any {
	info := model.SystemConfig{}
	core.Db.Model(model.SystemConfig{}).Where("has_type", hasType).Where("has_id", hasId).Find(&info)
	return info.Data[key]
}

// ConfigSave 保存配置
func ConfigSave(hasType string, hasId string, data map[string]any) {
	info := model.SystemConfig{}
	core.Db.Model(model.SystemConfig{}).Where("has_type", hasType).Where("has_id", hasId).Find(&info)
	info.Data = data
	core.Db.Save(info)
}

// ConfigSaveValue 保存配置值
func ConfigSaveValue(hasType string, hasId string, key string, value any) {
	info := model.SystemConfig{}
	core.Db.Model(model.SystemConfig{}).Where("has_type", hasType).Where("has_id", hasId).Find(&info)
	info.Data[key] = value
	core.Db.Save(info)
}
