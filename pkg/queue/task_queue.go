package queue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"simple-cicd/client"
	"simple-cicd/config"
	"simple-cicd/pkg/model"
	"sync"
)

var gTaskEtcdQueue *TaskQueue
var gStartedTaskEtcdQueue *TaskQueue

const (
	TASKPATH = "tasks"
)

type TaskQueue struct {
	RootPath string
	cli      *clientv3.Client
	locker   sync.Locker
}

func InitEtcdQueue() error {
	var err error
	gTaskEtcdQueue, err = NewQueue("tasks")
	gStartedTaskEtcdQueue, err = NewQueue("started-tasks")
	if err != nil {
		return err
	}
	return nil
}
func GetTaskQueue() *TaskQueue {
	return gTaskEtcdQueue
}
func GetStartTaskQueue() *TaskQueue {
	return gStartedTaskEtcdQueue
}
func NewQueue(key string) (*TaskQueue, error) {
	cli, err := client.NewEtcdClient()
	if err != nil {
		logrus.Warningf("get etcd client when new etcd queue failed, error:%s", err.Error())
		return nil, err
	}
	queue := &TaskQueue{
		RootPath: fmt.Sprintf("%s%s/", config.C.EtcdConfig.RootPath, key),
		cli:      cli,
		locker:   &sync.Mutex{},
	}
	return queue, err
}

func (q *TaskQueue) PopAll(ctx context.Context) (*model.Task, error) {

	q.locker.Lock()
	defer q.locker.Unlock()
	prefix := q.RootPath
	resp, err := q.cli.Get(ctx, prefix, clientv3.WithFirstKey()...)
	if err != nil {
		logrus.Warningf("get item:%s from etcd queue failed, error:%s", prefix, err.Error())
		return nil, err
	}
	if len(resp.Kvs) == 1 {
		_, err = q.cli.Delete(ctx, string(resp.Kvs[0].Key))
		if err != nil {
			return nil, err
		}
		item := &model.Task{}
		if err = json.Unmarshal(resp.Kvs[0].Value, item); err != nil {
			log.Errorf("[PopAll] err:%+v", err.Error())
			return nil, err
		}
		//item.Status=1
		//marshal, err := json.Marshal(item)
		//if err != nil {
		//	log.Errorf("[PopAll] err:%+v",err.Error())
		//	return nil, err
		//}
		//_, err = q.cli.Put(ctx, string(resp.Kvs[0].Key), string(marshal))
		//if err != nil {
		//	log.Errorf("[PopAll] err:%+v",err.Error())
		//	return nil, err
		//}
		return item, nil
	}
	return nil, errors.New("empty ")
}
func (q *TaskQueue) Pop(ctx context.Context, taskId string) (*model.Task, error) {

	q.locker.Lock()
	defer q.locker.Unlock()
	prefix := q.RootPath + taskId
	resp, err := q.cli.Get(ctx, prefix, clientv3.WithFirstKey()...)
	if err != nil {
		logrus.Warningf("get item:%s from etcd queue failed, error:%s", prefix, err.Error())
		return nil, err
	}
	if len(resp.Kvs) == 1 {
		_, err = q.cli.Delete(ctx, string(resp.Kvs[0].Key))
		if err != nil {
			return nil, err
		}
		item := &model.Task{}
		if err = json.Unmarshal(resp.Kvs[0].Value, item); err != nil {
			return nil, err
		}
		return item, nil
	}
	return nil, errors.New("empty ")
}

func (q *TaskQueue) Add(ctx context.Context, task *model.Task) error {
	data, err := json.Marshal(task)
	if err != nil {
		logrus.Warningf("marshal queue item:%s failed, error:%s", task.Name, err.Error())
		return err
	}
	key := q.RootPath + task.Id
	log.Debugf("[Add Task into queue] key:%s,value:%+v", key, *task)
	q.locker.Lock()
	defer q.locker.Unlock()
	if _, err = q.cli.Put(ctx, key, string(data)); err != nil {
		logrus.Warningf("put item:%s to etcd queue failed, error:%s", key, err.Error())
		return err
	}
	return nil
}
