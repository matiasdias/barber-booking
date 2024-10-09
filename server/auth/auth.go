package auth

import (
	"api/server/config"
	"encoding/json"
	"io/ioutil"
	"log"

	"golang.org/x/oauth2"
)

var (
	GoogleOauthConfig *oauth2.Config
	err               error
)

func loadAuthConfig(filePath string) (config.AuthConfig, error) {
	var config config.AuthConfig

	// Lê o conteúdo do arquivo JSON
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	// Decodifica o conteúdo do arquivo JSON para a estrutura Config
	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}

func InitAuthOauth() (err error) {
	googleConfig, err := loadAuthConfig("config/config.api.json")
	if err != nil {
		log.Printf("Erro ao carregar a configuração do Google: %v", err)
		return err
	}

	GoogleOauthConfig = &oauth2.Config{
		ClientID:     googleConfig.ClientID,
		ClientSecret: googleConfig.ClientSecret,
		RedirectURL:  googleConfig.RedirectURL,
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
