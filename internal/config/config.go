package config

import (
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfighcl"
	"log"
	"sync"
	"time"
)

type Config struct {
	TelegramBotToken     string        `hcl:"telegram_bot_token" env:"TELEGRAM_BOT_TOKEN" required:"true"`
	TelegramChannelID    int64         `hcl:"telegram_channel_id" env:"TELEGRAM_CHANNEL_ID" required:"true"`
	DatabaseDSN          string        `hcl:"database_dsn" env:"DATABASE_DSN" default:"postgres://postgres:postgres@localhost:5432/snakefishAnekdot?sslmode=disable"`
	FetchInterval        time.Duration `hcl:"fetch_interval" env:"FETCH_INTERVAL" default:"1440m"`
	NotificationInterval time.Duration `hcl:"notification_interval" env:"NOTIFICATION_INTERVAL" default:"60m"`
	FilterKeywords       []string      `hcl:"filter_keywords" env:"FILTER_KEYWORDS"`
	OpenAIKey            string        `hcl:"openai_key" env:"OPENAI_KEY"`
	OpenAIPrompt         string        `hcl:"openai_prompt" env:"OPENAI_PROMPT"`
	OpenAIModel          string        `hcl:"openai_model" env:"OPENAI_MODEL" default:"gpt-3.5-turbo"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() Config {
	once.Do(func() {
		loader := aconfig.LoaderFor(&cfg, aconfig.Config{
			EnvPrefix: "SFNKDT",
			Files:     []string{"./config.hcl", "./config.local.hcl", "$HOME/.config/snakefishAnekdot/config.hcl"},
			FileDecoders: map[string]aconfig.FileDecoder{
				".hcl": aconfighcl.New(),
			},
		})

		if err := loader.Load(); err != nil {
			log.Printf("[ERROR] failed to load config: %v", err)
		}
	})

	return cfg
}
