package model

// Kratos hello kratos.
type Kratos struct {
	Hello string
}

type Article struct {
	ID      int64
	Content string
	Author  string
}

// Position 定义了一个位置的结构体，包含两个整数坐标
type Position [2]int

// CharacterMap 定义了一个角色映射的结构体，包含角色名称和其位置
type CharacterMap map[string]Position

// ImageMap 定义了一个图像映射的结构体，包含图像名称和角色映射
type ImageMap map[string]CharacterMap
