package handlers

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"
	"goupride_bot/internal/services/ride"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	cmd := update.CallbackQuery.Data

	switch cmd {
	case "create_ride":
		ride.StartCreateRide(bot, update, db, cfg)
	default:
		//msgType, id := utils.GetKeyValue(update.CallbackQuery.Data)
		//services.Kamni200(bot, update, db, cfg, bike)
	}
}
