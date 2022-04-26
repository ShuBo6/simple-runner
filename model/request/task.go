package request

import "simple-cicd/model"

type TaskRequest struct {
	Name        string                      `json:"name"  binding:""`
	Type        string                      `json:"type" binding:"required"`
	DockerBuild *model.DockerTaskMetaData   `json:"docker_build"`
	ShellBuild  *model.ShellTaskMetaData    `json:"shell_build"`
	Pipeline    *model.PipelineTaskMetaData `json:"pipeline"`
	Args        string                      `json:"args"` //逗号分隔
	Env         string                      `json:"env"`
}

type ListTaskRequest struct {
	Scope int `json:"scope"` // 0 查所有，1查未开始，2查历史
}
