package commands

import (
	"github.com/jakemakesstuff/Cube/cube/aliases"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/permissions"
	"github.com/jakemakesstuff/Cube/cube/redis"
	"github.com/jakemakesstuff/Cube/cube/utils"
	"strings"
)

func init() {
	commandprocessor.Commands["rmalias"] = &commandprocessor.Command{
		Description:      "Allows administrators to remove an alias.",
		Usage:            "<alias>",
		Category:         categories.ADMINISTRATOR,
		PermissionsCheck: permissions.ADMINISTRATOR,
		Function: func(Args *commandprocessor.CommandArgs) {
			// Gets the currency/split arguments.
			split := utils.SpaceSplit(Args.RawArgs)

			// Processes all needed arguments.
			if 1 > len(split) {
				messages.NotEnoughArgs(Args.Channel, Args.Session)
				return
			}
			alias := strings.ToLower(split[0])
			a := aliases.GetAliases(Args.Channel.GuildID)
			cmd, ok := a[alias]
			if !ok {
				messages.Error(Args.Channel, "Invalid alias:", "The alias does not exist.", Args.Session)
				return
			}
			redis.Client.SRem("a:"+Args.Message.GuildID, alias+" "+cmd)

			// Send a confirmation.
			messages.GenericText(Args.Channel,
				"Alias deleted:", "The alias `"+alias+"` has been deleted.", nil, Args.Session)
		},
	}
}
