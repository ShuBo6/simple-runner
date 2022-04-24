package client

import (
	"context"
	"go.uber.org/zap"
	"simple-cicd/global"
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
					//zap.L().Info("[TaskHandler] global.ChannelTaskQueue close")
				} else if task != nil {
					go func() {
						task.UpdateTime = time.Now()
						task.Status = global.Running
						_ = global.EtcdCliAlias.Add(global.C.EtcdConfig.HistoryTaskPath, task)
						// 将任务执行的stdout信息保存到history，然后删除这个task
						args := strings.Split(task.Args, ",")
						out, err := service.ExecShell(task.Cmd, task.EnvMap, args...)
						if err != nil {
							zap.L().Error("[TaskHandler] ExecShell error", zap.Error(err))
							task.TaskData.StdErr = out
							task.TaskData.Error = err.Error()
							task.Status = global.Failed
						} else {
							task.TaskData.Stdout = out
							task.Status = global.Finished
						}
						go func() {
							task.DeleteTime = time.Now()
							err = global.EtcdCliAlias.Delete(task.Id)
							if err != nil {
								zap.L().Error("[TaskHandler] Delete error", zap.Error(err))
							}
						}()
						go func() {
							err = global.EtcdCliAlias.Add(global.C.EtcdConfig.HistoryTaskPath, task)
							if err != nil {
								zap.L().Error("[TaskHandler] Add to history error", zap.Error(err))
							}
						}()
					}()
				}

			case <-t.CloseChan3:
				zap.L().Info("[TaskHandler] exit")
				return

			}

		}
	}()

}
