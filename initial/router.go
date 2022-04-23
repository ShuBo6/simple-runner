package initial

import (
	"github.com/gin-gonic/gin"
	"simple-cicd/global"
	"simple-cicd/router"
)




func InitRouter() {
	global.DefaultRouter = gin.Default()
	global.DefaultRouter.Use(router.ZapLogger())
}
