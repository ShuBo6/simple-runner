package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	"os"
	"simple-cicd/client"
	"simple-cicd/config"
	"simple-cicd/pkg/queue"
	"simple-cicd/pkg/service"
	"simple-cicd/router"
	"sync"
)

func main() {

	client.Init()
	err := config.Load("conf/config.yaml")
	if err != nil {
		log.Error("load config path(conf/config.yaml) failed.")
		return
	}
	err = queue.InitEtcdQueue()
	if err != nil {
		log.Errorf("InitEtcdQueue failed.err:%s", err.Error())
		return
	}
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	queue.InitChannelQueue()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Infof("[Run] start Run worker")
		fmt.Println("[Run] start Run worker")
		for {
			select {
			case task, ok := <-queue.ChannelTaskQueue:
				if ok {
					go service.Exec(task)
				}

			}
		}
	}()
	router.Init()
	wg.Wait()

}
