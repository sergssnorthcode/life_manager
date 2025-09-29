package main

import (
	"flag"
	"log"

	"github.com/sergssnorth27/life_manager/config"
	"github.com/sergssnorth27/life_manager/internal/bot"
	"github.com/sergssnorth27/life_manager/internal/storage"
	"github.com/sergssnorth27/life_manager/logs"
)

func main() {
	configPath := flag.String("config", "config/config.dev.json", "Path to config")
	flag.Parse()
	conf, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Printf("Ошибка загрузки конфига %v", err)
	}
	logs.LoadLogs()

	db, err := storage.NewDB(conf.DbUrl)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer db.Close()

	var bot bot.TelegramBot

	err = bot.LoadBot(conf.TelegramBotToken, db)
	if err != nil {
		log.Printf("Ошибка загрузки бота: %v", err)
	}
	bot.GetUpdates()
}
