package config

import (
	"github.com/prometheus/common/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

var C = &config{}

type config struct {
	EtcdConfig *EtcdConfig `yaml:"etcd"`
}
type EtcdConfig struct {
	Endpoints []string `yaml:"endpoints"`
	Username  string   `yaml:"username"`
	Password  string   `yaml:"password"`
	RootPath  string   `yaml:"root_path"`
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
