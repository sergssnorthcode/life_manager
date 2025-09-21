package logs

import (
	"log"
	"os"
)

func LoadLogs() {
	file, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Ошибка открытия лог файла", err)
	}
	log.SetOutput(file)
}
