package global

import clientv3 "go.etcd.io/etcd/client/v3"

var (
	EtcdCli *clientv3.Client
)

