package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"simple-cicd/global"
	"simple-cicd/model"
	"simple-cicd/model/request"
	"simple-cicd/model/response"
	"time"
)

func CreateTask(ctx *gin.Context) *response.Response {
	var (
		err error
		v   []byte
		msg string = "CreateTask success"
	)
	var req request.TaskRequest
	err = ctx.ShouldBind(&req)
	if err != nil {
		return &response.Response{Code: response.ERROR, Message: err.Error()}
	}
	envMap := make(map[string]string)
	if req.Env != "" {
		err = json.Unmarshal([]byte(req.Env), &envMap)
		if err != nil {
			return &response.Response{Code: response.ERROR, Message: err.Error()}
		}
	}
	task := &model.Task{
		Id:         uuid.New().String(),
		Name:       req.Name,
		Type:       model.TaskType(req.Type),
		Cmd:        req.Cmd,
		Args:       req.Args,
		EnvMap:     envMap,
		CreateTime: time.Now(),
		Status:     global.Ready,
	}
	if err != nil {
		return &response.Response{Code: response.ERROR, Message: err.Error()}
	}
	err = global.EtcdCliAlias.Add(global.C.EtcdConfig.TaskPath, task)
	if err != nil {
		return &response.Response{Code: response.ERROR, Message: err.Error()}
	}
	return &response.Response{Code: response.SUCCESS, Data: v, Message: msg}
}
func ListTask(ctx *gin.Context) *response.Response {
	var (
		err error
		ret []model.Task
	)

	var req request.ListTaskRequest
	err = ctx.ShouldBind(&req)
	if err != nil {
		return &response.Response{Code: response.ERROR, Message: err.Error()}
	}
	// 0 查所有，1查未开始，2查历史
	switch req.Scope {
	case 0:
		ret, _ = global.EtcdCliAlias.Get(global.C.EtcdConfig.TaskPath)
		r2, _ := global.EtcdCliAlias.Get(global.C.EtcdConfig.HistoryTaskPath)
		ret = append(ret, r2...)
	case 1:
		ret, _ = global.EtcdCliAlias.Get(global.C.EtcdConfig.TaskPath)
	case 2:
		ret, _ = global.EtcdCliAlias.Get(global.C.EtcdConfig.HistoryTaskPath)
	}

	return &response.Response{Code: response.SUCCESS, Data: ret}
}
