package service

import "simple-cicd/client"

func Run() {
	go func() {
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
