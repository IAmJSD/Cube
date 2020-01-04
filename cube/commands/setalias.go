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
	commandprocessor.Commands["setalias"] = &commandprocessor.Command{
		Description:      "Allows administrators to set an alias to a command.",
		Usage:            "<alias> <command>",
		Category:         categories.ADMINISTRATOR,
		PermissionsCheck: permissions.ADMINISTRATOR,
		Function: func(Args *commandprocessor.CommandArgs) {
			// Gets the currency/split arguments.
			split := utils.SpaceSplit(Args.RawArgs)

			// Processes all needed arguments.
			if 2 > len(split) {
				messages.NotEnoughArgs(Args.Channel, Args.Session)
				return
			}
			alias := strings.ToLower(split[0])
			cmd := strings.ToLower(split[1])
			_, ok := commandprocessor.Commands[alias]
			if ok {
				messages.Error(Args.Channel, "Invalid alias:", "The alias cannot be an actual command.", Args.Session)
				return
			}
			_, ok = commandprocessor.Commands[cmd]
			if !ok {
				messages.Error(Args.Channel, "Invalid command:", "The command specified does not exist.", Args.Session)
				return
			}
			Aliases := aliases.GetAliases(Args.Channel.GuildID)
			_, ok = Aliases[alias]
			if ok {
				messages.Error(Args.Channel, "Invalid alias:", "The alias already exists.", Args.Session)
				return
			}
			redis.Client.SAdd("a:"+Args.Message.GuildID, alias+" "+cmd)

			// Send a confirmation.
			messages.GenericText(Args.Channel,
				"Alias set:", "`"+alias+"` is now a alias of `"+cmd+"`.", nil, Args.Session)
		},
	}
}
