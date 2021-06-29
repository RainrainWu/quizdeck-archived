package quizdeck

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var (
	Config ConfigSet = NewConfigSet()
)

type ConfigSet interface {
	GetDiscordAuthToken() string
	GetDiscordAppID() string
}

type configSet struct {
	discordAuthToken string
	discordAppID     string
}

func NewConfigSet() ConfigSet {
	err := godotenv.Load()
	if err != nil {
		log.Printf("fail to load env vars: %v\n", err.Error())
	}
	instance := configSet{
		discordAuthToken: os.Getenv("DISCORD_AUTH_TOKEN"),
		discordAppID:     os.Getenv("DISCORD_APP_ID"),
	}
	return &instance
}

func (c *configSet) GetDiscordAuthToken() string {
	return c.discordAuthToken
}

func (c *configSet) GetDiscordAppID() string {
	return c.discordAppID
}
