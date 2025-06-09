package ride

import (
	"fmt"
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"
	"goupride_bot/internal/database/table"
	"goupride_bot/internal/services"
	"goupride_bot/internal/utils"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartCreateRide(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	chatID := update.CallbackQuery.Message.Chat.ID
	if chatID == cfg.PublicChat {
		return
	}

	utils.SetAwaiting(update.CallbackQuery.From.ID, 600)

	text := `
<b>Создание группового заезда происходит в несколько этапов:</b>

1. Добавьте ссылку на карту с маршрутом 🗺
2. Укажите дату и время старта 🕒
3. Добавьте описание и фото, темп 📝📷

🔗 <b>Добавьте ссылку на маршрут</b>  
Поддерживаются: komoot.com, connect.garmin.com, ridewithgps.com, nakarte.me
`

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ParseMode = "HTML"

	buttons, err := services.DisableButtons(update.Message, db, cfg)
	if err == nil {
		msg.ReplyMarkup = buttons
	}

	bot.Send(msg)
}

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

	if len(ride.Files) > 0 {
		fileID := ride.Files[0].ID
		photoMsg := tgbotapi.NewPhoto(cfg.PublicChat, tgbotapi.FileID(fileID))
		photoMsg.Caption = FormatRideMessage(ride)
		photoMsg.ParseMode = "HTML"
		bot.Send(photoMsg)
	} else {
		msg := tgbotapi.NewMessage(cfg.PublicChat, FormatRideMessage(ride))
		msg.ParseMode = "HTML"
		bot.Send(msg)
	}
}

func FormatRideMessage(ride *table.Ride) string {
	date := ride.StartDate.Format("02.01.2006 в 15:04")
	return fmt.Sprintf(`%s

📍 <b>Маршрут:</b> <a href="%s">%s</a>

🕒 <b>Старт:</b> %s
`, ride.Description, ride.Link, ride.Link, date)
}
