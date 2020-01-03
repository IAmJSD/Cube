package roleshop

import (
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/currency"
	"github.com/jakemakesstuff/Cube/cube/embed_menus"
	"github.com/jakemakesstuff/Cube/cube/message_waiter"
	"github.com/jakemakesstuff/Cube/cube/styles"
	"strconv"
	"strings"
)

// roleShopConfig is the config used for a role shop.
type roleShopConfig struct {
	role        *discordgo.Role
	description *string
	cost        *int
	trial       bool
}

// addRole is the screen used to allow admins to add new roles to the bot.
func addRole(
	msg *discordgo.Message, MessageID string, MenuID string, session *discordgo.Session,
	Currency *currency.Currency, Config *roleShopConfig, ShopParent *embedmenus.EmbedMenu,
	Parent *embedmenus.EmbedMenu, Page int,
) {
	// Function to render the parent role shop.
	renderParentRoleShop := func() {
		ShowRoleShop(Page, Parent, Currency, true, MenuID, msg, session, MessageID)
	}

	// Function to render this embed.
	renderThisEmbed := func() {
		addRole(msg, MessageID, MenuID, session, Currency, Config, ShopParent, Parent, Page)
	}

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
	if Config.cost == nil {
		Description += "Unset\n"
	} else {
		Emoji := "💵"
		if Currency.Emoji != nil {
			Emoji = *Currency.Emoji
		}
		Description += Emoji + " " + strconv.Itoa(*Config.cost) + "\n"
	}
	RoleTrialsAllowed := "False"
	if Config.trial {
		RoleTrialsAllowed = "True"
	}
	Description += "**Role Trials Allowed:** " + RoleTrialsAllowed + "\n"

	// Handles if it can save.
	CanSave := Config.role != nil && Config.cost != nil && Config.description != nil
	if !CanSave {
		Description += "\n**When all unset values are set, you will have an option to save this.**"
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

	// Gets the role.
	menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       "👥",
			Name:        "Get Role",
			Description: "Gets the role for this config.",
		},
		Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
			// Remove all reactions.
			_ = session.MessageReactionsRemoveAll(msg.ChannelID, MessageID)

			// Get the role.
			_, _ = session.ChannelMessageEditComplex(&discordgo.MessageEdit{
				Content: nil,
				Embed: &discordgo.MessageEmbed{
					Title:       "Enter the role name:",
					Description: "Please type the name of the role.",
					Color:       styles.Generic,
				},
				ID:      MessageID,
				Channel: msg.ChannelID,
			})
			rolemsg := messagewaiter.WaitForMessage(msg.ChannelID, msg.Author.ID, 0)
			lower := strings.Trim(strings.ToLower(rolemsg.Content), " ")
			guild, _ := session.Guild(msg.GuildID)
			for _, v := range guild.Roles {
				if lower == strings.ToLower(v.Name) {
					// It's this role!
					Config.role = v
					break
				}
			}

			// Delete the message.
			_ = session.ChannelMessageDelete(msg.ChannelID, rolemsg.ID)

			// Redraw this embed.
			renderThisEmbed()
		},
	})

	// Sets the role description.
	menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       "💬",
			Name:        "Set Description",
			Description: "Sets a description (maximum 100 characters) of the role.",
		},
		Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
			// Remove all reactions.
			_ = session.MessageReactionsRemoveAll(msg.ChannelID, MessageID)

			// Get the role.
			_, _ = session.ChannelMessageEditComplex(&discordgo.MessageEdit{
				Content: nil,
				Embed: &discordgo.MessageEmbed{
					Title:       "Enter the description:",
					Description: "Please type the description of the role.",
					Color:       styles.Generic,
				},
				ID:      MessageID,
				Channel: msg.ChannelID,
			})
			description := messagewaiter.WaitForMessage(msg.ChannelID, msg.Author.ID, 0)
			d := description.Content
			if len(d) > 100 {
				d = description.Content[:100]
			}
			Config.description = &d

			// Delete the message.
			_ = session.ChannelMessageDelete(msg.ChannelID, description.ID)

			// Redraw this embed.
			renderThisEmbed()
		},
	})

	// Sets the cost.
	menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       "💰",
			Name:        "Set Cost",
			Description: "Sets the cost of the role.",
		},
		Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
			// Remove all reactions.
			_ = session.MessageReactionsRemoveAll(msg.ChannelID, MessageID)

			// Get the role.
			_, _ = session.ChannelMessageEditComplex(&discordgo.MessageEdit{
				Content: nil,
				Embed: &discordgo.MessageEmbed{
					Title:       "Enter the cost:",
					Description: "Please enter the cost as a number **with no other contents in the message**.",
					Color:       styles.Generic,
				},
				ID:      MessageID,
				Channel: msg.ChannelID,
			})
			possibleint := messagewaiter.WaitForMessage(msg.ChannelID, msg.Author.ID, 0)
			i, err := strconv.Atoi(possibleint.Content)
			if err == nil {
				// Make sure the cost is greater than 0.
				if 0 >= i {
					i = 1
				}

				// Set the cost.
				Config.cost = &i
			}

			// Delete the message.
			_ = session.ChannelMessageDelete(msg.ChannelID, possibleint.ID)

			// Redraw this embed.
			renderThisEmbed()
		},
	})

	// Toggle trials.
	var EnableDisableTrials string
	var EnabledEmoji string
	if Config.trial {
		EnableDisableTrials = "Disable Role Trials"
		EnabledEmoji = "❗"
	} else {
		EnableDisableTrials = "Enable Role Trials"
		EnabledEmoji = "✅"
	}
	menu.Reactions.Add(embedmenus.MenuReaction{
		Button: embedmenus.MenuButton{
			Emoji:       EnabledEmoji,
			Name:        EnableDisableTrials,
			Description: "Toggle role trials on this role.",
		},
		Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
			// Remove all reactions.
			_ = session.MessageReactionsRemoveAll(msg.ChannelID, MessageID)

			// Toggle role trials.
			Config.trial = !Config.trial

			// Redraw this embed.
			renderThisEmbed()
		},
	})

	// Add a save option if it can save.
	if CanSave {
		menu.Reactions.Add(embedmenus.MenuReaction{
			Button: embedmenus.MenuButton{
				Emoji:       "💾",
				Name:        "Save",
				Description: "Saves the configured role.",
			},
			Function: func(_ string, _ string, _ *embedmenus.EmbedMenu, _ *discordgo.Session) {
				// Remove all reactions.
				_ = session.MessageReactionsRemoveAll(msg.ChannelID, MessageID)

				// Save this to the shop.
				Currency.RoleShop = append(Currency.RoleShop, &currency.BuyableRole{
					Amount:       *Config.cost,
					RoleID:       Config.role.ID,
					Description:  *Config.description,
					TrialAllowed: Config.trial,
				})
				currency.SaveCurrency(msg.GuildID, Currency)

				// Redraw the parent role shop.
				renderParentRoleShop()
			},
		})
	}

	// Shows the menu.
	menu.Display(msg.ChannelID, MessageID, session)
}
