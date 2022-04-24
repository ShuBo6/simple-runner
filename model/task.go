package model



type TaskType string

const (
	ShellTask    = "shell"
	DockerTask   = "docker"
	PipelineTask = "pipeline"
)

type Task struct {
	Id       string            `mapstructure:"id" json:"id"`
	Cmd      string            `mapstructure:"cmd" json:"cmd"`
	EnvMap   map[string]string `mapstructure:"env_map" json:"env_map,omitempty"`
	Name     string            `mapstructure:"name" json:"name"`
	Type     TaskType          `mapstructure:"type" json:"type" yaml:"type" `
	Status   int               `mapstructure:"status" json:"status"`
	TaskData *TaskData         `mapstructure:"task_data" json:"task_data,omitempty" `
}
type TaskData struct {
	StdErr string `json:"stdErr"`
	Stdout string `json:"stdout"`
}

