package global

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"simple-cicd/model"
	"sync"
)

var (
	EtcdCli      *clientv3.Client
	EtcdCliAlias Etcd
)

type Etcd struct{}

func (q *Etcd) Add(task *model.Task) error {
	var (
		lock sync.Mutex
		kv   clientv3.KV
	)
	kv = clientv3.NewKV(EtcdCli)
	data, err := json.Marshal(task)
	if err != nil {
		zap.L().Error("Etcd.Add Marshal error", zap.String(task.Name, err.Error()))
		return err
	}
	key := fmt.Sprintf("%s/%s", C.EtcdConfig.TaskPath, task.Id)
	zap.L().Info("[Add Task into queue] ", zap.String("key", key), zap.String("value", string(data)))

	defer lock.Unlock()
	lock.Lock()
	if _, err = kv.Put(context.TODO(), key, string(data)); err != nil {
		zap.L().Error("put into etcd failed", zap.Error(err), zap.String("key", key))
		return err
	}
	return nil
}
func (q *Etcd) Delete(taskId string) error {
	var (
		lock sync.Mutex
		err  error
		kv   clientv3.KV
	)
	kv = clientv3.NewKV(EtcdCli)
	key := fmt.Sprintf("%s/%s", C.EtcdConfig.TaskPath, taskId)
	zap.L().Info("[Add Task into queue] ", zap.String("key", key))
	defer lock.Unlock()
	lock.Lock()
	_, err = kv.Delete(context.TODO(), key)
	return err
}
func (q *Etcd) Get(path string) ([]model.Task, error) {
	var (
		err  error
		resp *clientv3.GetResponse
		ret  []model.Task
		kv   clientv3.KV
	)
	kv = clientv3.NewKV(EtcdCli)
	// 先GET到当前的值，并监听后续变化
	if resp, err = kv.Get(context.TODO(), path, clientv3.WithPrefix()); err != nil {
		zap.L().Error(path+" :kv.Get err ", zap.Error(err))
		return nil, err
	}
	// 现在key是存在的
	if len(resp.Kvs) == 0 {
		err = errors.New(path + ":len(ret.Kvs) == 0")
		zap.L().Error(path+" :kv.Get err ", zap.Error(err))
		return nil, err
	}
	for i, _ := range resp.Kvs {
		task := model.Task{}
		err = json.Unmarshal(resp.Kvs[i].Value, &task)
		if err!=nil {
			zap.L().Error(path+" :kv.Get Unmarshal err ", zap.Error(err))
		}
		ret = append(ret, task)
	}
	return ret, nil

}

// Watch 使用cancelFunc控制监听器的退出,这里ctx感知到cancel则会关闭watcher
func (q *Etcd) Watch(path string, ctx context.Context, cancelFunc context.CancelFunc) {
	var (
		err                error
		getResp            *clientv3.GetResponse
		watchStartRevision int64
		kv                 clientv3.KV
	)
	// 先GET到当前的值，并监听后续变化
	if getResp, err = kv.Get(context.TODO(), path, clientv3.WithPrefix()); err != nil {
		zap.L().Error(path+" : kv.Get err ", zap.Error(err))
		return
	}
	// 现在key是存在的
	if len(getResp.Kvs) != 0 {
		zap.L().Info(path+" :当前值:", zap.Any("", getResp.Kvs[0].Value))
	}
	// 当前etcd集群事务ID, 单调递增的（监听path后续的变化,也就是通过监听版本变化）
	watchStartRevision = getResp.Header.Revision + 1
	// 创建一个watcher(监听器)
	watcher := clientv3.NewWatcher(EtcdCli)
	// 启动监听,Watch返回一个channel
	watchRespChan := watcher.Watch(ctx, path, clientv3.WithRev(watchStartRevision))
	//消费channel,for 遍历chan等价于 <-chan
	for watchResp := range watchRespChan {
		for _, event := range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				zap.L().Info(path+"监听到put", zap.String("值为", string(event.Kv.Value)), zap.Int64("lastCreateRevision:", event.Kv.CreateRevision), zap.Int64("lastModRevision:", event.Kv.ModRevision))

			case mvccpb.DELETE:
				fmt.Println("删除了", "Revision:", event.Kv.ModRevision)
			}
		}
	}

}
