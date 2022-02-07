package main

import (
	"simple-cicd/config"
	"simple-cicd/pkg/queue"
	"simple-cicd/router"
)

func main() {

	err := config.Load("conf/config.yaml")
	if err != nil {
		return
	}
	err = queue.InitEtcdQueue()
	if err != nil {
		return
	}
	router.Init()
	//cli, _ :=
	//get, err := cli.Get(cli.Ctx(),"a")
	//if err != nil {
	//	return
	//}
	//fmt.Println(client.TaskQueue.Pop())
}
