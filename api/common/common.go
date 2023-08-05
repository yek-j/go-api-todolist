package common

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
)

type JwtClaims struct {
	UserNM string `json:"usernm"`
	UserID string `json:"usernmid"`
	jwt.StandardClaims
}

func ValidateToken(strToken string) (*JwtClaims, error) {
	var claims JwtClaims

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("config")
	readErr := viper.ReadInConfig()
	
	fmt.Print("여기?")
	if readErr != nil {
		return nil, readErr
	}

	secretkey := []byte(viper.GetString("secretkey"))

	token, err := jwt.ParseWithClaims(strToken, &claims, func(t *jwt.Token) (interface{}, error) {
		return secretkey, nil
	})

	fmt.Print("여기")
	if err != nil {
		return nil, err 
	}
	fmt.Print("여기2")
	if token.Valid {
		return &claims, nil 
	} 
	return nil, fmt.Errorf("invalid token")
}