package initial

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"simple-cicd/global"
	"time"
)

func InitEtcd() {
	var err error
	global.EtcdCli, err = clientv3.New(clientv3.Config{
		Endpoints:   global.C.EtcdConfig.Endpoints,
		DialTimeout: time.Second,
		Username:    global.C.EtcdConfig.Username,
		Password:    global.C.EtcdConfig.Password,
	})
	if err != nil {
		zap.L().Error("etcd init failed", zap.Error(err))
	}
}
