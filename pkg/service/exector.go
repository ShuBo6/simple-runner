package service

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"simple-cicd/pkg/model"
)

func exector(task *model.Task) {
	taskData := new(model.TaskData)
	err := json.Unmarshal([]byte(task.Data), taskData)
	if err != nil {
		return
	}
	c := exec.Command("/usr/bin/sh", "-c", "-e", taskData.Cmd)
	c.Env = os.Environ()
	for k, v := range taskData.EnvMap {
		c.Env = append(c.Env, fmt.Sprintf("%s=%s", k, v))
	}

	output, err := c.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(output))
	taskData.Stdout = string(output)
	taskDataMarshal, err := json.Marshal(taskData)
	if err != nil {
		return
	}
	task.Data = string(taskDataMarshal)
}
