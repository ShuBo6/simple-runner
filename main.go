package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	"simple-cicd/client"
	"simple-cicd/config"
	"simple-cicd/pkg/queue"
	"simple-cicd/pkg/service"
	"simple-cicd/router"
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
	router.Init()
	gin.SetMode(gin.DebugMode)
	logrus.SetLevel(logrus.TraceLevel)
	service.EtcdHandler()
	service.Run()
}
