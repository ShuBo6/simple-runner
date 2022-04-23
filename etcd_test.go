package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	"simple-cicd/client"
	"simple-cicd/global"
	"simple-cicd/queue"
	"simple-cicd/router"
	"testing"
)

func TestName(t *testing.T) {

	client.Init()
	err := global.Load("conf/config.yaml")
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
	pop, err := queue.GetTaskQueue().Pop(context.Background(), "1644321561")
	if err != nil {
		return
	}
	fmt.Println(pop)
}
