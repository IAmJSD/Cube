package roleshop

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/embed_menus"
	"strconv"
)

// roleShopConfig is the config used for a role shop.
type roleShopConfig struct {
	role        *discordgo.Role
	description *string
	cost        *int
}

// addRole is the screen used to allow admins to add new roles to the bot.
func addRole(
	msg *discordgo.Message, MessageID string, MenuID string, session *discordgo.Session,
	Currency *currency.Currency, Config *roleShopConfig, ShopParent *embedmenus.EmbedMenu,
) {
	// Sets the description.
	Description := "**Role:** "
	if Config.role == nil {
		Description += "Unset\n"
	} else {
		Description += Config.role.Name + "\n"
	}
	Description += "**Description:** "
	if Config.description == nil {
		Description += "Unset\n"
	} else {
		Description += *Config.description + "\n"
	}
	Description += "**Cost:** "
	if Config.description == nil {
		Description += "Unset\n"
	} else {
		Emoji := "💵"
		if Currency.Emoji != nil {
			Emoji = *Currency.Emoji
		}
		Description += Emoji + strconv.Itoa(*Config.cost) + "\n"
	}

	// The main menu for adding a role.
	menu := embedmenus.NewEmbedMenu(discordgo.MessageEmbed{
		Title:       "Add Role",
		Description: Description,
	}, &embedmenus.MenuInfo{
		MenuID: MenuID,
		Author: msg.Author.ID,
		Info:   []string{},
	})

	// Handle the parent menus.
	menu.AddParentMenu(ShopParent)
	menu.AddBackButton()
	menu.Reactions.ReactionSlice[0].Button.Description += " **This will remove any unsaved changes.**"

	// Shows the menu.
	menu.Display(msg.ChannelID, MessageID, session)
}
