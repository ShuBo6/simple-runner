package initial

import (
	"simple-cicd/client"
	"simple-cicd/router"
)

func Init() {
	InitViper()
	Zap.InitZap()
	InitEtcd()
	InitChannelQueue()
	InitRouter()

}
func InitService() {
	client.TaskQueue.InitTaskHandler()
	router.RegisterRouter()
}
