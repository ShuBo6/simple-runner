package main

import (
	"fmt"
	"simple-cicd/initial"
	"simple-cicd/utils"
	"testing"
)

func TestSendMail(t *testing.T) {
	initial.Init()
	fmt.Println(utils.SendMail([]string{"814183583@qq.com"}, "test", "test"))
}
