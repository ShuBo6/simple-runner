package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"simple-cicd/global"
	"simple-cicd/model"
	"simple-cicd/model/request"
	"simple-cicd/model/response"
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
		return &response.Response{Code: response.ERROR, Error: err}
	}
	envMap := make(map[string]string)
	if req.Env != "" {
		err = json.Unmarshal([]byte(req.Env), envMap)
		if err != nil {
			return &response.Response{Code: response.ERROR, Error: err}
		}
	}
	task := &model.Task{
		Id:     uuid.New().String(),
		Name:   req.Name,
		Type:   model.TaskType(req.Type),
		Cmd:    req.Cmd,
		EnvMap: envMap,
	}
	if err != nil {
		return &response.Response{Code: response.ERROR, Error: err}
	}
	err = json.Unmarshal(v, task)
	if err != nil {
		return &response.Response{Code: response.ERROR, Error: err}
	}
	global.ChannelTaskQueue <- task
	return &response.Response{Code: response.SUCCESS, Data: v, Message: msg}
}
func ListTask(ctx *gin.Context) *response.Response {
	ret, err := json.Marshal(global.ChannelTaskQueue)
	if err != nil {
		return &response.Response{Code: response.ERROR, Error: err}
	}
	return &response.Response{Code: response.SUCCESS, Data: string(ret)}
}
