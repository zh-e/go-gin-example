package conf

import (
	"github.com/spf13/viper"
	"log"
)

var Config allConfig

type allConfig struct {
	Host  host
	Mysql mysql
	Redis redis
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

type redis struct {
	Host     string
	Port     string
	Password string
	DB       int
}

func init() {
	viper.SetConfigFile("./conf/app.yaml")
	viper.ReadInConfig()

	viper.Unmarshal(&Config)

	log.Println("配置加载完成")
}
