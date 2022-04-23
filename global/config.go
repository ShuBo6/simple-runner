package global

var C = &config{}

type config struct {
	Zap        Zap         `mapstructure:"zap" json:"zap" yaml:"zap"`
	EtcdConfig *EtcdConfig `mapstructure:"etcd" yaml:"etcd"`
	System     *System     `mapstructure:"system" yaml:"system"`
}
type System struct {
	Port string `yaml:"port"`
}

type EtcdConfig struct {
	Endpoints []string `mapstructure:"endpoints" yaml:"endpoints"`
	Username  string   `mapstructure:"username" yaml:"username"`
	Password  string   `mapstructure:"password" yaml:"password"`
	RootPath  string   `mapstructure:"root_path" yaml:"root_path"`
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

