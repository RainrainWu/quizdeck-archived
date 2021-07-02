package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchellh/mapstructure"
)

type Authorization struct {
	Signature string `json:"X-Signature-Ed25519"`
	Timestamp string `json:"X-Signature-Timestamp"`
}

type Acknowledge struct {
	Type int `json:"type"`
}

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
	return interaction
}

func HandleRequest(ctx context.Context, event interface{}) (string, error) {

	fmt.Printf("event: %v\n", event)
	result := loadInteraction(event)
	if result.Type == 1 {
		ack := Acknowledge{Type: 1}
		out, _ := json.Marshal(ack)
		return string(out), nil
	}
	return fmt.Sprintf("current command: %s", result.Data.Options[0].Name), nil
}

func main() {

	lambda.Start(HandleRequest)
}
