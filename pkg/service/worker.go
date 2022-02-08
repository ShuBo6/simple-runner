package service

import (
	"github.com/prometheus/common/log"
	"simple-cicd/client"
	"sync"
)

func Run(wg *sync.WaitGroup) {
	go func() {
		log.Infof("[Run] start Run worker")
		defer wg.Done()
		wg.Add(1)
		for {
			select {
			case task, ok := <-client.TaskChan:
				if ok {
					Exec(task)
				}

			}
		}
	}()

}
