package model

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

type TaskQueue struct {
	TaskQueue chan *Task
	cli       *clientv3.Client
	locker    sync.Locker
}

func NewQueue() (*TaskQueue, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"http://192.168.123.24:2379"},
	})
	if err != nil {
		logrus.Warningf("get etcd client when new etcd queue failed, error:%s", err.Error())
		return nil, err
	}
	queue := &TaskQueue{
		TaskQueue: make(chan *Task, 999),
		cli:       cli,
		locker:    &sync.Mutex{},
	}
	return queue, err
}
func (q *TaskQueue) Pop(ctx context.Context, taskId string) (*Task, error) {

	q.locker.Lock()
	defer q.locker.Unlock()
	prefix := fmt.Sprintf(ETCDPATH, taskId)
	resp, err := q.cli.Get(q.cli.Ctx(), prefix, clientv3.WithFirstKey()...)
	if err != nil {
		logrus.Warningf("get item:%s from etcd queue failed, error:%s", prefix, err.Error())
		return nil, err
	}
	if len(resp.Kvs) == 1 {
		_, err = q.cli.Delete(ctx, string(resp.Kvs[0].Key))
		if err != nil {
			return nil, err
		}
		item := &Task{}
		if err = json.Unmarshal(resp.Kvs[0].Value, item); err != nil {
			return nil, err
		}
		return item, nil
	}
	return nil, errors.New("empty ")
}

func (q *TaskQueue) Add(ctx context.Context, task *Task, taskId string) error {
	q.TaskQueue <- task
	data, err := json.Marshal(task)
	if err != nil {
		logrus.Warningf("marshal queue item:%s failed, error:%s", task.Name, err.Error())
		return err
	}
	key := fmt.Sprintf(ETCDPATH, taskId)
	q.locker.Lock()
	defer q.locker.Unlock()
	if _, err = q.cli.Put(ctx, key, string(data)); err != nil {
		logrus.Warningf("put item:%s to etcd queue failed, error:%s", key, err.Error())
		return err
	}
	return nil
}
