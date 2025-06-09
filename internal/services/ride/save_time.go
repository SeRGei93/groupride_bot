package ride

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"
	"goupride_bot/internal/services"
	"goupride_bot/internal/utils"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SaveTime(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	chatID := update.Message.Chat.ID
	userId := update.Message.From.ID
	if chatID == cfg.PublicChat {
		return
	}

	ride, err := db.Ride.FindNoReadyRideByUser(userId)
	if err != nil {
		text := `Произошла ошибка`
		bot.Send(tgbotapi.NewMessage(chatID, text))
		log.Fatalf("failed attach time to new ride: %s", err)
		return
	}

	layout := "02.01.2006 15:04"
	parsedTime, err := time.Parse(layout, update.Message.Text)
	if err != nil {
		text := "⚠️ Неверный формат даты. Введите в виде: 08.06.2025 12:30"
		bot.Send(tgbotapi.NewMessage(chatID, text))
		return
	}

	ride.StartDate = parsedTime

	err = db.Ride.UpdateRide(*ride)
	if err != nil {
		text := `Произошла ошибка`
		bot.Send(tgbotapi.NewMessage(chatID, text))
		log.Fatalf("failed attach time to new ride: %s", err)
		return
	}

	utils.NextAwaitRideStep(userId)

	text := `✍️ Введите описание и прикрепите изображение 📷`
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "HTML"

	buttons, err := services.DisableButtons(update.Message, db, cfg)
	if err == nil {
		msg.ReplyMarkup = buttons
	}

	bot.Send(msg)
}
