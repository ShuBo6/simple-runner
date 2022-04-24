package model

import (
	"time"
)

type TaskType string

const (
	ShellTask    = "shell"
	DockerTask   = "docker"
	PipelineTask = "pipeline"
)

type Task struct {
	CreateTime time.Time         `json:"create_time"`
	UpdateTime time.Time         `json:"update_time"`
	DeleteTime time.Time         `json:"delete_time"`
	Id         string            `mapstructure:"id" json:"id"`
	Cmd        string            `mapstructure:"cmd" json:"cmd"`
	Args       string            `mapstructure:"args" json:"args"`
	EnvMap     map[string]string `mapstructure:"env_map" json:"env_map,omitempty"`
	Name       string            `mapstructure:"name" json:"name"`
	Type       TaskType          `mapstructure:"type" json:"type" yaml:"type" `
	Status     int               `mapstructure:"status" json:"status"`
	TaskData   TaskData          `mapstructure:"task_data" json:"task_data,omitempty" `
}
type TaskData struct {
	StdErr string `json:"stdErr,omitempty"`
	Stdout string `json:"stdout,omitempty"`
	Error  string `json:"error,omitempty"`
}
