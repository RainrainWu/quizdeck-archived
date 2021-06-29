package quizdeck_test

import (
	"os"
	"strconv"
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
			"DEBUG_MODE",
			strconv.FormatBool(quizdeck.Config.GetDebugMode()),
		},
		{
			"DISCORD_AUTH_TOKEN",
			quizdeck.Config.GetDiscordAuthToken(),
		},
		{
			"DISCORD_APP_ID",
			quizdeck.Config.GetDiscordAppID(),
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
