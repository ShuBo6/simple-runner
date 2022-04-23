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
	Id     string            `json:"id"`
	Cmd    string            `json:"cmd"`
	EnvMap map[string]string `json:"env_map"`

	Name string   `json:"name"`
	Type TaskType `yaml:"type" json:"type"`
	//Status int      `json:"status"`
}
type TaskData struct {

	//Stdout string            `json:"stdout"`
}
