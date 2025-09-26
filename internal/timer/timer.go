package timer

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartTimer(bot *tgbotapi.BotAPI, chatID int64, duration time.Duration) {
	go func() {
		time.Sleep(duration)
		msg := tgbotapi.NewMessage(chatID, "Пора отдохнуть ...")
		bot.Send(msg)
	}()
}
