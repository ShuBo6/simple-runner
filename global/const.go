package global

const (
	Ready    = 0
	Running  = 1
	Finished = 2
	Failed   = 3

	ShellTask    = "shell"
	DockerTask   = "docker"
	PipelineTask = "pipeline"
)
