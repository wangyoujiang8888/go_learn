package viper_test

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

func NewConfig() *viper.Viper {
	conf := viper.New()
	env := os.Getenv("GO_ENV")
	fmt.Println(env)
	if env == "" {
		env = "local"
		//panic("GO_ENV is not found!")
	}
	conf.SetConfigName("config_" + env) // 设置配置文件名 (不带后缀)

	//conf.AddConfigPath("./conf")     // 第一个搜索路径
	//conf.AddConfigPath("../conf") // 第一个搜索路径
	//conf.AddConfigPath("../../conf") // 第一个搜索路径
	conf.AddConfigPath("/Users/river/www/go_test/example/viper_test") // 第一个搜索路径
	//conf.AddConfigPath("./conf")     // 第一个搜索路径
	//conf.AddConfigPath("./")         // 第一个搜索路径
	conf.SetConfigType("yml")  // 文件类型
	err := conf.ReadInConfig() // 读取配置数据
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return conf
}
