package model

import (
	"encoding/json"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type TaskQueue struct {
	TaskQueue chan *Task
	EtcdCli   *clientv3.Client
}

func (q *TaskQueue) Pop(path string) *Task {

	clientv3.OpGet(path).ValueBytes()
	return nil
}

func (q *TaskQueue) Add(task *Task, path string) {
	data, err := json.Marshal(task)
	if err != nil {
		return
	}
	clientv3.OpPut(path, string(data))
}
