package config

import (
	"os"
	"sync"
	"log"
)

type Config struct {
	Server  Server
	DB      DBConfig
	Mailgun Mailgun
	JWT     JWT
}

type DBConfig struct {
	DB_URL   string
}

type Server struct {
	Port string
}

type Mailgun struct {
	Domain          string
	Private_Key     string
	Sender_email    string
	Recipient_email string
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
		log.Println("Initializing config...")
		config.DB.DB_URL = getEnv("DATABASE_URL", "")
		log.Println("DB_URL:", config.DB.DB_URL)

		// Server
		config.Server.Port = getEnv("SERVER_PORT", "")
		log.Println("SERVER_PORT:", config.Server.Port)

		// Mailgun
		config.Mailgun.Domain = getEnv("MAILGUN_DOMAIN", "")
		log.Println("MAILGUN_DOMAIN:", config.Mailgun.Domain)

		config.Mailgun.Private_Key = getEnv("MAILGUN_PRIVATE_API_KEY", "")
		log.Println("MAILGUN_PRIVATE_API_KEY:", config.Mailgun.Private_Key)

		config.Mailgun.Sender_email = getEnv("SENDER_EMAIL", "")
		log.Println("SENDER_EMAIL:", config.Mailgun.Sender_email)

		config.Mailgun.Recipient_email = getEnv("RECIPIENT_EMAIL", "")
		log.Println("RECIPIENT_EMAIL:", config.Mailgun.Recipient_email)

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
