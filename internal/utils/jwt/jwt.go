package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrTokenExpired     = errors.New("token 已过期, 请重新登录")
	ErrTokenNotValidYet = errors.New("token 无效, 请重新登录")
	ErrTokenMalformed   = errors.New("token 不正确, 请重新登录")
	ErrTokenInvalid     = errors.New("这不是一个 token, 请重新登录")
)

type MyClaims struct {
	UserId int `json:"user_id"`
	// UUID    string `json:"uuid"`
	jwt.RegisteredClaims
}

func CreateToken(secret, issuer string, expireHour, userId int) (string, error) {
	// 创建一个JWT对象
	//conf:=config.GetConfig().JWt
	if expireHour <= 0 {
		expireHour = 24
	}
	claim := MyClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expireHour))),
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(secret, tokenString string) (*MyClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效的签名算法")
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := jwtToken.Claims.(*MyClaims); ok && jwtToken.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}
