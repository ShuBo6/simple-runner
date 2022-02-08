package client

import (
	clientv3 "go.etcd.io/etcd/client/v3"
	"simple-cicd/config"
	"simple-cicd/pkg/model"
	"time"
)

//
var (
	TaskChan chan *model.Task
)

//
func Init() {
	TaskChan = make(chan *model.Task, 999)
}

func NewEtcdClient() (*clientv3.Client, error) {
	conf := config.C.EtcdConfig
	return clientv3.New(clientv3.Config{
		Endpoints:   conf.Endpoints,
		DialTimeout: time.Second,
		Username:    conf.Username,
		Password:    conf.Password,
	})
}
