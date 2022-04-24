package router

import (
	"simple-cicd/api"
	"simple-cicd/global"
	"simple-cicd/model/response"
)

var handler = response.Handler{}

func RegisterRouter() {
	global.DefaultRouter.GET("/health", api.Health)
	global.DefaultRouter.POST("/task", handler.Handler()(api.CreateTask))
	global.DefaultRouter.GET("/task", handler.Handler()(api.ListTask))

	_ = global.DefaultRouter.Run("0.0.0.0:5080")

}
