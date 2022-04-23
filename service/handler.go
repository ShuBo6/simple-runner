package service

//func EtcdHandler(wg *sync.WaitGroup) {
//	go func() {
//		ctx := context.Background()
//		defer wg.Done()
//		log.Infof("[EtcdHandler] start EtcdHandler")
//		wg.Add(1)
//		for {
//			time.Sleep(2 * time.Second)
//			q := queue.GetTaskQueue()
//			task, err := q.PopAll(ctx)
//			if task != nil{
//				if err != nil {
//					log.Warnf("[EtcdHandler] pop taskQueue failed,err:%+v", err)
//					continue
//				}
//				task.Status = 1
//				err = queue.GetStartTaskQueue().Add(ctx, task)
//				if err != nil {
//					log.Errorf("[EtcdHandler] save StartTaskQueue failed,err:%+v", err)
//					return
//				}
//				log.Infof("[EtcdHandler] get task:%+v", *task)
//				client.TaskChan <- task
//			}
//
//		}
//	}()
//}
