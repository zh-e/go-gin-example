package conf

import (
	"github.com/spf13/viper"
)

var Config allConfig

type allConfig struct {
	Host  host
	Mysql mysql
}

type host struct {
	Port string
}

type mysql struct {
	User     string
	Password string
	Host     string
	Port     string
	Database string

	MaxIdleConn int
	MaxOpenConn int
}

func InitConfig() {
	viper.SetConfigFile("./conf/app.yaml")
	viper.ReadInConfig()

	viper.Unmarshal(&Config)

}
