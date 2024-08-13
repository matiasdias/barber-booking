package database

import (
	"api/server/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	_ "github.com/lib/pq"
)

var (
	Driver = "postgres"

	APIConfigInfo config.APIConfig
	Config        config.Config
	err           error
	dbconfig      config.DatabaseConfig
)

func loadAPIConfig(filePath string) (config.APIConfig, error) {
	var config config.APIConfig

	// Lê o conteúdo do arquivo JSON
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, err
	}

	// Decodifica o conteúdo do arquivo JSON para a estrutura APIConfig
	if err := json.Unmarshal(data, &config); err != nil {
		return config, err
	}

	return config, nil
}

func loadEnvironmentConfig(filePath string) (config.Config, error) {
	var config config.Config

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

// Connection conecta com o banco de dados
func Connection() (*sql.DB, error) {

	APIConfigInfo, err = loadAPIConfig("config/config.api.json")
	if err != nil {
		return nil, fmt.Errorf("failed to load API config: %w", err)
	}
	Config, err = loadEnvironmentConfig("config/config.api.json")
	if err != nil {
		return nil, fmt.Errorf("failed to load environment config: %w", err)
	}
	if Config.Environment != "development" && Config.Environment != "production" {
		return nil, fmt.Errorf("invalid environment: %s", Config.Environment)
	}
	if Config.Environment == "production" {
		dbconfig = Config.Production
	} else {
		dbconfig = Config.Development
	}
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbconfig.DBHost, dbconfig.DBPort, dbconfig.DBUser, dbconfig.DBName, dbconfig.DBPassword,
	)

	db, err := sql.Open(Driver, connectionString)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.Ping(); err != nil {
		defer db.Close()
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return db, nil

}
