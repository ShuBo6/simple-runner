package service

import (
	"context"
	"github.com/prometheus/common/log"
	"simple-cicd/client"
	"simple-cicd/pkg/queue"
)

func EtcdHandler() {
	go func() {
		for {
			q := queue.GetTaskQueue()
			task, err := q.Pop(context.Background(), q.RootPath)
			if err != nil {
				continue
			}
			log.Debugf("[EtcdHandler] get task:%+v", *task)
			client.TaskChan <- task
		}

	}()
}
