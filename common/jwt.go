package common

import (
	"account-go/model"
	"github.com/dgrijalva/jwt-go"
	"time"
)

// jwt的秘钥
var jwtKey = []byte("a_secret_crect")

type Claims struct {
	UserId int
	jwt.StandardClaims
}

// ReleaseToken 获取token
func ReleaseToken(user model.User) (string, error) {
	// 设定过期时间
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: user.Id,
		StandardClaims: jwt.StandardClaims{
			// 过期时间
			ExpiresAt: expirationTime.Unix(),
			// 发放时间
			IssuedAt: time.Now().Unix(),
			// 发放人
			Issuer: "oceanlearn.tech",
			// 主题
			Subject: "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return token, claims, err
}
