package utils

import (
	"e-food/constants"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/martian/log"
	"strings"
	"time"
)

func ValidateHeader(bearerHeader string) (interface{}, error) {
	bearerToken := strings.Split(bearerHeader, " ")[1]
	claims := jwt.MapClaims{}
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
	if token.Valid {
		return claims["user"].(string), nil
	}
	return nil, errors.New("invalid token")
}

func GenerateJWT(userEmail, fname, lname string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = userEmail
	claims["fname"] = fname
	claims["lname"] = lname
	claims["exp"] = time.Now().Add(time.Minute * 90).Unix()

	tokenString, err := token.SignedString(constants.MySecretKeyForJWT)
	if err != nil {
		log.Errorf("Error generating Token: " + err.Error())
		return "", err
	}
	return tokenString, nil
}
