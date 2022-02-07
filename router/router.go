package router

import (
	"github.com/gin-gonic/gin"
	"simple-cicd/controller"
)

var defaultRouter *gin.Engine

func Init() {
	defaultRouter = gin.Default()
	defaultRouter.GET("/test", controller.Test)
	defaultRouter.PUT("/task", controller.CreateTask)
	defaultRouter.Run("0.0.0.0:5080")
}
