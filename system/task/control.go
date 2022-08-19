package task

import (
	"context"
	"encoding/json"
	"github.com/duxphp/duxgo-admin/system/service"
	coreConfig "github.com/duxphp/duxgo/config"
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
	path := coreConfig.Get("app").GetString("logger.service.path")
	return &lumberjack.Logger{
		Filename:   path + "/" + name + ".log",
		MaxSize:    1,
		MaxBackups: 5,
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
