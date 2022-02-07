package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-cicd/pkg/model"
	"simple-cicd/pkg/queue"
	"time"
)

func CreateTask(ctx *gin.Context) {
	err := queue.GetTaskQueue().Add(ctx, &model.Task{
		Id: fmt.Sprintf("%d", time.Now().Unix()),
	})
	if err != nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "create task failed.",
			"err":     err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "ok",
		"err":     "",
	})
}
