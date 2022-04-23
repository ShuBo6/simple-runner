package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"simple-cicd/global"
	"simple-cicd/initial"
	"simple-cicd/model"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	ret := make([]model.Task, 3)
	b,_:=json.Marshal(ret)
	fmt.Println(string(b))
[{"id":"3959a9f8-faa1-4861-a695-b9f153b5ca7f","cmd":"go version","env_map":{},"name":"test12","type":"shell"}{"id":"90c3e7ca-812c-47bb-86aa-7ae8edd981c5","cmd":"go version","env_map":{},"name":"test12","type":"shell"}{"id":"a42fb315-c0dd-488b-a9d0-a2ea256b4cbc","cmd":"go version","env_map":{},"name":"test12","type":"shell"}]
}

func TestWatch(t *testing.T) {

	initial.Init()

	var (
		err                error
		kv                 clientv3.KV
		watcher            clientv3.Watcher
		getResp            *clientv3.GetResponse
		watchStartRevision int64
		watchRespChan      <-chan clientv3.WatchResponse
		watchResp          clientv3.WatchResponse
		event              *clientv3.Event
	)

	client:=global.EtcdCli

	// KV
	kv = clientv3.NewKV(client)

	// 模拟etcd中KV的变化
	go func() {
		for {
			kv.Put(context.TODO(), "/cron/jobs/job7", "i am job7")

			kv.Delete(context.TODO(), "/cron/jobs/job7")

			time.Sleep(1 * time.Second)
		}
	}()

	// 先GET到当前的值，并监听后续变化
	if getResp, err = kv.Get(context.TODO(), "/cron/jobs/job7"); err != nil {
		fmt.Println(err)
		return
	}

	// 现在key是存在的
	if len(getResp.Kvs) != 0 {
		fmt.Println("当前值:", string(getResp.Kvs[0].Value))
	}

	// 当前etcd集群事务ID, 单调递增的（监听/cron/jobs/job7后续的变化,也就是通过监听版本变化）
	watchStartRevision = getResp.Header.Revision + 1

	// 创建一个watcher(监听器)
	watcher = clientv3.NewWatcher(client)

	// 启动监听
	fmt.Println("从该版本向后监听:", watchStartRevision)

	ctx, cancelFunc := context.WithCancel(context.TODO())
	//5秒钟后取消
	time.AfterFunc(5*time.Second, func() {
		cancelFunc()
	})
	//这里ctx感知到cancel则会关闭watcher
	watchRespChan = watcher.Watch(ctx, "/cron/jobs/job7", clientv3.WithRev(watchStartRevision))

	// 处理kv变化事件
	for watchResp = range watchRespChan {
		for _, event = range watchResp.Events {
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("修改为:", string(event.Kv.Value), "Revision:", event.Kv.CreateRevision, event.Kv.ModRevision)
			case mvccpb.DELETE:
				fmt.Println("删除了", "Revision:", event.Kv.ModRevision)
			}
		}
	}

}
