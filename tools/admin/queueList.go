package admin

import (
	"fmt"
	"github.com/duxphp/duxgo-admin/system/service"
	"github.com/duxphp/duxgo-ui/lib/form"
	"github.com/duxphp/duxgo-ui/lib/table"
	"github.com/duxphp/duxgo-ui/lib/table/column"
	"github.com/duxphp/duxgo/core"
	"github.com/duxphp/duxgo/exception"
	"github.com/hibiken/asynq"
	"github.com/jianfengye/collection"
	"github.com/labstack/echo/v4"
)

func QueueList(ctx echo.Context) error {
	if core.QueueInspector == nil {
		return exception.BusinessError("队列服务未开启，暂时无法使用")
	}
	return service.NewManageExpand().SetTable(queueTable).ListPage(ctx)
}

func QueueAjax(ctx echo.Context) error {
	return service.NewManageExpand().SetTable(queueTable).ListData(ctx)
}

func queueTable(ctx echo.Context) *table.Table {
	tableT := table.NewTable()

	tableT.SetUrl("/admin/tools/queueList/ajax")

	tableT.AddTab("处理中")
	tableT.AddTab("待处理")
	tableT.AddTab("定时任务")
	tableT.AddTab("重试任务")
	tableT.AddTab("存档任务")

	queueList, _ := core.QueueInspector.Queues()
	options := map[any]any{}
	for _, name := range queueList {
		options[name] = name
	}
	tableT.AddFilter("队列类型", "queue").SetUI(form.NewSelect().SetOptions(options)).SetQuick(true).SetDefault("default")

	tableT.SetDataFun(func(filter map[string]any) (collect collection.ICollection) {
		fmt.Println("test0")
		qname := filter["queue"].(string)
		var tasks []*asynq.TaskInfo
		if filter["tab"] == 0 {
			tasks, _ = core.QueueInspector.ListActiveTasks(qname)
		}
		if filter["tab"] == 1 {
			tasks, _ = core.QueueInspector.ListPendingTasks(qname)
		}
		if filter["tab"] == 2 {
			tasks, _ = core.QueueInspector.ListScheduledTasks(qname)
		}
		if filter["tab"] == 3 {
			tasks, _ = core.QueueInspector.ListRetryTasks(qname)
		}
		if filter["tab"] == 4 {
			tasks, _ = core.QueueInspector.ListArchivedTasks(qname)
		}
		fmt.Println("test1")

		type taskI struct {
			Id      string
			Type    string
			Payload string
			Time    string
		}

		fmt.Println("test2")
		var data []taskI
		var date string

		for _, task := range tasks {
			if !task.NextProcessAt.IsZero() {
				date = task.NextProcessAt.Format("2006-01-02 15:04:05")
			}
			if !task.CompletedAt.IsZero() {
				date = task.CompletedAt.Format("2006-01-02 15:04:05")
			}
			if date == "" {
				date = "-"
			}
			data = append(data, taskI{
				Id:      task.ID,
				Type:    task.Type,
				Payload: string(task.Payload),
				Time:    date,
			})
		}
		fmt.Println("test3")

		return collection.NewObjCollection(data)
	}, "id")

	tableT.AddCol("ID", "FamilyId").SetUI(column.NewContext())
	tableT.AddCol("类型", "Type").SetUI(column.NewContext())
	tableT.AddCol("参数", "Payload").SetUI(column.NewContext())

	tableT.AddCol("时间", "Time").SetUI(column.NewContext()).SetWidth(220)

	return tableT
}
