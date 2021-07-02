package main

import (
	"bytes"
	"context"
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/mitchellh/mapstructure"

	"github.com/RainrainWu/quizdeck"
)

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

func validate(signature, timestamp string) bool {

	if signature == "" || timestamp == "" {
		fmt.Println("err input")
		return false
	}

	var msg bytes.Buffer
	pub, _ := hex.DecodeString(quizdeck.Config.GetDiscordAppPublicKey())
	pub = ed25519.PublicKey(pub)
	fmt.Println("decode pub: ", pub)
	sig, err := hex.DecodeString(signature)
	fmt.Println("decode sig: ", sig)
	if err != nil {
		return false
	}
	if len(sig) != ed25519.SignatureSize || sig[63]&224 != 0 {
		return false
	}

	msg.WriteString(timestamp)
	return ed25519.Verify(pub, msg.Bytes(), sig)
}

func loadInteraction(digest interface{}) DiscordInteraction {

	interaction := DiscordInteraction{}
	mapstructure.Decode(digest, &interaction)
	return interaction
}

func HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	fmt.Println("event: ", request)
	fmt.Println("sig: ", request.Headers["X-Signature-Ed25519"])
	fmt.Println("tsp: ", request.Headers["X-Signature-Timestamp"])
	if !validate(
		request.Headers["X-Signature-Ed25519"],
		request.Headers["X-Signature-Timestamp"],
	) {
		return events.APIGatewayProxyResponse{Body: "", StatusCode: 401}, nil
	}

	result := loadInteraction(request.Body)
	if result.Type == 1 {
		ack := Acknowledge{Type: 1}
		out, _ := json.Marshal(ack)
		return events.APIGatewayProxyResponse{Body: string(out), StatusCode: 200}, nil
	}
	return events.APIGatewayProxyResponse{
		Body:       fmt.Sprintf("current command: %s", result.Data.Options[0].Name),
		StatusCode: 200,
	}, nil
}

func main() {

	lambda.Start(HandleRequest)
}
