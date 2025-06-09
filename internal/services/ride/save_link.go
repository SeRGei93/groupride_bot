package ride

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"
	"goupride_bot/internal/database/table"
	"goupride_bot/internal/services"
	"goupride_bot/internal/utils"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SaveLink(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	chatID := update.Message.Chat.ID
	userId := update.Message.From.ID
	if chatID == cfg.PublicChat {
		return
	}

	link := update.Message.Text

	err := db.Ride.CreateRide(table.Ride{
		UserID: userId,
		Active: true,
		Ready:  false,
		Link:   link,
	})

	if err != nil {
		text := `Произошла ошибка`
		bot.Send(tgbotapi.NewMessage(chatID, text))
		log.Fatalf("failed create new ride: %s", err)
		return
	}

	utils.NextAwaitRideStep(userId)

	text := `🕒 Введите дату и время старта в формате: 08.06.2025 12:30`
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "HTML"

	buttons, err := services.DisableButtons(update.Message, db, cfg)
	if err == nil {
		msg.ReplyMarkup = buttons
	}

	bot.Send(msg)
}
