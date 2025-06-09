package handlers

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"

	"goupride_bot/internal/services/ride"
	"goupride_bot/internal/utils"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Messages(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	if update.Message != nil && update.Message.Chat.IsPrivate() {
		if utils.IsAwaiting(update.Message.From.ID) {

			step, err := utils.GetAwaitRideStep(update.Message.From.ID)
			if err != nil {
				return
			}

			switch step {
			case utils.StepLink:
				ride.SaveLink(bot, update, db, cfg)
			case utils.StepTime:
				ride.SaveTime(bot, update, db, cfg)
			case utils.StepDescription:
				ride.SaveDescription(bot, update, db, cfg)
			}
		}
	}
}
