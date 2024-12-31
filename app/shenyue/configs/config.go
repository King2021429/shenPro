package configs

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

// DbConfig 结构体用于存储 Db 部分的配置
type DbConfig struct {
	Addr         string
	Dsn          string
	Active       int
	Idle         int
	IdleTimeout  string
	QueryTimeout string
	ExecTimeout  string
	TranTimeout  string
}

// DbLocalConfig 结构体用于存储 DbLocal 部分的配置
type DbLocalConfig struct {
	Addr string
	Dsn  string
}

// Config 结构体用于存储整个 TOML 文件的配置
type Config struct {
	Db      DbConfig
	DbLocal DbLocalConfig
}

var conf *Config

func GetConfig() *Config {
	return conf
}

func InitConfig() {
	if _, err := toml.DecodeFile("app/shenyue/configs/application.toml", &conf); err != nil {
		fmt.Println("读取TOML文件出错:", err)
		return
	}
}
