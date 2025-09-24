package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	TelegramBotToken string
}

//config.dev.json
//config.prod.json

func LoadConfig(configPath string) (*Config, error) {
	file, err := os.Open(configPath)
	if err != nil {
		log.Printf("Ошибка открытия конфиг файла, %v", err)
		return nil, err
	}
	var cfg Config
	err = json.NewDecoder(file).Decode(&cfg)
	if err != nil {
		log.Printf("Ошибка декодирования конфиг файла, %v", err)
		return nil, err
	}
	return &cfg, nil
}
