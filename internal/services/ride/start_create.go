package ride

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"
	"goupride_bot/internal/services"
	"goupride_bot/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCreate(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	chatID := update.CallbackQuery.Message.Chat.ID
	if chatID == cfg.PublicChat {
		return
	}

	utils.SetAwaiting(update.CallbackQuery.From.ID, 600)

	text := `
<b>Создание группового заезда происходит в несколько этапов:</b>

1. Добавьте ссылку на карту с маршрутом 🗺
2. Укажите дату и время старта 🕒
3. Добавьте описание и фото, темп 📝📷

🔗 <b>Добавьте ссылку на маршрут</b>  
Поддерживаются: komoot.com, connect.garmin.com, ridewithgps.com, nakarte.me
`

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "HTML"

	buttons, err := services.DisableButtons(update.Message, db, cfg)
	if err == nil {
		msg.ReplyMarkup = buttons
	}

	bot.Send(msg)
}
