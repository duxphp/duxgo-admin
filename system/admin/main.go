package admin

import (
	"fmt"
	"github.com/duxphp/duxgo-admin/system/model"
	toolsModel "github.com/duxphp/duxgo-admin/tools/model"
	coreConfig "github.com/duxphp/duxgo/config"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/pkg/chart"
	coreRegister "github.com/duxphp/duxgo/register"
	"github.com/duxphp/duxgo/response"
	"github.com/go-resty/resty/v2"
	"github.com/golang-module/carbon/v2"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"syscall"
	"time"
)

var menuData *map[string]any

type statsData struct {
	Label string  `mapstructure:"label"`
	Value float64 `mapstructure:"value"`
}

func Main(ctx echo.Context) error {

	assign := map[string]any{}
	startTime := carbon.Now().SubDays(6).ToDateString()

	// 访问量
	apiNumData := []*statsData{}
	err := core.Db.Model(&model.VisitorApi{}).Select(`SUM(pv) as value, DATE_FORMAT(date, "%Y-%m-%d") as label`).Where("date  >= ?", startTime).Group("date").Find(&apiNumData).Error
	if err != nil {
		return err
	}
	assign["apiNum"] = CoverData(apiNumData)

	apiNumMaps := []map[string]any{}
	err = mapstructure.Decode(apiNumData, &apiNumMaps)
	if err != nil {
		return err
	}
	apiNumChart := chart.New("date").Legend(true, "right", "top").Line().Date(startTime, carbon.Now().ToDateString(), "day", "01-02")
	apiNumChart.Data("访问量", apiNumMaps)
	assign["apiNumChart"] = apiNumChart.Render()

	// 响应速度
	type timeData struct {
		Label string
		Max   string
		Min   string
	}

	apiTimeData := []*timeData{}
	err = core.Db.Model(&model.VisitorApi{}).Select(`MAX(max_time) as max, MAX(min_time) as min, DATE_FORMAT(date, "%Y-%m-%d") as label`).Where("date  >= ?", startTime).Group("date").Find(&apiTimeData).Error
	if err != nil {
		return err
	}

	apiTimeMax := []*statsData{}
	apiTimeMin := []*statsData{}
	for _, datum := range apiTimeData {
		max, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", cast.ToFloat64(datum.Max)), 64)
		apiTimeMax = append(apiTimeMax, &statsData{
			Label: cast.ToString(datum.Label),
			Value: max,
		})
		min, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", cast.ToFloat64(datum.Min)), 64)
		apiTimeMin = append(apiTimeMin, &statsData{
			Label: cast.ToString(datum.Label),
			Value: min,
		})
	}
	assign["apiTime"] = CoverData(apiTimeMax)

	apiTimeMaxMaps := []map[string]any{}
	err = mapstructure.Decode(apiTimeMax, &apiTimeMaxMaps)
	if err != nil {
		return err
	}
	apiTimeMinMaps := []map[string]any{}
	err = mapstructure.Decode(apiTimeMin, &apiTimeMinMaps)
	if err != nil {
		return err
	}
	apiTimeChart := chart.New("date").Legend(true, "right", "top").Line().Date(startTime, carbon.Now().ToDateString(), "day", "01-02")
	apiTimeChart.Data("最大延迟", apiTimeMaxMaps)
	apiTimeChart.Data("最小延迟", apiTimeMinMaps)
	assign["apiTimeChart"] = apiTimeChart.Render()

	// 文件统计
	fileNumData := []*statsData{}
	err = core.Db.Model(&toolsModel.ToolFile{}).Select(`COUNT(*) as value, DATE_FORMAT(created_at, "%Y-%m-%d") as label`).Where("created_at  >= ?", startTime).Group(`DATE_FORMAT(created_at,"%Y-%m-%d")`).Find(&fileNumData).Error
	if err != nil {
		return err
	}
	assign["fileNum"] = CoverData(fileNumData)

	fileNumMaps := []map[string]any{}
	err = mapstructure.Decode(fileNumData, &fileNumMaps)
	if err != nil {
		return err
	}
	fileNumChart := chart.New("date").Legend(true, "right", "top").Column().Date(startTime, carbon.Now().ToDateString(), "day", "01-02")
	fileNumChart.Data("上传量", fileNumMaps)
	assign["fileNumChart"] = fileNumChart.Render()

	// 操作统计
	operateData := []*statsData{}
	err = core.Db.Model(&model.VisitorOperate{}).Select(`COUNT(*) as value, DATE_FORMAT(created_at,"%Y-%m-%d")  as label`).Where("created_at  >= ?", startTime).Where("type = ?", "admin").Group(`DATE_FORMAT(created_at,"%Y-%m-%d")`).Find(&operateData).Error
	if err != nil {
		return err
	}
	assign["operate"] = CoverData(operateData)

	operateMaps := []map[string]any{}
	err = mapstructure.Decode(operateData, &operateMaps)
	if err != nil {
		return err
	}
	operateChart := chart.New("date").Legend(true, "right", "top").Column().Date(startTime, carbon.Now().ToDateString(), "day", "01-02")
	operateChart.Data("日志量", operateMaps)
	assign["operateChart"] = operateChart.Render()

	sqlVersion := []map[string]any{}
	core.Db.Raw(`SHOW VARIABLES LIKE "version"`).Scan(&sqlVersion)

	redisInfo, _ := core.Redis.Info(core.Ctx).Result()

	myRegex, _ := regexp.Compile(`redis_version:(.*)`)
	found := myRegex.FindStringSubmatch(redisInfo)
	redisVer := found[1]
	assign["ver"] = map[string]any{
		"go":   runtime.Version(),
		"echo": echo.Version,
		"dux":  core.Version,
		//"mysql": sqlVersion[0]["Value"],
		"redis": redisVer,
	}

	client := resty.New().SetTimeout(2 * time.Second).R()
	resp, _ := client.Get("http://ip.dhcp.cn/?ip")
	ip := "-"
	if resp.String() != "" {
		ip = resp.String()
	}
	assign["sys"] = map[string]any{
		"os":         runtime.GOOS,
		"ip":         ip,
		"pid":        os.Getpid(),
		"uploadSize": coreConfig.Get("storage").Get("driver.maxSize"),
		"timeout":    coreConfig.Get("app").Get("server.timeout"),
		"freeDisk":   getFreeDisk() / 1024 / 1024 / 1024,
	}

	return response.New(ctx).Render("adminHome.gohtml", assign)
}

func Menu(ctx echo.Context) error {
	if menuData == nil {
		menu := coreRegister.AppMenu["admin"].Render()
		menuData = &menu
	}
	return response.New(ctx).Send("ok", menuData)
}

func Notification(ctx echo.Context) error {
	return response.New(ctx).Send("ok", map[string]any{
		"list": []any{},
		"num":  0,
	})
}

func CoverData(data []*statsData) map[string]any {

	var dataTmpDay float64
	var dataTmpRate int
	var dataTmpBefore float64
	var dataTmpSum float64
	var dataTmpTrend = 1

	if len(data) >= 1 {
		dataTmpLast := data[len(data)-1]
		dataTmpDay = dataTmpLast.Value
		for _, datum := range data {
			dataTmpSum += datum.Value
		}
	}

	if len(data) >= 2 {
		dataTmpBefore = data[len(data)-2].Value
	}
	dataTmpRate = lo.Ternary[int](dataTmpSum > 0, cast.ToInt(dataTmpDay/dataTmpSum*100), 0)

	if dataTmpBefore < dataTmpDay {
		dataTmpTrend = 2
	}
	if dataTmpBefore > dataTmpDay {
		dataTmpTrend = 0
	}
	return map[string]any{
		"day":   dataTmpDay,
		"rate":  dataTmpRate,
		"trend": dataTmpTrend,
	}
}

func getFreeDisk() uint64 {
	var stat syscall.Statfs_t
	wd, _ := os.Getwd() // 获取目录
	syscall.Statfs(wd, &stat)
	// size := fmt.Sprintf("%.f", float64(stat.Bavail * uint64(stat.Bsize)) / 1024 / 1024 / 1024) G
	size := stat.Bavail * uint64(stat.Bsize) // k
	return size
}
