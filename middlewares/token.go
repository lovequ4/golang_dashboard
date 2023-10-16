package middlewares

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("secret")

func GenerateToken(userId int, name string, role string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"name":   name,
		"role":   role,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	tokenString, _ := token.SignedString(jwtSecret)
	return "Bearer " + tokenString, nil
}

//解析JWT
func ParseToken(tokenString string) (*jwt.MapClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
