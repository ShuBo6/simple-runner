package global

var C = &config{}

type config struct {
	Zap        Zap         `mapstructure:"zap" json:"zap" yaml:"zap"`
	EtcdConfig *EtcdConfig `mapstructure:"etcd" yaml:"etcd"`
	System     *System     `mapstructure:"system" yaml:"system"`
	Email      *Email      `mapstructure:"email" yaml:"email"`
}
type System struct {
	Port string `yaml:"port"`
}

type EtcdConfig struct {
	Endpoints       []string `mapstructure:"endpoints" yaml:"endpoints"`
	Username        string   `mapstructure:"username" yaml:"username"`
	Password        string   `mapstructure:"password" yaml:"password"`
	RootPath        string   `mapstructure:"root_path" yaml:"root_path"`
	TaskPath        string   `mapstructure:"task_path" yaml:"task_path"`
	HistoryTaskPath string   `mapstructure:"history_task_path" yaml:"history_task_path"`
}
type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"`
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`
}
type Email struct {
	From   string `mapstructure:"from" json:"from" yaml:"from"` // 发件人  你自己要发邮件的邮箱
	Name   string `mapstructure:"name" json:"name" yaml:"name"`
	Host   string `mapstructure:"host" json:"host" yaml:"host"` // 服务器地址 例如 smtp.qq.com  请前往QQ或者你要发邮件的邮箱查看其smtp协议
	User   string `mapstructure:"user" json:"user" yaml:"user"`
	Secret string `mapstructure:"secret" json:"secret" yaml:"secret"` // 密钥    用于登录的密钥 最好不要用邮箱密码 去邮箱smtp申请一个用于登录的密钥
	Port   int    `mapstructure:"port" json:"port" yaml:"port"`       // 端口     请前往QQ或者你要发邮件的邮箱查看其smtp协议 大多为 465
}
