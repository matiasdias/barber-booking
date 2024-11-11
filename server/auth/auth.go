package auth

import (
	"api/server/config"
	"fmt"
	"log"

	"golang.org/x/oauth2"
)

var (
	GoogleOauthConfig *oauth2.Config
	err               error
)

func InitAuthOauth() (err error) {
	config, err := config.LoadAuthConfig("config/config.api.json")
	if err != nil {
		log.Printf("Erro ao carregar a configuração do Google: %v", err)
		return err
	}
	if config.Oauth.ClientID == "" || config.Oauth.ClientSecret == "" {
		log.Println("Erro: GOOGLE_CLIENT_ID ou GOOGLE_CLIENT_SECRET não estão definidos.")
		return fmt.Errorf("credenciais do Google não estão definidas")
	}
	GoogleOauthConfig = &oauth2.Config{
		ClientID:     config.Oauth.ClientID,
		ClientSecret: config.Oauth.ClientSecret,
		RedirectURL:  config.Oauth.RedirectURL,
		Scopes: []string{
			"openid",
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/auth",
			TokenURL: "https://accounts.google.com/o/oauth2/token",
		},
	}
	return nil
}
