package initial

func Init() {
	InitChannelQueue()
	InitViper()
	Zap.InitZap()
	InitEtcd()
	InitRouter()
}
