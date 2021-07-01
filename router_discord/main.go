package lab

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchellh/mapstructure"

	"github.com/RainrainWu/quizdeck"
)

type DiscordInteraction struct {
	Type   int    `json:"type"`
	Token  string `json:"token"`
	Member struct {
		User struct {
			ID            int64  `json:"id"`
			Username      string `json:"username"`
			Avatar        string `json:"avatar"`
			Discriminator string `json:"discriminator"`
			PublicFlags   int    `json:"public_flags"`
		} `json:"user"`
		Roles        []string    `json:"roles"`
		PremiumSince interface{} `json:"premium_since"`
		Permissions  string      `json:"permissions"`
		Pending      bool        `json:"pending"`
		Nick         interface{} `json:"nick"`
		Mute         bool        `json:"mute"`
		JoinedAt     time.Time   `json:"joined_at"`
		IsPending    bool        `json:"is_pending"`
		Deaf         bool        `json:"deaf"`
	} `json:"member"`
	ID      string `json:"id"`
	GuildID string `json:"guild_id"`
	Data    struct {
		Options []struct {
			Name  string `json:"name"`
			Value string `json:"value"`
		} `json:"options"`
		Name string `json:"name"`
		ID   string `json:"id"`
	} `json:"data"`
	ChannelID string `json:"channel_id"`
}

func loadInteraction(digest interface{}) DiscordInteraction {

	interaction := DiscordInteraction{}
	mapstructure.Decode(digest, &interaction)
	return integration
}

func HandleRequest(ctx context.Context, event interface{}) (string, error) {

	result := loadInteraction(event)
	fmt.Println("parse result: %v\n", result)
	return fmt.Sprintf("current debug mode: %v", quizdeck.Config.GetDebugMode()), nil
}

func main() {

	lambda.Start(HandleRequest)
}
