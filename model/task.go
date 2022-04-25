package model

import (
	"time"
)

type TaskType string

type BaseTask struct {
	CreateTime time.Time         `json:"create_time"`
	UpdateTime time.Time         `json:"update_time"`
	DeleteTime time.Time         `json:"delete_time"`
	Id         string            `mapstructure:"id" json:"id"`
	Args       string            `mapstructure:"args" json:"args"`
	EnvMap     map[string]string `mapstructure:"env_map" json:"env_map,omitempty"`
	Name       string            `mapstructure:"name" json:"name"`
	Type       TaskType          `mapstructure:"type" json:"type" yaml:"type" `
	Status     int               `mapstructure:"status" json:"status"`
	TaskData   TaskData          `mapstructure:"task_data" json:"task_data,omitempty" `
}
type Task struct {
	TaskMetaData string `json:"task_meta_data"`
	BaseTask
}
type ShellTaskMetaData struct {
	Cmd string `mapstructure:"cmd" json:"cmd"`
}
type DockerTaskMetaData struct {
	Path      string `json:"path"`
	GitUrl    string `json:"git_url"`
	GitRef    string `json:"git_ref"`
	Tag       string `json:"tag"`
	ImageName string `json:"image_name"`
}
//todo
type PipelineTaskMetaData struct {

	Path      string `json:"path"`
	GitUrl    string `json:"git_url"`
	GitRef    string `json:"git_ref"`

}
type TaskData struct {
	StdErr string `json:"stdErr,omitempty"`
	Stdout string `json:"stdout,omitempty"`
	Error  string `json:"error,omitempty"`
}
