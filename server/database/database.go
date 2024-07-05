package database

import (
	"api/server/config"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	_ "github.com/lib/pq"
)

var (
	Driver = "postgres"

	APIConfigInfo config.APIConfig
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

func loadDatabase(filePath string) (config.DatabaseConfig, error) {
	var db config.DatabaseConfig
	file, err := os.Open(filePath)
	if err != nil {
		return db, fmt.Errorf("Failed to open config file: %w", err)
	}
	defer file.Close()
	err = json.NewDecoder(file).Decode(&db)
	if err != nil {
		return db, fmt.Errorf("Failed to decode config file: %w", err)
	}
	return db, nil
}

// Connection conecta com o banco de dados
func Connection() (*sql.DB, error) {

	dbConfig, err := loadDatabase("config/config.api.json")
	if err != nil {
		return nil, fmt.Errorf("failed to load database config: %w", err)
	}
	APIConfigInfo, err = loadAPIConfig("config/config.api.json")
	if err != nil {
		return nil, fmt.Errorf("failed to load API config: %w", err)
	}

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		dbConfig.DBHost, dbConfig.DBPort, dbConfig.DBUser, dbConfig.DBName, dbConfig.DBPassword,
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
