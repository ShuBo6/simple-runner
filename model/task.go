package model

const (
	Ready    = 0
	Running  = 1
	Finished = 2
	Failed   = 3
)

type TaskType string

const (
	ShellTask    = "shell"
	DockerTask   = "docker"
	PipelineTask = "pipeline"
)

type Task struct {
	Id     string            `mapstructure:"id" json:"id"`
	Cmd    string            `mapstructure:"cmd" json:"cmd"`
	EnvMap map[string]string `mapstructure:"env_map" json:"env_map"`
	Name   string            `mapstructure:"name" json:"name"`
	Type   TaskType          `mapstructure:"type" json:"type" yaml:"type" `
	//Status int      `json:"status"`
}
type TaskData struct {

	//Stdout string            `json:"stdout"`
}
