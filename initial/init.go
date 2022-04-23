package initial

import "simple-cicd/router"

func Init() {
	InitChannelQueue()
	InitViper()
	Zap.InitZap()
	InitEtcd()
	InitRouter()
	router.RegisterRouter()
}
