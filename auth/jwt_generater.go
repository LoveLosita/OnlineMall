package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var JwtSecret = []byte("OnlineMallJWT") // 用于签名和验证 Token 的密钥

func GenerateJWT(userID int) (string, error) {
	// 创建 JWT
	//fmt.Println(userID)//测试用
	//fmt.Println(string(jwtSecret))//测试用
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,                                // 获取用户id
		"exp":     time.Now().Add(24 * time.Hour).Unix(), // 设置 Token 过期时间为 24 小时
	})

	// 使用密钥签名 Token
	tokenString, err := token.SignedString(JwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
