package service

import (
	"context"
	"fmt"
	"github.com/prometheus/common/log"
	"os"
	"os/exec"
	"simple-cicd/pkg/model"
	"simple-cicd/pkg/queue"
)

func Exec(task *model.Task) {
	log.Infof("[executor] task: %+v starting.", *task)
	c := exec.Command("/bin/sh", "-c", "-xe", task.Data.Cmd)
	//c := exec.Command( task.Data.Cmd)
	c.Env = os.Environ()
	for k, v := range task.Data.EnvMap {
		c.Env = append(c.Env, fmt.Sprintf("%s=%s", k, v))
	}

	output, err := c.CombinedOutput()
	if err != nil {
		log.Errorf("[Exec] err:%+v", err.Error())
		task.Data.Stdout = err.Error()
		task.Status = 3
	} else {
		log.Infof(string(output))
		task.Data.Stdout = string(output)
		task.Status = 2
	}
	err = queue.GetStartTaskQueue().Add(context.Background(), task)
	if err != nil {
		log.Errorf("[EtcdHandler] save StartTaskQueue failed,err:%+v", err)
		return
	}
	log.Infof("[executor] task: %s finished.", task.Name)
}
