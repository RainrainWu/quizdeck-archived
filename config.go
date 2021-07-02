package quizdeck

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

var (
	Config ConfigSet = NewConfigSet()
)

type ConfigSet interface {
	GetDebugMode() bool
	GetDiscordAuthToken() string
	GetDiscordAppID() string
	GetDiscordAppPublicKey() string
}

type configSet struct {
	debugMode           string
	discordAuthToken    string
	discordAppID        string
	discordAppPublicKey string
}

func NewConfigSet() ConfigSet {
	err := godotenv.Load("/Users/rainwu/Repositories/quizdeck/.env")
	if err != nil {
		log.Printf("fail to load env vars: %v\n", err.Error())
	}
	instance := configSet{
		debugMode:           os.Getenv("DEBUG_MODE"),
		discordAuthToken:    os.Getenv("DISCORD_AUTH_TOKEN"),
		discordAppID:        os.Getenv("DISCORD_APP_ID"),
		discordAppPublicKey: os.Getenv("DISCORD_APP_PUBLIC_KEY"),
	}
	return &instance
}

func (c *configSet) GetDebugMode() bool {
	return strings.ToLower(c.debugMode) == "true"
}

func (c *configSet) GetDiscordAuthToken() string {
	return c.discordAuthToken
}

func (c *configSet) GetDiscordAppID() string {
	return c.discordAppID
}

func (c *configSet) GetDiscordAppPublicKey() string {
	return c.discordAppPublicKey
}
