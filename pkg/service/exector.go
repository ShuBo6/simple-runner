package service

import (
	"fmt"
	"github.com/prometheus/common/log"
	"os"
	"os/exec"
	"simple-cicd/pkg/model"
)

func Exec(task *model.Task) {
	log.Debugf("[executor] task: %s starting.")
	c := exec.Command("/usr/bin/sh", "-c", "-e", task.Data.Cmd)
	c.Env = os.Environ()
	for k, v := range task.Data.EnvMap {
		c.Env = append(c.Env, fmt.Sprintf("%s=%s", k, v))
	}

	output, err := c.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	log.Debugf(string(output))
	task.Data.Stdout = string(output)
	log.Debugf("[executor] task: %s finished.")
}
