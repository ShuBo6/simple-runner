package request

type TaskRequest struct {
	Name string `json:"name"  binding:"required"`
	Type string `json:"type" binding:"required"`
	Cmd  string `json:"cmd"  binding:"required"`
	Env  string `json:"env"`
}
