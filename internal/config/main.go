package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	Bot         Bot
}

type Bot struct {
	TelegramToken string  `yaml:"token"`
	UseWebhook    bool    `yaml:"use_webhook"`
	WebhookURL    string  `yaml:"webhook_url"`
	AdminUsers    []int64 `yaml:"admin_users"`
	BanUsers      []int64 `yaml:"ban_users"`
	PublicChat    int64   `yaml:"public_chat"`
}

var (
	PatchFlag string
)

func MustLoad() *Config {
	flag.StringVar(&PatchFlag, "config", getEnv("CONFIG_PATH", ""), "config file path")
	flag.Parse()

	configPath := PatchFlag
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}

func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
