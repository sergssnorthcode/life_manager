package bot

import (
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sergssnorth27/life_manager/internal/timer"
)

type TelegramBot struct {
	bot *tgbotapi.BotAPI
}

func (tg *TelegramBot) LoadBot(token string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("Не удалось создать бота: %v", err)
		return err
	}
	bot.Debug = true
	tg.bot = bot
	return nil
}

func (tg *TelegramBot) GetUpdates() {
	updateConf := tgbotapi.NewUpdate(0)
	updateConf.Timeout = 60
	updates, err := tg.bot.GetUpdatesChan(updateConf)
	if err != nil {
		log.Printf("Не получилось получить обновления: %v", err)
	}
	for update := range updates {
		if update.Message != nil {
			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет 👋🏻")
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Запустить таймер ⏰"),
					),
				)
				tg.bot.Send(msg)
			case "Запустить таймер ⏰":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите время")
				msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("30 минут", "timer_30_min"),
					),
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("45 минут", "timer_45_min"),
					),
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("60 минут", "timer_60_min"),
					),
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("2 часа", "timer_120_min"),
					),
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("2 часа", "timer_30_sec"),
					),
				)
				tg.bot.Send(msg)
			}

		}

		if update.CallbackQuery != nil {
			switch update.CallbackQuery.Data {
			case "timer_30_min":
				duration := 30 * time.Second
				timer.StartTimer(tg.bot, update.CallbackQuery.Message.Chat.ID, duration)
			case "timer_30_sec":
				duration := 30 * time.Second
				timer.StartTimer(tg.bot, update.CallbackQuery.Message.Chat.ID, duration)

			}
		}
	}
}

// Reply keyboard
// msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите действие:")
// msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("Просто текст"),
// 		tgbotapi.NewKeyboardButtonContact("Отправить контакт"),
// 		tgbotapi.NewKeyboardButtonLocation("Отправить локацию"),
// 	),
// 	tgbotapi.NewKeyboardButtonRow(
// 		tgbotapi.NewKeyboardButton("Просто текст2"),
// 	),

// Reply keyboard
// inlineMsg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите вариант:")
// inlineMsg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonData("Кнопка с callback", "my_callback"),
// 		tgbotapi.NewInlineKeyboardButtonURL("Ссылка", "https://example.com"),
// 	),
// )
// tg.bot.Send(inlineMsg)
