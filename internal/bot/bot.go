package bot

import (
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sergssnorth27/life_manager/internal/timer"
)

type TelegramBot struct {
	bot        *tgbotapi.BotAPI
	userStates map[int64]string
	tasks      map[int64][]Task
}

type Task struct {
	id   int
	text string
}

func (tg *TelegramBot) LoadBot(token string) error {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Printf("Не удалось создать бота: %v", err)
		return err
	}
	bot.Debug = true
	tg.bot = bot
	tg.userStates = make(map[int64]string)
	tg.tasks = make(map[int64][]Task)
	return nil
}

func (tg *TelegramBot) GetUpdates() {
	updateConf := tgbotapi.NewUpdate(0)
	updateConf.Timeout = 60
	updates := tg.bot.GetUpdatesChan(updateConf)
	for update := range updates {
		var chatId int64
		if update.Message != nil {
			chatId = update.Message.Chat.ID
			if tg.userStates[chatId] == "adding_task" {
				lastTaskId := len(tg.tasks[chatId])
				task := Task{
					id:   lastTaskId + 1,
					text: update.Message.Text,
				}
				tg.tasks[chatId] = append(tg.tasks[chatId], task)
				msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Задача добавлена: id : %v, %v", task.id, task.text))
				tg.bot.Send(msg)
				tg.userStates[chatId] = ""
			}

			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(chatId, "Привет 👋🏻")
				msg.ReplyMarkup = tgbotapi.NewReplyKeyboard(
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Запустить таймер ⏰"),
					),
					tgbotapi.NewKeyboardButtonRow(
						tgbotapi.NewKeyboardButton("Список задач 📝"),
					),
				)
				tg.bot.Send(msg)
			case "Список задач 📝":
				actionOptions := []struct {
					text string
					data string
				}{
					{"Завершить задачу", "complete_task"},
					{"Добавить задачу", "adding_task"},
					{"Удалить задачу", "delete_task"},
				}
				msgTasks := tgbotapi.NewMessage(chatId, "Список задач")
				var tasksInlineKeyboardRows [][]tgbotapi.InlineKeyboardButton
				for _, task := range tg.tasks[chatId] {
					tasksInlineKeyboardRows = append(tasksInlineKeyboardRows, tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(task.text, fmt.Sprintf("task_%v", task.id))))
				}
				msgTasks.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(tasksInlineKeyboardRows...)
				tg.bot.Send(msgTasks)

				msg := tgbotapi.NewMessage(chatId, "Выберите действие")
				var inlineKeyboardRows [][]tgbotapi.InlineKeyboardButton
				for _, option := range actionOptions {
					inlineKeyboardRows = append(inlineKeyboardRows, tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(option.text, option.data)))
				}
				msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(inlineKeyboardRows...)
				tg.bot.Send(msg)

			case "Запустить таймер ⏰":
				options := []struct {
					text string
					data string
				}{
					{"30 минут", "timer_30_min"},
					{"45 минут", "timer_45_min"},
					{"1 час", "timer_60_min"},
					{"1,5 часа", "timer_90_min"},
					{"2 часа", "timer_120_min"},
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Выберите время")
				var inlineKeyboardRows [][]tgbotapi.InlineKeyboardButton
				for _, option := range options {
					inlineKeyboardRows = append(inlineKeyboardRows, tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData(option.text, option.data)))
				}
				msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(inlineKeyboardRows...)
				tg.bot.Send(msg)
			}

		}

		if update.CallbackQuery != nil {
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, "")
			if _, err := tg.bot.Request(callback); err != nil {
				log.Printf("Ошибка при ответе на callback: %v", err)
			}
			chatId := update.CallbackQuery.Message.Chat.ID
			switch update.CallbackQuery.Data {
			case "adding_task":
				msg := tgbotapi.NewMessage(chatId, "Напишите задачу которую хотите добавить:")
				tg.userStates[chatId] = "adding_task"
				tg.bot.Send(msg)
			}

			dataToDuration := map[string]time.Duration{
				"timer_30_min":  30 * time.Minute,
				"timer_45_min":  45 * time.Minute,
				"timer_60_min":  60 * time.Minute,
				"timer_90_min":  90 * time.Minute,
				"timer_120_min": 120 * time.Minute,
			}
			if duration, ok := dataToDuration[update.CallbackQuery.Data]; ok {
				minutes := int(duration.Minutes())

				msg := tgbotapi.NewMessage(chatId, fmt.Sprintf("Запустил таймер на %v", minutes))
				tg.bot.Send(msg)
				timer.StartTimer(tg.bot, chatId, duration)
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
