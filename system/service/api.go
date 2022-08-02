package service

import (
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo/core"
)

// ServiceApi 接口数据
var ServiceApi map[string]*apiItem

type apiItem struct {
	SecretId  string
	SecretKey string
}

func InitApi() {
	// 查询Api数据
	var data []model.SystemApi
	core.Db.Where("status = ?", true).Find(&data)
	apiData := map[string]*apiItem{}
	for _, datum := range data {
		apiData[datum.SecretId] = &apiItem{
			SecretId:  datum.SecretId,
			SecretKey: datum.SecretKey,
		}
	}
	ServiceApi = apiData
}
