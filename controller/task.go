package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"io/ioutil"
	"net/http"
	"simple-cicd/pkg/model"
	"simple-cicd/pkg/queue"
	"time"
)

func CreateTask(ctx *gin.Context) {

	task := &model.Task{}

	reqBodyData, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "get reqBodyData err",
			"err":     err.Error(),
		})
		return
	}
	err = json.Unmarshal(reqBodyData, task)
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "Unmarshal reqBodyData err",
			"err":     err.Error(),
		})
		return
	}
	task.Id = fmt.Sprintf("%d", time.Now().Unix())
	queue.ChannelTaskQueue <- task
	//err = queue.GetTaskQueue().Add(context.Background(), task)
	//if err != nil {
	//	ctx.JSON(http.StatusOK, map[string]string{
	//		"message": "create task failed.",
	//		"err":     err.Error(),
	//	})
	//	return
	//}
	log.Infof("[CreateTask] task:%+v", *task)
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "ok",
		"err":     "",
	})
}
