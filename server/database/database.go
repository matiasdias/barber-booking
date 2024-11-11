package database

import (
	"api/server/config"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	Driver = "postgres"

	APIConfigInfo config.APIConfig
	Config        config.Config
	err           error
	dbconfig      config.DatabaseConfig
)

// Connection conecta com o banco de dados
func Connection() (*sql.DB, error) {

	APIConfigInfo, err = config.LoadPortConfig("config/config.api.json")
	if err != nil {
		return nil, fmt.Errorf("failed to load API config: %w", err)
	}
	Config, err = config.LoadEnvironmentConfig("config/config.api.json")
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
