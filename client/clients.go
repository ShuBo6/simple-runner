package client

import "simple-cicd/pkg/model"

var (
	TaskQueue *model.TaskQueue
)

func Init() {
	TaskQueue, _ = model.NewQueue()
}
