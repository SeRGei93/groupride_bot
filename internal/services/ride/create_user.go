package ride

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCreateRide(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	if update.Message.Chat.ID == cfg.PublicChat {
		return
	}

	text := `
	‼️ Создание группового заезда происходит в несколько этапов:
	1. Загрузите .gpx файл с треком
	2. Укажите время старта
	3. Введите описание и прикрепите картинку.
	
	Прикрепите .gpx файл`

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ParseMode = "HTML"

	bot.Send(msg)
}
