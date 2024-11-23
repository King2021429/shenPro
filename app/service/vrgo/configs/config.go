package configs

import (
	"go-common/library/cache/redis"
	"go-common/library/database/sql"
	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/rpc/warden"
	"time"
)

var (
	conf *Config
)

type Config struct {
	DB    *DB
	Redis *Redis

	Host       *Host
	BM         *bm.ServerConfig
	HTTPClient *HTTPClient
	Warden     *warden.ServerConfig
	GRPCClient *warden.ClientConfig
}

// Host host config .
type Host struct {
	Api     string
	Manager string
	OaHr    string
}

type DB struct {
	Mysql *sql.Config
}

// HTTPClient conf.
type HTTPClient struct {
	Normal *bm.ClientConfig
	Slow   *bm.ClientConfig
}

// Redis conf
type Redis struct {
	Default *struct {
		*redis.Config
		Expire time.Duration
	}
}

func GetConfig() *Config {
	return conf
}

func AfterLoad() {

}

func Init() {
	defer AfterLoad()
	loadConfig()
}

func loadConfig() {
	//conf, pl = new(Config), new(paladin.Map)
	//
	//files := []string{"http.toml", "grpc.toml", "db.toml", "redis.toml", "application.toml"}
	//for _, v := range files {
	//	checkErr(paladin.Watch(v, pl))
	//	checkErr(paladin.Get(v).UnmarshalTOML(&conf))
	//}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
