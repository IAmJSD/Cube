package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/role_shop"
)

func init() {
	commandprocessor.Commands["roleshop"] = &commandprocessor.Command{
		Description:      "Opens the role shop.",
		Category:         categories.CURRENCY,
		PermissionsCheck: currency.CurrencyEnabled,
		Function: func(Args *commandprocessor.CommandArgs) {
			// Gets the currency.
			cur := currency.GetCurrency(Args.Message.GuildID)

			// Gets the menu ID.
			MenuID := uuid.Must(uuid.NewRandom()).String()

			// Create the message.
			m, err := Args.Session.ChannelMessageSendComplex(Args.Message.ChannelID, &discordgo.MessageSend{
				Embed: &discordgo.MessageEmbed{
					Title: "Loading...",
				},
			})
			if err != nil {
				return
			}

			// Show the role shop.
			roleshop.ShowRoleShop(1, nil, cur, false, MenuID, Args.Message, Args.Session, m.ID)
		},
	}
}
