package service

import (
	"github.com/prometheus/common/log"
	"simple-cicd/client"
)

func Run() {
	log.Infof("[Run] start Run worker")
	for {
		select {
		case task, ok := <-client.TaskChan:
			if ok {
				go Exec(task)
			}

		}
	}

}
