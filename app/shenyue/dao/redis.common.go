package dao

import (
	"context"
	"fmt"
	"time"
)

func (d *Dao) RcSet(ctx context.Context, key string, value string, expiration time.Duration) (err error) {
	// 设置一个键值对到Redis
	err = d.rdb.Set(key, value, expiration).Err()
	if err != nil {
		fmt.Println("设置键值对失败:", err)
		return
	}
	fmt.Println("成功设置键值对到Redis")
	return
}

func (d *Dao) RcGet(ctx context.Context, key string) (value string, err error) {
	// 从Redis获取刚才设置的值
	value, err = d.rdb.Get(key).Result()
	if err != nil {
		fmt.Println("获取值失败:", err)
		return
	}
	return
}

func (d *Dao) RcDel(ctx context.Context, key string) (err error) {
	err = d.rdb.Del(key).Err()
	if err != nil {
		fmt.Println("删除键值对失败:", err)
		return
	}
	return
}
