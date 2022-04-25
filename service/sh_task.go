package service

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os/exec"
	"strings"
)

var (
	GitCloneTemplate = `
	pwd
	git clone {{git_url}}
	git checkout {{git_ref}}
	ls -alh
`
	DockerBuildTemplate = `
	pwd
	docker build -t {{tag}} -f {{dockerfile}} .
	docker push {{tag}}
`
)

func checkCmd(cmd ...string) bool {
	for _, c := range cmd {
		_, err := exec.LookPath(c)
		if err != nil {
			return false
		}
	}
	return true
}
func baseShellExec(cmd string, env map[string]string, args ...string) (string, error) {
	zap.L().Info("baseShellExec", zap.String("cmd:", cmd))
	if !checkCmd(cmd) {
		return "", errors.New(fmt.Sprintf("cmd[%s] not found ", cmd))
	}
	c := exec.Command(cmd, args...)
	//path, _ := os.Getwd()
	c.Dir = "/tmp"
	for k, v := range env {
		c.Env = append(c.Env, fmt.Sprintf("%s=%s", k, v))
	}
	output, err := c.CombinedOutput()
	if err != nil {
		zap.L().Error("baseShellExec", zap.Error(err))
		return string(output), err
	}
	zap.L().Info("baseShellExec", zap.String("cmd output", string(output)))
	return string(output), nil
}
func ExecShell(cmd string, env map[string]string, args ...string) (string, error) {
	return baseShellExec(cmd, env, args...)
}
func GitClone(gitUrl, gitRef string, env map[string]string) (string, error) {
	cmd := strings.ReplaceAll(GitCloneTemplate, "{{git_url}}", gitUrl)
	cmd = strings.ReplaceAll(GitCloneTemplate, "{{git_ref}}", gitRef)
	return baseShellExec("/bin/sh", env, "-xe", cmd)
}
func DockerBuild(tag, path string, env map[string]string) (string, error) {
	cmd := strings.ReplaceAll(DockerBuildTemplate, "{{tag}}", tag)
	cmd = strings.ReplaceAll(DockerBuildTemplate, "{{dockerfile}}", path)
	return baseShellExec("/bin/sh", env, "-xe", cmd)
}
func GoVersion() (string, error) {
	return baseShellExec("go", nil, "version")
}
