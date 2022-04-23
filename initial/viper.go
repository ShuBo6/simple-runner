package initial

import (
	"github.com/prometheus/common/log"
	"simple-cicd/config"
)

func InitViper() {
	err := config.Load("conf/config.yaml")
	if err != nil {
		log.Error("load config path(conf/config.yaml) failed.")
		return
	}
}