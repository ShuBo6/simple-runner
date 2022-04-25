# 配置文件conf/config.yaml 模板如下

```yaml
# zap logger configuration
zap:
  level: 'info'
  format: 'console'
  prefix: '[runner]'
  director: 'logs'
  link-name: 'latest_log'
  show-line: true
  encode-level: 'LowercaseColorLevelEncoder'
  stacktrace-key: 'disableStacktrace'
  log-in-console: true

system:
  port: "5080"
etcd:
  endpoints:
    - "http://192.168.123.4:2379"
  #    - "http://127.0.0.1:2379"
  username: ""
  password: ""
  root_path: "/runner"
  task_path: "/runner/task_path"
  history_task_path: "/runner/history_task_path"

email:
  from: "xx@qq.com"
  name: "xxxxx"
  user: "xxxxx"
  host: "smtp.qq.com"
  secret: "xxxxxxxxx"
  port: 587
```

# 创建shell任务

```shell
curl --location --request POST 'localhost:5080/task' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "test",
    "type": "shell",
    "shell_build": {
        "cmd": "go"
    },
    "args": "version"
}'
```

# 创建docker build任务

```shell
curl --location --request POST 'localhost:5080/task' \
--header 'Content-Type: application/json' \
--data-raw '{
    "name": "test",
    "type": "docker",
    "docker_build": {
        "git_url": "ssh://git@home.shubo6.cn:30001/shubo6/docker-hello-world.git",
        "git_ref": "master",
        "path": "Dockerfile",
        "image_name": "runner",
        "tag": "1.0.2"
    },
    "args": "version"
}'
```

# 查询shell任务列表

```shell
curl --location --request GET 'localhost:5080/task'
```