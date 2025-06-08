package handlers

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"
	"goupride_bot/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	switch update.Message.Command() {
	case "start":
		services.Start(bot, update, db, cfg)
	case "export_csv":
		services.ExportCsv(bot, update, db, cfg)
	case "export_gifts":
		services.ExportGifts(bot, update, db, cfg)
	case "send_notify":
		services.SendNotify(bot, update, db, cfg)
	case "info":
		services.Info(bot, update, db, cfg)
	}
}
