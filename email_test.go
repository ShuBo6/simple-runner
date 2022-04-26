package main

import (
	"encoding/json"
	"fmt"
	"simple-cicd/initial"
	"simple-cicd/model"
	"simple-cicd/model/request"
	"simple-cicd/utils"
	"testing"
)

func TestSendMail(t *testing.T) {
	initial.Init()
	fmt.Println(utils.SendMail([]string{"814183583@qq.com"}, "test", "test"))
}
func TestPrintStruct(t *testing.T) {
	a, _ := json.Marshal(request.TaskRequest{})
	b, _ := json.Marshal(model.PipelineTaskMetaData{})
	c, _ := json.Marshal(model.Notify{})
	fmt.Println(string(a))
	fmt.Println(string(b))
	fmt.Println(string(c))
}
