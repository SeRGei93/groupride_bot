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
	}
}
