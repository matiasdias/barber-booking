package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

func replaceEnvVariables(cfg *APIConfig) {
	cfg.APIPort = strings.ReplaceAll(cfg.APIPort, "${API_PORT}", os.Getenv("API_PORT"))
}
func replaceEnvOauthVariables(cfg *GoogleConfig) {
	cfg.Oauth.ClientID = strings.ReplaceAll(cfg.Oauth.ClientID, "${GOOGLE_CLIENT_ID}", os.Getenv("GOOGLE_CLIENT_ID"))
	cfg.Oauth.ClientSecret = strings.ReplaceAll(cfg.Oauth.ClientSecret, "${GOOGLE_CLIENT_SECRET}", os.Getenv("GOOGLE_CLIENT_SECRET"))
}

func LoadPortConfig(filePath string) (APIConfig, error) {
	var config APIConfig

	// Lê o conteúdo do arquivo JSON
	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	// Decodifica o conteúdo do arquivo JSON para a estrutura APIConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}

	replaceEnvVariables(&config)

	return config, nil
}

func LoadEnvironmentConfig(filePath string) (Config, error) {
	var config Config

	// Lê o conteúdo do arquivo JSON
	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	// Decodifica o conteúdo do arquivo JSON para a estrutura Config
	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}

func LoadAuthConfig(filePath string) (GoogleConfig, error) {
	var config GoogleConfig

	// Lê o conteúdo do arquivo JSON
	data, err := os.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	// Decodifica o conteúdo do arquivo JSON para a estrutura Config
	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}

	replaceEnvOauthVariables(&config)

	return config, nil
}

func LoadJwtConfig(filePath string) (JwtConfig, error) {
	var jwtRead JwtConfig

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
