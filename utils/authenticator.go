package utils

import (
	"e-food/constants"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/martian/log"
	"strings"
)

type TokenInfo struct {
	IsValid bool
	Email   string
}

func ValidateHeader(bearerHeader string) (interface{}, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	claims := jwt.MapClaims{}
	tokenInfo := &TokenInfo{}
	token, err := jwt.ParseWithClaims(bearerToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding token")
		}
		return constants.MySecretKeyForJWT, nil
	})
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	tokenInfo.IsValid = token.Valid
	tokenInfo.Email = claims["user"].(string)
	return tokenInfo, nil
}
