package utils

import (
	"github.com/bwmarrin/snowflake"
)

// GenerateUserID 生成唯一用户 ID
func GenerateUserID() (int64, error) {
	// 创建一个节点，节点 ID 范围是 0 - 1023
	node, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}
	return node.Generate().Int64(), nil
}
