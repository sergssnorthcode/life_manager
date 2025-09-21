package main

import (
	"log"

	"github.com/sergssnorth27/life_manager/config"
	"github.com/sergssnorth27/life_manager/logs"
)

func main() {
	conf := config.LoadConfig()
	logs.LoadLogs()
	log.Println(conf.TelegramBotToken)

}
