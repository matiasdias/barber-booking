package token

import (
	"api/server/config"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	Jwt config.JwtConfig
	err error
)

// Claims define os dados que vão no token
type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func InitJwt() error {
	Jwt, err = config.LoadJwtConfig("config/config.api.json")
	if err != nil {
		log.Printf("Erro ao carregar a configuração do JWT: %v", err)
		return err
	}
	return nil
}

func GenerateJWT(email *string) (string, error) {
	expirationTime := time.Now().Add(time.Duration(Jwt.Jwt.ExpireTime) * time.Second).Unix()
	claims := &Claims{
		Email: *email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(Jwt.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil

}

func GenerateRefreshToken(email *string) (string, error) {
	expirationTime := time.Now().Add(30 * 24 * time.Hour).Unix()
	claims := &Claims{
		Email: *email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(Jwt.Jwt.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	secret := []byte(Jwt.Jwt.Secret)
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return claims, nil
}

func ValidateRefreshToken(refreshToken string) (*Claims, error) {
	secret := []byte(Jwt.Jwt.Secret)
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || claims.ExpiresAt < time.Now().Unix() {
		return nil, fmt.Errorf("token expired")
	}

	return claims, nil
}
