package client

import (
	"simple-cicd/model"
)

//
var (
	TaskChan chan *model.Task
)

//
func Init() {
	TaskChan = make(chan *model.Task, 999)
}
