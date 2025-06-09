package ride

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"
	"goupride_bot/internal/services"
	"goupride_bot/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CancelCreate(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	chatID := update.CallbackQuery.Message.Chat.ID
	if chatID == cfg.PublicChat {
		return
	}

	sendCanceledMessage(bot, update, db, cfg)
	utils.DeleteAwaiting(update.CallbackQuery.From.ID)

	ride, err := db.Ride.FindNoReadyRideByUser(update.CallbackQuery.From.ID)
	if err != nil {
		return
	}

	db.Ride.DeleteRide(*ride)
}

func sendCanceledMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	text := `❌ Заезд удален`
	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	msg.ParseMode = "HTML"

	buttons, err := services.StartButtons(update.CallbackQuery.Message, db, cfg)
	if err == nil {
		msg.ReplyMarkup = buttons
	}

	bot.Send(msg)
}
