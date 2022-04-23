package initial

import (
	"simple-cicd/global"
	"simple-cicd/model"
)

func InitChannelQueue() {
	global.ChannelTaskQueue = make(chan *model.Task, 10)
}
