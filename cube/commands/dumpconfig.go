package commands

import (
	"bytes"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/messages"
	"github.com/jakemakesstuff/Cube/cube/permissions"
	"github.com/jakemakesstuff/Cube/cube/redis"
	"github.com/jakemakesstuff/Cube/cube/wallets"
)

func init() {
	commandprocessor.Commands["dumpconfig"] = &commandprocessor.Command{
		Description:      "Allows you to dump a guilds configuration.",
		Category:         categories.ADMINISTRATOR,
		PermissionsCheck: permissions.ADMINISTRATOR,
		Function: func(Args *commandprocessor.CommandArgs) {
			// Dumps the config.
			var Prefix *string
			p, err := redis.Client.Get("p:" + Args.Message.GuildID).Result()
			if err == nil {
				Prefix = &p
			}
			Config := map[string]interface{}{
				"Wallets":        wallets.GetAll(Args.Message.GuildID),
				"CurrencyConfig": currency.GetCurrency(Args.Message.GuildID),
				"Prefix":         Prefix,
			}

			// Marshal the config into JSON.
			ConfigBytes, err := json.MarshalIndent(&Config, "", "  ")
			if err != nil {
				// TODO: Report to Sentry!
				return
			}

			// DM the config.
			c, err := Args.Session.UserChannelCreate(Args.Message.Author.ID)
			if err != nil {
				messages.Error(Args.Channel, "Failed to DM:", "Failed to DM you! Do you have DM's off or me blocked?", Args.Session)
				return
			}
			_, err = Args.Session.ChannelMessageSendComplex(c.ID, &discordgo.MessageSend{File: &discordgo.File{
				Name:        "config.json",
				ContentType: "application/json",
				Reader:      bytes.NewReader(ConfigBytes),
			}})
			if err != nil {
				messages.Error(Args.Channel, "Failed to DM:", "Failed to DM you! Do you have DM's off or me blocked?", Args.Session)
				return
			}

			// Say that I DM'd.
			messages.GenericText(Args.Channel, "DM'd config:", "I have DM'd you the guild config!", nil, Args.Session)
		},
	}
}
