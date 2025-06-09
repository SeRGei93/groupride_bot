package ride

import (
	"fmt"
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"
	"goupride_bot/internal/database/table"
	"goupride_bot/internal/services"
	"goupride_bot/internal/utils"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SaveDescription(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
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

	var files []table.File
	if update.Message.Photo != nil && len(update.Message.Photo) > 0 {
		photo := update.Message.Photo[len(update.Message.Photo)-1]
		files = append(files, table.File{
			ID:   photo.FileID,
			Type: "photo",
		})
	}

	description := update.Message.Text
	if len(files) > 0 {
		description = update.Message.Caption
		ride.Files = files // files
	}

	ride.Description = description
	ride.Ready = true

	err = db.Ride.UpdateRide(*ride)
	if err != nil {
		text := `Произошла ошибка`
		bot.Send(tgbotapi.NewMessage(chatID, text))
		log.Fatalf("failed attach time to new ride: %s", err)
		return
	}

	utils.DeleteAwaiting(userId)

	text := "✅ Спасибо! Заезд готов к публикации."
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "HTML"

	buttons, err := services.StartButtons(update.Message, db, cfg)
	if err == nil {
		msg.ReplyMarkup = buttons
	}

	bot.Send(msg)

	postMessage, err := createRidePost(ride, cfg)
	if err != nil {
		text := `Произошла ошибка`
		bot.Send(tgbotapi.NewMessage(chatID, text))
		log.Fatalf("failed create new post: %s", err)
		return
	}

	bot.Send(postMessage)
}

func createRidePost(ride *table.Ride, cfg config.Bot) (tgbotapi.Chattable, error) {
	if len(ride.Files) > 0 {
		fileID := ride.Files[0].ID
		photoMsg := tgbotapi.NewPhoto(cfg.Channel, tgbotapi.FileID(fileID))
		photoMsg.Caption = formatRideMessage(ride)
		photoMsg.ParseMode = "HTML"
		return photoMsg, nil
	} else {
		msg := tgbotapi.NewMessage(cfg.Channel, formatRideMessage(ride))
		msg.ParseMode = "HTML"
		return msg, nil
	}
}

func formatRideMessage(ride *table.Ride) string {
	date := ride.StartDate.Format("02.01.2006 в 15:04")
	return fmt.Sprintf(`%s

📍 <b>Маршрут:</b> <a href="%s">%s</a>

🕒 <b>Старт:</b> %s
`, ride.Description, ride.Link, ride.Link, date)
}
