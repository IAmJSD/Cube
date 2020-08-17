package commands

import (
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/messages"
)

func init() {
	commandprocessor.Commands["privacypolicy"] = &commandprocessor.Command{
		Description: "Displays the privacy policy.",
		Category:    categories.INFORMATIONAL,
		Function: func(Args *commandprocessor.CommandArgs) {
			messages.GenericText(Args.Channel, "Privacy Policy", "We cache a very limited amount of information from Discord in order to be able to make the bot work properly. For example, we hold the relationship of certain ID's. We additionally hold all guild configurations and wallet information to ensure the bot can function as expected.\n\nIf you would like to request all data held or for the data to be removed, contact JakeMakesStuff#0001.", nil, Args.Session)
		},
	}
}
