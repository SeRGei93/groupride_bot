package services

import (
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func StartButtons(message *tgbotapi.Message, db database.Database, cfg config.Bot) (*tgbotapi.InlineKeyboardMarkup, error) {
	//userID := message.From.ID

	result := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("➕ Создать заезд", "create_ride")),
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonURL("🚴‍♂️ Канал со всеми заездами", "http://t.me/groupride_by")),
		//tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("📋 Мои заезды", fmt.Sprintf("user_events:%d", userID))),
	)

	return &result, nil
}

func DisableButtons(message *tgbotapi.Message, db database.Database, cfg config.Bot) (*tgbotapi.InlineKeyboardMarkup, error) {
	result := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(tgbotapi.NewInlineKeyboardButtonData("🚫 Отмена", "cancel")),
	)

	return &result, nil
}
