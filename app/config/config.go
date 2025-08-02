package config

import (
	"os"
	"sync"
)

type Config struct {
	Server      Server
	DB          DBConfig
	Mail        Mail
	JWT         JWT
	InviteGroup InviteGroup
}

type DBConfig struct {
	DB_HOST string
	DB_PORT string
	DB_USER string
	DB_PASSWORD string
	DB_NAME string
	ROOTPASS string
	DB_URL string
}

type Server struct {
	Port string
}

type Mail struct {
	GmailPass   string
	SenderEmail string
}

type JWT struct {
	Secret string
}

type InviteGroup struct {
	BaseURL string
}

var (
	once   sync.Once
	config Config
)

func GetConfig() *Config {
	// goroutine実行中でも一度だけ実行される
	once.Do(func() {
		// DB
		config.DB.DB_HOST = getEnv("DB_HOST", "")
		config.DB.DB_PORT = getEnv("DB_PORT", "")
		config.DB.DB_USER = getEnv("DB_USER", "")
		config.DB.DB_PASSWORD = getEnv("DB_PASSWORD", "")
		config.DB.DB_NAME = getEnv("DB_NAME", "")
		config.DB.ROOTPASS = getEnv("ROOTPASS", "")
		config.DB.DB_URL = getEnv("DATABASE_URL", "")

		// Server
		config.Server.Port = getEnv("SERVER_PORT", "")

		// Mail
		config.Mail.GmailPass = getEnv("GMAIL_APP_PASS", "")
		config.Mail.SenderEmail = getEnv("SENDER_EMAIL", "")

		// JWT
		config.JWT.Secret = getEnv("JWT_SECRET", "")

		// InviteGroup
		config.InviteGroup.BaseURL = getEnv("INVITE_BASEURL", "http://localhost:3000")
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
