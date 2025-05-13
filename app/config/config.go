package config

import (
	"os"
	"sync"
)

type Config struct {
	Server  Server
	DB      DBConfig
	Mail    Mail
	JWT     JWT
}

type DBConfig struct {
	DB_URL   string
}

type Server struct {
	Port string
}

type Mail struct{
	GmailPass string
	SenderEmail string
}

type JWT struct {
	Secret string
}

var (
	once   sync.Once
	config Config
)

func GetConfig() *Config {
	// goroutine実行中でも一度だけ実行される
	once.Do(func() {
		config.DB.DB_URL = getEnv("DATABASE_URL", "")

		// Server
		config.Server.Port = getEnv("SERVER_PORT", "")

		// Mail
		config.Mail.GmailPass = getEnv("GMAIL_APP_PASS", "")
		config.Mail.SenderEmail = getEnv("SENDER_EMAIL", "")

		// JWT
		config.JWT.Secret = getEnv("JWT_SECRET", "")
	})
	return &config
}

// getEnv は環境変数を取得し、指定されたデフォルト値を返す
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
