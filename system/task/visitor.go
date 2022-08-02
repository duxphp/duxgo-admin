package task

import (
	"bufio"
	"context"
	"github.com/duxphp/duxgo-admin/system/model"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/exception"
	"github.com/hibiken/asynq"
	"github.com/tidwall/gjson"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"io"
	"net/url"
	"os"
	"time"
)

type visitor struct {
	Uv      int
	Pv      int
	MinTime float64
	MaxTime float64
}

type logDataT struct {
	Id      string
	Uri     string
	Method  string
	Time    string
	Ip      string
	Latency float64
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

// Visitor 访客统计
func Visitor(ctx context.Context, t *asynq.Task) error {
	// 统计访客
	requestID, _ := core.Redis.Get(core.Ctx, "visitor:requestID").Result()
	// 打开日志文件
	fd, err := os.Open(core.Config["app"].GetString("logger.request.path"))
	if err != nil {
		return nil
	}
	defer fd.Close()
	bufferRead := bufio.NewReader(fd)
	logData := []*logDataT{}

	var lastId string
	status := false
	for {
		line, err := bufferRead.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}
		lineData := gjson.Parse(line)
		data := logDataT{
			Id:      lineData.Get("id").String(),
			Uri:     lineData.Get("uri").String(),
			Method:  lineData.Get("method").String(),
			Time:    lineData.Get("time").String(),
			Ip:      lineData.Get("ip").String(),
			Latency: lineData.Get("latency").Float(),
		}
		if !status && (requestID == "" || requestID == data.Id) {
			status = true
		}
		if !status {
			continue
		}
		logData = append(logData, &data)
		lastId = data.Id
	}

	visitorData := map[string]*model.VisitorApi{}
	visitorUv := map[string]*visitor{}

	for _, item := range logData {
		urlParse, err := url.Parse(item.Uri)
		if err != nil {
			continue
		}
		t, _ := time.Parse(time.RFC3339, item.Time)
		date := t.Format("2006-01-02")
		key := urlParse.Path + ":" + item.Method + ":" + date
		if visitorData[key] == nil {
			var info = model.VisitorApi{
				Date:   datatypes.Date(t),
				Url:    urlParse.Path,
				Method: item.Method,
			}
			err = core.Db.Model(&model.VisitorApi{}).Where("url = ?", urlParse.Path).Where("method = ?", item.Method).Where("date = ?", date).FirstOrCreate(&info).Error
			if err != nil {
				exception.Error(err)
				continue
			}
			visitorData[key] = &info
		}

		if visitorUv[key] == nil {
			visitorUv[key] = &visitor{
				Uv:      1,
				Pv:      1,
				MinTime: 0,
				MaxTime: 0,
			}
		}
		visitorUv[key].Pv += 1
		if visitorUv[key].MinTime == 0 || visitorUv[key].MinTime > item.Latency {
			visitorUv[key].MinTime = item.Latency
		}
		if visitorUv[key].MaxTime == 0 || visitorUv[key].MaxTime < item.Latency {
			visitorUv[key].MaxTime = item.Latency
		}

		pvKey := "visit:" + key + ":" + item.Ip
		pvLock, _ := core.Redis.Get(core.Ctx, pvKey).Result()
		if pvLock == "" {
			core.Redis.Set(core.Ctx, pvKey, 1, 24*time.Hour)
			visitorUv[key].Uv += 1
		}
	}

	for k, v := range visitorUv {
		visit := visitorData[k]
		data := map[string]any{
			"uv": gorm.Expr("uv + ?", v.Uv),
			"pv": gorm.Expr("pv + ?", v.Pv),
		}
		if visit.MinTime == 0 || v.MinTime < visit.MinTime {
			data["min_time"] = v.MinTime
		}
		if visit.MaxTime == 0 || v.MaxTime > visit.MaxTime {
			data["max_time"] = v.MaxTime
		}
		err := core.Db.Model(&model.VisitorApi{}).Where("id = ?", visit.ID).Updates(data).Error
		if err != nil {
			exception.Error(err)
			continue
		}
	}
	core.Redis.Set(core.Ctx, "visitor:requestID", lastId, 0)
	return nil
}
