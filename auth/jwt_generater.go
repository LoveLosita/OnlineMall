package auth

import (
	"OnlineMall/respond"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

var JwtSecret = []byte("OnlineMallJWT") // 用于签名和验证 Token 的密钥

// GenerateTokens 生成访问令牌和刷新令牌
func GenerateTokens(userID int) (string, string, error) {
	// 创建访问令牌
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,                                  // 获取用户ID
		"exp":     time.Now().Add(15 * time.Minute).Unix(), // 设置访问令牌过期时间为 15 分钟
	})

	// 使用密钥签名访问令牌
	accessTokenString, err := accessToken.SignedString(JwtSecret)
	if err != nil {
		return "", "", err
	}

	// 创建刷新令牌
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,                                    // 获取用户ID
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(), // 设置刷新令牌过期时间为 7 天
	})

	// 使用密钥签名刷新令牌
	refreshTokenString, err := refreshToken.SignedString(JwtSecret)
	if err != nil {
		return "", "", err
	}

	return accessTokenString, refreshTokenString, nil
}

// ValidateToken 验证令牌
func ValidateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 确保签名方法是我们支持的 HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, respond.InvalidTokenSingingMethod
		}
		return JwtSecret, nil
	})
	return token, err
}
