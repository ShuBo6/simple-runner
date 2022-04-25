配置文件conf/config.yaml 模板如下
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