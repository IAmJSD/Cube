package commands

import (
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/permissions"
	"github.com/jakemakesstuff/Cube/cube/redis"
)

func init() {
	commandprocessor.Commands["prefix"] = &commandprocessor.Command{
		Description:      "This is used to set the prefix of the bot.",
		Category:         categories.ADMINISTRATOR,
		PermissionsCheck: permissions.ADMINISTRATOR,
		Usage:            "<new prefix>",
		Function: func(Args *commandprocessor.CommandArgs) {
			// Handles ensuring the prefix is correct.
			Prefix := ""
			for _, v := range Args.RawArgs {
				if v == ' ' {
					if Prefix == "" {
						// Just whitespace at the beginning.
						continue
					} else {
						// Whitespace cannot be in the prefix.
						messages.Error(Args.Channel, "Invalid Prefix", "Your prefix cannot have whitespace in it.", Args.Session)
						return
					}
				} else {
					// Add to the prefix.
					Prefix += string(v)
				}
			}

			// Sets the prefix.
			redis.Client.Set("p:"+Args.Message.GuildID, Prefix, 0)
			messages.GenericText(Args.Channel, "Prefix Set", "The prefix was successfully set to `"+Prefix+"`.", nil, Args.Session)
		},
	}
}
