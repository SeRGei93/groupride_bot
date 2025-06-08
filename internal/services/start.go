package services

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Start(bot *tgbotapi.BotAPI, update tgbotapi.Update, db database.Database, cfg config.Bot) {
	if update.Message.Chat.ID == cfg.PublicChat {
		return
	}

	text := `
<b>КАМНИ 200 🔥 18+</b>
16–30 июня

Гравийный маршрут 200 км в формате гонки/бревета/покатухи — кому что ближе. Индивидуальное прохождение на условии самообеспечения. В зачёт принимаются страва-треки с окном прохождения с 16 по 30 июня 2025 года включительно.

<b>ПРИЗОВОЙ ФОНД</b> формируется самими участниками, претендовать на призы могут только те, кто сделал вклад. Ставить можно любые новые вещи на любое место либо условие. Например: ведро раков на 75 место, проездной на трамвай поломавшему велик, пачка минералки тому, кто потеряет сознание — дайте волю фантазии. Не обязательно ставить свою квартиру. Любой донат от участников приветствуется как дань уважения гравийному сообществу.

<b>ОБЯЗАТЕЛЬНОЕ СНАРЯЖЕНИЕ:</b> исправный велик, шлем, ремкомплект, питание, вода, навигация, аптечка, передний и задний свет.
Легенда маршрута: 70% неасфальтированная поверхность.

🗺 <b>МАРШРУТ:</b> <a href="https://ehai.club/kamni/Kamni200_2025_v1.gpx">GPX</a> | <a href="https://nakarte.me/#m=10/54.26482/27.30927&l=Y&nktl=JBZ7YVT6aBOO5xd2fESKEQ">Nakarte</a>
❗️До старта возможны изменения
	`

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	buttons, err := addButtons(update.Message, "kamni200", db, cfg)
	if err == nil {
		msg.ReplyMarkup = buttons
	}
	msg.ParseMode = "HTML"

	bot.Send(msg)
}
