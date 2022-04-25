package client

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.uber.org/zap"
	"simple-cicd/global"
	"simple-cicd/model"
	"simple-cicd/service"
	"strings"
	"time"
)

var TaskQueue *TaskQ

type TaskQ struct {
	CloseChan1, CloseChan2, CloseChan3 chan bool
	Context1, Context2                 context.Context
	CancelFunc1, CancelFunc2           context.CancelFunc
}

func NewTaskQueue() *TaskQ {
	c1, cf1 := context.WithCancel(context.Background())
	c2, cf2 := context.WithCancel(context.Background())
	return &TaskQ{
		CloseChan1: make(chan bool),
		CloseChan2: make(chan bool),
		CloseChan3: make(chan bool),
		Context1:   c1, Context2: c2,
		CancelFunc1: cf1,
		CancelFunc2: cf2,
	}
}
func (t *TaskQ) Close() {
	t.CancelFunc1()
	t.CancelFunc2()

	t.CloseChan1 <- true
	t.CloseChan2 <- true
	t.CloseChan3 <- true

}
func (t *TaskQ) InitTaskHandler() {
	//defer close(global.ChannelTaskQueue)

	go func() {
		global.EtcdCliAlias.Watch(global.C.EtcdConfig.TaskPath, t.Context1, t.CancelFunc1)
		<-t.CloseChan1
	}()

	go func() {
		global.EtcdCliAlias.Watch(global.C.EtcdConfig.HistoryTaskPath, t.Context2, t.CancelFunc2)
		<-t.CloseChan2
	}()
	go func() {
		zap.L().Info("[TaskHandler] start")
		for {
			select {
			case task, ok := <-global.ChannelTaskQueue:
				if !ok {
					zap.L().Debug("[TaskHandler] global.ChannelTaskQueue wait")
				} else if task != nil {
					go processTask(task)
				}

			case <-t.CloseChan3:
				zap.L().Info("[TaskHandler] exit")
				return

			}

		}
	}()

}
func processTask(task *model.Task) {

	var (
		out string
		err error
	)

	task.UpdateTime = time.Now()
	task.Status = global.Running
	//更新任务状态
	_ = global.EtcdCliAlias.Add(global.C.EtcdConfig.TaskPath, task)
	args := strings.Split(task.Args, ",")
	switch task.Type {
	case global.ShellTask:

		metaData := model.ShellTaskMetaData{}
		err = json.Unmarshal([]byte(task.TaskMetaData), &metaData)
		if err != nil {
			zap.L().Error("[processTask] json.Unmarshal failed")
			return
		}
		out, err = service.ExecShell(metaData.Cmd, task.EnvMap, args...)
	case global.DockerTask:
		metaData := model.DockerTaskMetaData{}
		err = json.Unmarshal([]byte(task.TaskMetaData), &metaData)
		if err != nil {
			zap.L().Error("[processTask] json.Unmarshal failed")
			return
		}
		out, err = service.GitClone(metaData.GitUrl, metaData.GitRef, task.EnvMap)
		if err == nil {
			out, err = service.DockerBuild(fmt.Sprintf("%s:%s", metaData.ImageName, metaData.Tag),
				metaData.Path, task.EnvMap)
		}

	default:
		zap.L().Error("[processTask] ExecShell error", zap.Error(errors.New(string("未知的类型"+task.Type))))
		err = errors.New(string("未知的类型" + task.Type))
	}

	if err != nil {
		zap.L().Error("[processTask] ExecShell error", zap.Error(err))
		task.TaskData.StdErr = out
		task.TaskData.Error = err.Error()
		task.Status = global.Failed
	} else {
		task.TaskData.Stdout = out
		task.Status = global.Finished
	}
	// 将任务执行的stdout信息保存到history，然后删除这个task
	go func() {
		task.DeleteTime = time.Now()
		err = global.EtcdCliAlias.Delete(task.Id)
		if err != nil {
			zap.L().Error("[processTask] Delete error", zap.Error(err))
		}
	}()
	go func() {
		err = global.EtcdCliAlias.Add(global.C.EtcdConfig.HistoryTaskPath, task)
		if err != nil {
			zap.L().Error("[processTask] Add to history error", zap.Error(err))
		}
	}()

}
