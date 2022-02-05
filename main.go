package main

import (
	"simple-cicd/client"
	"simple-cicd/router"
)

func main() {
	router.Init()
	client.Init()
	//cli, _ :=
	//get, err := cli.Get(cli.Ctx(),"a")
	//if err != nil {
	//	return
	//}
	//fmt.Println(client.TaskQueue.Pop())
}
