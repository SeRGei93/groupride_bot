package main

import (
	"goupride_bot/internal/clients"
	"goupride_bot/internal/config"
	"goupride_bot/internal/database"
	"goupride_bot/internal/handlers"
	"goupride_bot/internal/utils"
	"log/slog"
	"os"
	"time"

	"gorm.io/driver/sqlite"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log.Info("starting gravel bot")

	db := database.InitDatabase(sqlite.Open(cfg.StoragePath))

	bot := clients.InitBot(cfg.Bot)
	handlers.Init(bot, db, cfg.Bot)

	// очистка очереди ожидания
	go func() {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()
		for range ticker.C {
			utils.CleanupOldAwaiting()
		}
	}()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
