package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// JWT 密钥
const jwtSecret = "ced2850a8efb4b52aa5db779152fca5d"

// GenerateToken 生成 JWT 令牌
func GenerateToken(userID uint) (string, error) {
	claims := &jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userID)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// AuthMiddleware 用于验证 JWT
// Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzUzODk4MDUsImlzcyI6IjAifQ.4YsG5MX6eGgoLXCwTldfH-JfdSV-mMSv152A3RYsVlQ
func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "未提供认证信息"})
		c.Abort()
		return
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的认证格式"})
		c.Abort()
		return
	}

	tokenStr := parts[1]
	claims := &jwt.StandardClaims{}
	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil || !tkn.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "无效的令牌"})
		c.Abort()
		return
	}

	// 将用户 ID 等信息存储到上下文，方便后续使用
	c.Set("userID", claims.Issuer)
	// 生成唯一的请求ID
	requestID := uuid.New().String()
	c.Set("request_id", requestID)
	c.Next()
}

// AdminAuthMiddleware 用于验证管理员权限
func AdminAuthMiddleware(c *gin.Context) {
	// 先调用普通的身份验证中间件
	AuthMiddleware(c)
	if c.IsAborted() {
		return
	}
	// 从上下文中获取用户信息，假设用户信息存储在 claims 中
	userID := c.GetString("userID")
	// 这里可以根据 userID 从数据库中查询用户角色
	fmt.Println(userID)
	// 简单示例中直接从上下文中获取角色信息
	role := c.GetString("role")
	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{"error": "没有管理员权限"})
		c.Abort()
		return
	}
	c.Next()

}
