package utils

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ix-pay/ixpay/config"
)

type JwtUtil struct {
	jwtKey []byte
}

func SetupJwt(cfg *config.Config) *JwtUtil {
	return &JwtUtil{
		jwtKey: []byte(cfg.JWTSecret),
	}
}

func (j *JwtUtil) GenerateJWT(userID int64) (string, error) {
	claims := jwt.MapClaims{
		"user_id": strconv.FormatInt(userID, 10),
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.jwtKey)
}

func (j *JwtUtil) ParseJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return j.jwtKey, nil
	})
}
