package commands

import (
	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
	"github.com/jakemakesstuff/Cube/cube/categories"
	"github.com/jakemakesstuff/Cube/cube/command_processor"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/embed_menus"
	"github.com/jakemakesstuff/Cube/cube/permissions"
)

// CreateCurrencyMenu is used to create the currency config menu.
func CreateCurrencyMenu(MenuID string, GuildID string, msg *discordgo.Message, cur *currency.Currency) *embedmenus.EmbedMenu {
	// Creates the embed menu.
	Menu := embedmenus.NewEmbedMenu(
		discordgo.MessageEmbed{
			Title:       "Cube Currency Configuration",
			Description: "Using this menu, you can configure Cube's currency.",
			Color:       255,
		}, &embedmenus.MenuInfo{
			MenuID: MenuID,
			Author: msg.Author.ID,
			Info:   []string{},
		},
	)

	// Add a toggle for the currency.
	var EnableDisableCurrency string
	if cur.Enabled {
		EnableDisableCurrency = "Disable Currency"
	} else {
		EnableDisableCurrency = "Enable Currency"
	}
	Menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       "🔐",
			Name:        EnableDisableCurrency,
			Description: "Toggles the currency in this guild.",
		},
		Function: func(ChannelID string, MessageID string, menu *embedmenus.EmbedMenu, client *discordgo.Session) {
			cur.Enabled = !cur.Enabled
			_ = client.MessageReactionsRemoveAll(ChannelID, MessageID)
			defer CreateCurrencyMenu(MenuID, GuildID, msg, cur).Display(ChannelID, MessageID, client)
			currency.SaveCurrency(GuildID, cur)
		},
	})

	// Returns the embed menu.
	return &Menu
}

func init() {
	commandprocessor.Commands["currencyconfig"] = &commandprocessor.Command{
		Description:      "Used to configure the currency in this guild.",
		Category:         categories.CURRENCY,
		PermissionsCheck: permissions.ADMINISTRATOR,
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

			// Show the embed.
			CreateCurrencyMenu(MenuID, Args.Message.GuildID, Args.Message, cur).Display(Args.Channel.ID, m.ID, Args.Session)
		},
	}
}
