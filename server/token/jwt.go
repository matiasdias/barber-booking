package token

import (
	"api/server/config"
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func loadJwtConfig(filePath string) (config.JwtConfig, error) {
	var jwtRead config.JwtConfig

	// Lê o conteúdo do arquivo JSON
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return jwtRead, err
	}

	// Decodifica o conteúdo do arquivo JSON para a estrutura APIConfig
	if err := json.Unmarshal(data, &jwtRead); err != nil {
		return jwtRead, err
	}

	return jwtRead, nil
}

func InitJwt() error {
	Jwt, err = loadJwtConfig("config/config.api.json")
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

func ValidateToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(Jwt.Jwt.Secret), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return claims, nil
}
