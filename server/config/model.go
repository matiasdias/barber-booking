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
	APIPort string `json:"api_port"`
}

type Credentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
}

type GoogleConfig struct {
	Oauth Credentials `json:"google_oauth"`
}

type Jwt struct {
	Secret     string `json:"secret"`
	ExpireTime int64  `json:"expiration_time"`
}

type JwtConfig struct {
	Jwt Jwt `json:"jwt_config"`
}
