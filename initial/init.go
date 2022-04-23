package initial

import "simple-cicd/router"

func Init() {
	InitChannelQueue()
	InitViper()
	router.Init()
}