package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Site struct {
		URL        string `yaml:"url"`
		DateFormat string `yaml:"date_format"`
		Selectors  struct {
			EventRoot string `yaml:"event_root"`
			Title     string `yaml:"title"`
			DateAttr  string `yaml:"date_attr"`
		} `yaml:"selectors"`
	} `yaml:"site"`
	Telegram struct {
		BotToken string
		ChatIDs  []string `yaml:"chat_ids"`
	} `yaml:"telegram"`
}

func LoadConfig(ymlPath string) (*Config, error) {
	data, err := os.ReadFile(ymlPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	tryLoadDotEnv()

	config.Telegram.BotToken = getTelegramBotToken()

	return &config, nil
}

func tryLoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables directly")
	}
}

func getTelegramBotToken() string {
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	if botToken == "" {
		log.Fatal("TELEGRAM_BOT_TOKEN is not set")
	}

	return botToken
}
