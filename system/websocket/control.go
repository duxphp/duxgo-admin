package websocket

import (
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo/websocket"
)

func ControlData(client *websocket.Client, message *websocket.Message) error {
	if client.Auth != "admin" {
		return nil
	}
	data := service.GetMonitorData()
	client.SendMsg("control", "ok", data)
	return nil
}

func ControlLog(client *websocket.Client, message *websocket.Message) error {
	if client.Auth != "admin" {
		return nil
	}
	data := service.GetMonitorLog()
	client.SendMsg("controlLog", "ok", data)
	return nil
}
