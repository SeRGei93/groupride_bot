package handlers

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"
	"goupride_bot/internal/services"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Init(bot *tgbotapi.BotAPI, db database.Database, cfg config.Bot) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.CallbackQuery != nil {
			Callbacks(bot, update, db, cfg)
		} else if update.Message != nil && update.Message.IsCommand() {
			Commands(bot, update, db, cfg)
		} else if update.Message != nil && update.Message.NewChatMembers != nil {
			services.NewMember(bot, update, db, cfg)
		} else {
			Messages(bot, update, db, cfg)
		}
	}
}
