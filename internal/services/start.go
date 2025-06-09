package services

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"
	"goupride_bot/internal/database/table"
	"log"

	"gorm.io/gorm"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	if update.Message.Chat.ID == cfg.PublicChat {
		return
	}

	from := update.Message.From
	userID := from.ID

	_, err := db.User.FindUser(from.ID)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Fatalf("failed create new user: %s", err)
		return
	}

	if err != nil && err == gorm.ErrRecordNotFound {
		// Создание или обновление пользователя
		err = db.User.CreateUser(table.User{
			ID:        userID,
			NickName:  from.UserName,
			FirstName: from.FirstName,
			LastName:  from.LastName,
			Rides:     nil,
		})

		if err != nil {
			log.Fatalf("failed create new user: %s", err)
			return
		}
	}

	text := "👋 Привет! Я помогу тебе создать заезд, в который ты пригласишь своих друзей 🚴‍♀️🚴‍♂️"

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	buttons, err := StartButtons(update.Message, db, cfg)
	if err == nil {
		msg.ReplyMarkup = buttons
	}
	msg.ParseMode = "HTML"

	bot.Send(msg)
}
