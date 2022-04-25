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
		err     error
		v       []byte
		msg     string = "CreateTask success"
		envMap         = make(map[string]string)
		task           = &model.Task{}
		marshal []byte
	)
	var req request.TaskRequest
	err = ctx.ShouldBind(&req)
	if err != nil {
		return &response.Response{Code: response.ERROR, Message: "type 不能为空"}
	}
	if req.Env != "" {
		err = json.Unmarshal([]byte(req.Env), &envMap)
		if err != nil {
			return &response.Response{Code: response.ERROR, Message: err.Error()}
		}
	}
	task.BaseTask = model.BaseTask{
		Id:         uuid.New().String(),
		Name:       req.Name,
		Type:       global.ShellTask,
		Args:       req.Args,
		EnvMap:     envMap,
		CreateTime: time.Now(),
		Status:     global.Ready,
	}
	switch req.Type {
	case global.ShellTask:
		if req.ShellBuild == nil {
			return &response.Response{Code: response.ERROR, Message: "ShellBuild 不能为空"}
		}
		marshal, err = json.Marshal(req.ShellBuild)
		if err != nil {
			return &response.Response{Code: response.ERROR, Message: err.Error()}
		}

	case global.DockerTask:
		if req.DockerBuild == nil {
			return &response.Response{Code: response.ERROR, Message: "DockerBuild 不能为空"}
		}
		marshal, err = json.Marshal(req.DockerBuild)
		if err != nil {
			return &response.Response{Code: response.ERROR, Message: err.Error()}
		}
		//todo PipelineTask
	case global.PipelineTask:
		return &response.Response{Code: response.ERROR, Message: "PipelineTask 暂时不支持"}
	default:
		return &response.Response{Code: response.ERROR, Message: "未知的类型"}
	}
	task.TaskMetaData = string(marshal)
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
