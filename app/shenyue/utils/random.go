package utils

import (
	"github.com/bwmarrin/snowflake"
	"math/rand"
	"time"
)

// GenerateUserIdSnow 生成唯一用户 ID
func GenerateUserIdSnow() (int64, error) {
	// 创建一个节点，节点 ID 范围是 0 - 1023
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}
	return node.Generate().Int64(), nil
}

// GenerateUserId 生成唯一用户 ID
func GenerateUserId() int64 {
	// 获取当前时间戳（秒级）
	timestamp := time.Now().Unix()
	// 生成一个随机数
	rand.Seed(time.Now().UnixNano())
	randomNum := rand.Int63n(1000)
	// 组合时间戳和随机数
	id := timestamp*1000 + randomNum
	return id
}
