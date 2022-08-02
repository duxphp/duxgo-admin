package service

import (
	"bufio"
	"encoding/json"
	"github.com/dustin/go-humanize"
	coreConfig "github.com/duxphp/duxgo/config"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/util/function"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/process"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"runtime/pprof"
	"time"
)

type MonitorInfo struct {
	OsName      string // 操作系统
	BootTime    string // 启动时间
	LogSize     uint64 // 日志大小
	LogSizeF    string // 日志大小格式化
	UploadSize  uint64 // 上传大小
	UploadSizeF string // 上传大小格式化
	TmpSize     uint64 // 缓存大小
	TmpSizeF    string // 缓存大小格式化
}

// GetMonitorInfo 获取监控信息
func GetMonitorInfo() *MonitorInfo {
	data := MonitorInfo{}
	data.LogSize = getDirSize("/logs")
	data.LogSizeF = humanize.Bytes(data.LogSize)
	data.UploadSize = getDirSize("/uploads")
	data.UploadSizeF = humanize.Bytes(data.UploadSize)
	data.TmpSize = getDirSize("/tmp")
	data.TmpSizeF = humanize.Bytes(data.TmpSize)
	data.BootTime = core.BootTime.Format("2006-01-02 15:04:05")
	sysInfo, _ := host.Info()
	data.OsName = sysInfo.Platform + " " + sysInfo.PlatformVersion
	return &data

}

type MonitorData struct {
	CpuPercent     float64
	MemPercent     float64
	ThreadCount    int
	GoroutineCount int
	Timestamp      int64
}

// GetMonitorData 获取监控数据
func GetMonitorData() *MonitorData {
	// CPU占用率
	p, _ := process.NewProcess(int32(os.Getpid()))
	cpuPercent, _ := p.Percent(time.Second)
	// 内存占用率
	memPercent, _ := p.MemoryPercent()
	// 创建的线程数
	threadCount := pprof.Lookup("threadcreate").Count()
	// Goroutine数
	GoroutineCount := runtime.NumGoroutine()

	return &MonitorData{
		CpuPercent:     function.Round(cpuPercent, 2),
		MemPercent:     function.Round(float64(memPercent), 2),
		ThreadCount:    threadCount,
		GoroutineCount: GoroutineCount,
		Timestamp:      time.Now().UnixMilli(),
	}
}

// GetMonitorLog 获取监控日志
func GetMonitorLog() []map[string]any {
	path := coreConfig.Get("app").GetString("logger.service.path")
	loadFiles, _ := filepath.Glob(path + "/monitor*.log")
	loadData := passingFiles(loadFiles)
	return loadData
}

func getDirSize(path string) uint64 {
	var size int64
	wd, _ := os.Getwd()
	filepath.Walk(wd+path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})
	return uint64(size)
}

// 解析文件
func passingFiles(files []string) []map[string]any {
	loadData := []map[string]any{}
	for _, file := range files {
		fileData, err := parsingFile(file)
		if err != nil {
			continue
		}
		loadData = append(loadData, fileData...)
	}
	return loadData
}

// 解析单文件
func parsingFile(file string) ([]map[string]any, error) {
	fd, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer fd.Close()
	bufferRead := bufio.NewReader(fd)
	data := []map[string]any{}
	for {
		line, err := bufferRead.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		curData := map[string]any{}
		err = json.Unmarshal([]byte(line), &curData)
		if err != nil {
			continue
		}
		data = append(data, curData)
	}
	return data, nil
}
