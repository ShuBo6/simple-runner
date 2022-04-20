package queue

import "simple-cicd/pkg/model"

var (
	ChannelTaskQueue chan *model.Task
)

func InitChannelQueue() {
	ChannelTaskQueue = make(chan *model.Task, 10)
}
