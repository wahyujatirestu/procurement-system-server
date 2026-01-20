package config

import (
	"errors"
	"os"
	"time"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type DBConfig struct {
	Host		string
	Port 		string
	User		string
	Password	string
	DBName		string
}

type APIConfig struct {
	Port	string
}

type JWTConfig struct {
	AppName				string
	JwtSignatureKey		[]byte
	JwtSigningMethod	*jwt.SigningMethodHMAC
	AccessTokenLifetime time.Duration
}

type Config struct {
	DB	DBConfig
	API APIConfig
	JWT JWTConfig
}

func NewConfig() (*Config, error) {
	_ = godotenv.Load()

	cfg := &Config{
		DB: DBConfig{
			Host: os.Getenv("DB_HOST"),
			Port: os.Getenv("DB_PORT"),
			User: os.Getenv("DB_USERNAME"),
			Password: os.Getenv("DB_PASSWORD"),
			DBName: os.Getenv("DB_NAME"),
		},
		API: APIConfig{
			Port: os.Getenv("API_PORT"),
		},
		JWT: JWTConfig{
			AppName: os.Getenv("JWT_APP_NAME"),
			JwtSignatureKey: []byte(os.Getenv("ACCESS_TOKEN")),
			JwtSigningMethod: jwt.SigningMethodHS256,
			AccessTokenLifetime: 24 * time.Hour,
		},
	}

	if cfg.DB.Host == "" || cfg.DB.Port == "" || cfg.DB.User == "" || cfg.DB.Password == "" || cfg.DB.DBName == "" {
		return nil, errors.New("invalid database configuration")
	}

	return cfg, nil
}