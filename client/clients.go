package client

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"simple-cicd/config"
	"time"
)

//
//var (
//	TaskQueue *model.TaskQueue
//)
//
//func Init() {
//	err :=
//	if err != nil {
//		return
//	}
//}

func NewEtcdClient() (*clientv3.Client, error) {
	conf := config.C.EtcdConfig
	return clientv3.New(clientv3.Config{
		Endpoints:   conf.Endpoints,
		DialTimeout: time.Second,
		Username:    conf.Username,
		Password:    conf.Password,
	})
}
