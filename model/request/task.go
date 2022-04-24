package request

type TaskRequest struct {
	Name string `json:"name"  binding:"required"`
	Type string `json:"type" binding:"required"`
	Cmd  string `json:"cmd"  binding:"required"`
	Args string `json:"args"`
	Env  string `json:"env"`
}

type ListTaskRequest struct {
	Scope int `json:"scope"` // 0 查所有，1查未开始，2查历史
}
