package main

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	//router.Init()
	cli, _ := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://192.168.123.24:2379"},
	})
	get, err := cli.Get(cli.Ctx(),"a")
	if err != nil {
		return
	}
	fmt.Println(get.Kvs)
}
