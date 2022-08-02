package admin

import (
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/pkg/chart"
	"github.com/duxphp/duxgo/response"
	"github.com/duxphp/duxgo/util/function"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

func QueueMain(ctx echo.Context) error {

	var data []map[string]any
	queueCharts := chart.New("custom").Legend(true, "right", "top").Column(true)
	queueList, _ := core.QueueInspector.Queues()

	var activeData []map[string]any
	var pendingData []map[string]any
	var scheduledData []map[string]any
	var retryData []map[string]any
	var archivedData []map[string]any

	for _, name := range queueList {
		queueInfo, _ := core.QueueInspector.GetQueueInfo(name)
		if queueInfo == nil {
			continue
		}
		var m map[string]any
		mapstructure.Decode(queueInfo, &m)
		m["MemoryUsage"] = function.FormatFileSize(queueInfo.MemoryUsage)
		data = append(data, m)

		activeData = append(activeData, map[string]any{
			"label": queueInfo.Queue,
			"value": queueInfo.Active,
		})

		pendingData = append(pendingData, map[string]any{
			"label": queueInfo.Queue,
			"value": queueInfo.Pending,
		})

		scheduledData = append(scheduledData, map[string]any{
			"label": queueInfo.Queue,
			"value": queueInfo.Scheduled,
		})

		retryData = append(retryData, map[string]any{
			"label": queueInfo.Queue,
			"value": queueInfo.Retry,
		})

		archivedData = append(archivedData, map[string]any{
			"label": queueInfo.Queue,
			"value": queueInfo.Archived,
		})
	}
	queueCharts.Data("执行中", activeData)
	queueCharts.Data("待处理", pendingData)
	queueCharts.Data("定时任务", scheduledData)
	queueCharts.Data("重试任务", retryData)
	queueCharts.Data("存档任务", archivedData)

	taskChart := chart.New("date").Legend(true, "right", "top").Line()
	taskData, _ := core.QueueInspector.History("default", 7)

	var successData []map[string]any
	var failureData []map[string]any
	for _, datum := range taskData {
		successData = append(successData, map[string]any{
			"label": datum.Date.Format("2006-01-02"),
			"value": datum.Processed,
		})
		failureData = append(failureData, map[string]any{
			"label": datum.Date.Format("2006-01-02"),
			"value": datum.Failed,
		})
	}

	taskChart.Data("成功任务", successData)
	taskChart.Data("失败任务", failureData)

	return response.New(ctx).Render("adminToolsQueue.tpl", map[string]any{
		"QueueList":  data,
		"QueueChart": queueCharts.Render(),
		"TaskChart":  taskChart.Render(),
	})
}
