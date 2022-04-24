package initial

import (
	"simple-cicd/client"
	"simple-cicd/global"
	"simple-cicd/model"
)

type ChannelQueue struct {
	closeChan <-chan bool
}

func InitChannelQueue() {
	global.ChannelTaskQueue = make(chan *model.Task, 10)
	client.TaskQueue = client.NewTaskQueue()
}
