package service

import (
	"fmt"
	"testing"
)

func TestDockerBuild(t *testing.T) {
	fmt.Println(GitClone("ssh://git@home.shubo6.cn:30001/shubo6/docker-hello-world.git", "master", nil))
	fmt.Println(DockerBuild("t:1.0.1", "docker-hello-world","Dockerfile",  nil))

}
