package handlers

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"
	"goupride_bot/internal/services"
	"goupride_bot/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	cmd := update.CallbackQuery.Data

	switch cmd {
	case "rules":
		services.Rules(bot, update, db, cfg)
	case "kamni200":
		services.SetBike(bot, update, db, cfg)
	case "kamni200_off":
		services.Kamni200Off(bot, update, db, cfg)
	case "add_gift":
		services.AddGift(bot, update, db, cfg)
	default:
		_, bike := utils.GetKeyValue(update.CallbackQuery.Data)
		services.Kamni200(bot, update, db, cfg, bike)
	}
}
