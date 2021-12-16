package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
	"time"
)

var jwtSecret = []byte(viper.GetString("JWT_SECRET"))

type Claims struct {
	Uid      uint   `json:"uid"`
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func GenerateToken(uid uint, username, password string) (string, error) {
	claims := Claims{
		uid,
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "fiber-user",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
