package dao

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Dao
type Dao struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewDao() (dao *Dao) {
	dao = &Dao{
		db:  NewGorm(),
		rdb: NewRedis(),
	}
	return dao
}

func NewGorm() (newdb *gorm.DB) {
	//配置MySQL连接参数
	username := "jiang"            //账号
	password := "D77jB5bTNSyKF7jb" //密码
	host := "106.15.138.33"        //数据库地址，可以是Ip或者域名
	port := 3306                   //数据库端口
	Dbname := "jiang"              //数据库名
	//username := "root"   //账号
	//password := "123456" //密码
	//host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	//port := 3306         //数据库端口
	//Dbname := "blog"     //数据库名
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	newdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	return newdb
}

func NewRedis() (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

// Close close the resource.
func (d *Dao) Close() {

}

// Ping ping the resource.
func (d *Dao) Ping(ctx context.Context) (err error) {
	return nil
}
