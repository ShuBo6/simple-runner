package model

const (
	Ready    = 0
	Running  = 1
	Finished = 2
	Failed   = 3
)

type Task struct {
	Id     string
	Data   string
	Name   string
	Status int
}
type TaskData struct {
	Cmd    string
	EnvMap map[string]string
	Stdout string
}
