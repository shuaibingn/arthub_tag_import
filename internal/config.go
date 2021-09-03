package internal

import (
	"flag"
	"log"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type ConfigData struct {
	Domain string
	Token  string
	Depot  string
	File   struct {
		Path string
		Name string
	}
}

var GlobalConfig = func() *ConfigData {
	// 读取配置文件
	c := new(ConfigData)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".\\")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("read config error, %v\n", err)
	}

	// 读取命令行参数
	pflag.String("domain", "", "arthub api")
	pflag.String("token", "", "arthub token")
	pflag.String("depot", "", "arthub depot")
	pflag.String("file.name", "", "arthub tag file name")
	pflag.String("file.path", "", "arthub tag file path")

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if err := viper.BindPFlags(pflag.CommandLine); err != nil {
		log.Fatalf("bind from cmd error, %v\n", err)
	}

	// 序列化到结构体
	if err := viper.Unmarshal(c); err != nil {
		log.Fatalf("unmarshal config error, %v\n", err)
	}
	return c
}()
