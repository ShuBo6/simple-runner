package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"simple-cicd/client"
	"simple-cicd/config"
	"simple-cicd/pkg/queue"
	"simple-cicd/pkg/service"
	"simple-cicd/router"
	"sync"
	"syscall"
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
	gin.SetMode(gin.DebugMode)
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.JSONFormatter{})
	//router.Init()
	wg := &sync.WaitGroup{}
	service.EtcdHandler(wg)
	service.Run(wg)
	router.Init()
	wg.Wait()
	go processSignal()
}
func processSignal() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGTSTP)
	sn := <-signalChan
	logrus.Infof("exit server because signal %d", sn)
}
