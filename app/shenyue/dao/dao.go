package dao

import (
	"context"
	"github.com/go-redis/redis"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"shenyue-gin/app/shenyue/model"
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
	//username := "jiang"            //账号
	//password := "D77jB5bTNSyKF7jb" //密码
	//host := "106.15.138.33"        //数据库地址，可以是Ip或者域名
	//port := 3306                   //数据库端口
	//Dbname := "jiang"              //数据库名
	//username := "root"   //账号
	//password := "123456" //密码
	//host := "127.0.0.1"  //数据库地址，可以是Ip或者域名
	//port := 3306         //数据库端口
	//Dbname := "blog"     //数据库名
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)

	dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local"
	//dsn := "jiang:D77jB5bTNSyKF7jb@tcp(106.15.138.33:3306)/jiang?charset=utf8&parseTime=True&loc=Local"
	newdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 自动迁移表结构
	newdb.AutoMigrate(&model.User{})
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
