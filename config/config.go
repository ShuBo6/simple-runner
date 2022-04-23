package config

import (
	"fmt"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var C = &config{}

type config struct {
	EtcdConfig *EtcdConfig `yaml:"etcd"`
	System     *System     `yaml:"system"`
}
type System struct {
	Port string `yaml:"port"`
}
type EtcdConfig struct {
	Endpoints []string `yaml:"endpoints"`
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
	RootPath  string   `yaml:"root_path"`
}

func InitViper() {
	_v := viper.New()
	_v.SetConfigFile("conf/config.yaml")
	if err := _v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf(`读取config.yaml文件失败, err: %v`, err))
	}
}
func Load(configPath string) error {
	abs, _ := filepath.Abs(configPath)
	log.Infof("load config from config file %s", abs)
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(data, C); err != nil {
		return err
	}
	return nil
}
