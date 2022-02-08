package model

const (
	Ready    = 0
	Running  = 1
	Finished = 2
	Failed   = 3
)

type Task struct {
	Id     string   `json:"id"`
	Data   TaskData `json:"data"`
	Name   string   `json:"name"`
	Status int      `json:"status"`
}
type TaskData struct {
	Cmd    string            `json:"cmd"`
	EnvMap map[string]string `json:"env_map"`
	Stdout string            `json:"stdout"`
}
