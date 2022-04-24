package initial

import (
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"simple-cicd/global"
)

func InitViper() {
	_v := viper.New()
	_v.SetConfigFile("conf/config.yaml")
	if err := _v.ReadInConfig(); err != nil {
		panic(fmt.Sprintf(`读取config.yaml文件失败, err: %v`, err))
	}
	err := _v.Unmarshal(&global.C)
	zap.L().Info("global",zap.Any("",global.C))
	if err != nil {
		panic(fmt.Sprintf(`读取config.yaml文件失败, err: %v`, err))
	}
}
