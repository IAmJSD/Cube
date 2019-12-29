package commands

import (
	"fmt"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"time"
)

func init() {
	commandprocessor.Commands["ping"] = &commandprocessor.Command{
		Description:      "Pings the bot.",
		Category:          categories.INFORMATIONAL,
		Function: func(Args *commandprocessor.CommandArgs) {
			EndTime := time.Now()
			NowSubThen := EndTime.Sub(*Args.StartTime)
			Description := fmt.Sprintf(
				"WebSocket Latency: **%s**\nProcessing Latency: **%s**\n",
				Args.Session.HeartbeatLatency(),
				NowSubThen,
			)
			messages.GenericText(Args.Channel, "🏓 Pong!", Description, nil, Args.Session)
		},
	}
}
