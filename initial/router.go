package initial

import (
	"github.com/gin-gonic/gin"
	"simple-cicd/api"
	"simple-cicd/model/response"
)

var defaultRouter *gin.Engine
var handler=response.Handler{}
func InitRouter() {
	defaultRouter = gin.Default()
	defaultRouter.GET("/health", api.Health)
	defaultRouter.POST("/task", handler.Handler()(api.CreateTask))
	defaultRouter.GET("/task", handler.Handler()(api.ListTask))

	defaultRouter.Run("0.0.0.0:5080")
}