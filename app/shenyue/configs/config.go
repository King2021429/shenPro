package configs

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

// DbConfig 结构体用于存储 Db 部分的配置
type DbConfig struct {
	Dsn string
	//Addr         string
	//Active       int
	//Idle         int
	//IdleTimeout  string
	//QueryTimeout string
	//ExecTimeout  string
	//TranTimeout  string
}

// EmailConfig 结构体用于email部分的配置
type EmailConfig struct {
	EmailUser string
	EmailPwd  string
	EmailHost string
	EmailPort int
}

type MoonshotConfig struct {
	Url           string
	Authorization string
}

// Config 结构体用于存储整个 TOML 文件的配置
type Config struct {
	Db       DbConfig
	Email    EmailConfig
	Moonshot MoonshotConfig
}

var conf *Config

func GetConfig() *Config {
	return conf
}

func InitConfig() {
	var url string
	// 根据环境读取配置文件
	//v := os.Getenv("env")
	url = "app/shenyue/configs/local.toml"
	//url = "/www/wwwroot/goproject/shenyue-gin/app/shenyue/configs/prod.toml"

	if _, err := toml.DecodeFile(url, &conf); err != nil {
		fmt.Println("读取TOML文件出错:", err)
		return
	}
	fmt.Println(conf)
	fmt.Println(conf.Db)
	fmt.Println(conf.Email)
}
