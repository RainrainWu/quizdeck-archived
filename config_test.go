package quizdeck_test

import (
	"os"
	"testing"

	"github.com/RainrainWu/quizdeck"
	"github.com/stretchr/testify/assert"
)

func TestConfigSet(t *testing.T) {

	params := []struct {
		envName string
		actual  string
	}{
		{
			"DISCORD_AUTH_TOKEN",
			quizdeck.Config.GetDiscordAuthToken(),
		},
		{
			"DISCORD_APP_ID",
			quizdeck.Config.GetDiscordAppID(),
		},
		{
			"DISCORD_APP_PUBLIC_KEY",
			quizdeck.Config.GetDiscordAppPublicKey(),
		},
	}
	for _, param := range params {
		t.Run(
			param.envName,
			func(t *testing.T) {
				assert.Equal(t, os.Getenv(param.envName), param.actual)
			},
		)
	}
}
