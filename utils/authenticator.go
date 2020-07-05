package utils

import (
	"e-food/constants"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/martian/log"
	"strings"
)

//type

func ValidateHeader(bearerHeader string) (interface{}, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	token, err := jwt.Parse(bearerToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error decoding token")
		}
		return constants.MySecretKeyForJWT, nil
	})
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	return token.Valid, nil
}
