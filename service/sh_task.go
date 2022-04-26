package service

import (
	"errors"
	"fmt"
	"go.uber.org/zap"
	"os/exec"
	"regexp"
	"strings"
)

var (
	GitCloneTemplate = `
	pwd 
	if [[ -d {{dir}} ]]
	then
    	rm -rf {{dir}}
	fi
	git clone {{git_url}} 
	cd {{dir}} 
	git checkout {{git_ref}} 
	ls -alh 
	`

	DockerBuildTemplate = `
	pwd
	cd {{dir}}
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
	//fmt.Println(c)
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
	strs := SplitUrl(gitUrl)
	name := strs[3]
	cmd := strings.ReplaceAll(GitCloneTemplate, "{{git_url}}", gitUrl)
	cmd = strings.ReplaceAll(cmd, "{{git_ref}}", gitRef)
	cmd = strings.ReplaceAll(cmd, "{{dir}}", name)
	return baseShellExec("/bin/bash", env, "-cxe", cmd)
}
func DockerBuild(tag, path, filepath string, env map[string]string) (string, error) {
	cmd := strings.ReplaceAll(DockerBuildTemplate, "{{tag}}", tag)
	cmd = strings.ReplaceAll(cmd, "{{dockerfile}}", filepath)
	cmd = strings.ReplaceAll(cmd, "{{dir}}", path)
	return baseShellExec("/bin/bash", env, "-cxe", cmd)
}
//func Pipeline(start, end *model.Stag, cmd string, env map[string]string) (string, error) {
//
//	return
//}

func GoVersion() (string, error) {
	return baseShellExec("go", nil, "version")
}
func CheckRepoUrl(repoUrl string) bool {
	reg := regexp.MustCompile(`^(?:git@(coding|git).jd.com:|http[s]?://(coding|git).jd.com/)([\w._-]+)/([\w._-]+)(?:.git|/)?$`)
	return reg.MatchString(repoUrl)
}
func SplitUrl(url string) []string {
	reg := regexp.MustCompile(`^(ssh://git@home.shubo6.cn:30001/)([\w._-]+)/([\w_-]+)(?:.git|/)?$`)
	return reg.FindStringSubmatch(url)
}

//func GetGitClient(queryUrl string) git_client.GitClientInterface {
//	strs := SplitUrl(queryUrl)
//	repoType := strs[2]
//	if repoType == "" {
//		repoType = strs[1]
//	}
//
//	return gitClient
//}
