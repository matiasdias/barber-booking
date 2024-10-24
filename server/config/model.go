package config

type DatabaseConfig struct {
	DBName     string `json:"db_name"`
	DBUser     string `json:"db_user"`
	DBPassword string `json:"db_password"`
	DBHost     string `json:"db_host"`
	DBPort     string `json:"db_port"`
}

type Config struct {
	Environment string         `json:"environment"`
	Development DatabaseConfig `json:"development"`
	Production  DatabaseConfig `json:"production"`
}

type APIConfig struct {
	APIPort int `json:"api_port"`
}

type AuthConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
}

type Jwt struct {
	Secret     string `json:"secret"`
	ExpireTime int64  `json:"expiration_time"`
}

type JwtConfig struct {
	Jwt Jwt `json:"jwt_config"`
}
