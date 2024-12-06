package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// 定义一个密钥，用于签名和验证Token，应该保存在安全的地方，这里只是示例
var jwtSecret = []byte("your_secret_key")

// 生成Token函数
func GenerateToken(username, password string) (string, error) {
	// 设置Token的过期时间，这里设置为1小时后过期
	expirationTime := time.Now().Add(time.Hour)
	// 创建一个新的Token对象，指定签名方法为HS256，并设置声明（Claims）
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"password": password,
		"exp":      expirationTime.Unix(),
	})
	// 使用密钥对Token进行签名，生成最终的Token字符串
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// 解析Token函数
func ParseToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法是否为HS256
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method)
		}
		return jwtSecret, nil
	})
}

// 从解析后的Token中获取用户名和密码等信息
func GetClaimsFromToken(token *jwt.Token) (map[string]interface{}, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}
	return nil, fmt.Errorf("invalid token claims")
}
