package task

import (
	"context"
	"encoding/json"
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo/core"
	"github.com/hibiken/asynq"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
)

var (
	logger *log.Logger
)

func ControlInit() {
	logger = log.New(os.Stderr, "", 0)
	logger.SetOutput(getLumberjack("monitor"))
}

func getLumberjack(name string) *lumberjack.Logger {
	path := core.Config["app"].GetString("logger.service.path")
	return &lumberjack.Logger{
		Filename:   path + "/" + name + ".log",
		MaxSize:    1024,
		MaxBackups: 0,
		MaxAge:     30,
		Compress:   true,
	}
}

// Control 服务监控
func Control(ctx context.Context, t *asynq.Task) error {
	data := service.GetMonitorData()
	dataJson, err := json.Marshal(data)
	if err == nil {
		logger.Println(string(dataJson))
	}
	return nil
}