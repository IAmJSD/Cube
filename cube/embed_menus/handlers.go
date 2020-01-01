package embedmenus

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

// Add is used to add a menu reaction.
func (mr *MenuReactions) Add(reaction MenuReaction) {
	Slice := append(mr.ReactionSlice, reaction)
	mr.ReactionSlice = Slice
}

// Display is used to show a menu.
func (emm EmbedMenu) Display(ChannelID string, MessageID string, client *discordgo.Session) *error {
	MenuCache[emm.MenuInfo.MenuID] = &emm

	EmbedCopy := emm.Embed
	EmbedCopy.Footer = &discordgo.MessageEmbedFooter{
		Text: emm.MenuInfo.MenuID,
	}
	Fields := make([]*discordgo.MessageEmbedField, 0)
	for _, k := range emm.Reactions.ReactionSlice {
		Fields = append(Fields, &discordgo.MessageEmbedField{
			Name:   fmt.Sprintf("%s %s", k.Button.Emoji, k.Button.Name),
			Value:  k.Button.Description,
			Inline: false,
		})
	}
	EmbedCopy.Fields = Fields

	_, err := client.ChannelMessageEditComplex(&discordgo.MessageEdit{
		Embed:   EmbedCopy,
		ID:      MessageID,
		Channel: ChannelID,
	})
	if err != nil {
		return &err
	}
	for _, k := range emm.Reactions.ReactionSlice {
		err := client.MessageReactionAdd(ChannelID, MessageID, k.Button.Emoji)
		if err != nil {
			return &err
		}
	}
	return nil
}

// NewChildMenu is used to create a new child menu.
func (emm EmbedMenu) NewChildMenu(embed discordgo.MessageEmbed, item MenuButton) *EmbedMenu {
	NewEmbedMenu := NewEmbedMenu(embed, emm.MenuInfo)
	NewEmbedMenu.parent = &emm
	Reaction := MenuReaction{
		Button: item,
		Function: func(ChannelID string, MessageID string, _ *EmbedMenu, client *discordgo.Session) {
			_ = client.MessageReactionsRemoveAll(ChannelID, MessageID)
			NewEmbedMenu.Display(ChannelID, MessageID, client)
		},
	}
	emm.Reactions.Add(Reaction)
	return &NewEmbedMenu
}

// AddBackButton is used to add a back Button to the page.
func (emm EmbedMenu) AddBackButton() {
	Reaction := MenuReaction{
		Button: MenuButton{
			Description: "Goes back to the parent menu.",
			Name:        "Back",
			Emoji:       "⬆",
		},
		Function: func(ChannelID string, MessageID string, _ *EmbedMenu, client *discordgo.Session) {
			_ = client.MessageReactionsRemoveAll(ChannelID, MessageID)
			emm.parent.Display(ChannelID, MessageID, client)
		},
	}
	emm.Reactions.Add(Reaction)
}

// NewEmbedMenu is used to create a new menu handler.
func NewEmbedMenu(embed discordgo.MessageEmbed, info *MenuInfo) EmbedMenu {
	var reactions []MenuReaction
	menu := EmbedMenu{
		Reactions: &MenuReactions{
			ReactionSlice: reactions,
		},
		Embed:    &embed,
		MenuInfo: info,
	}
	return menu
}

// HandleMenuReactionEdit is used to handle menu reactions.
func HandleMenuReactionEdit(client *discordgo.Session, reaction *discordgo.MessageReactionAdd, MenuID string) {
	_ = client.MessageReactionRemove(reaction.ChannelID, reaction.MessageID, reaction.Emoji.Name, reaction.UserID)
	menu := MenuCache[MenuID]
	if menu == nil {
		return
	}

	if menu.MenuInfo.Author != reaction.UserID {
		return
	}

	for _, v := range menu.Reactions.ReactionSlice {
		if v.Button.Emoji == reaction.Emoji.Name {
			v.Function(reaction.ChannelID, reaction.MessageID, menu, client)
			return
		}
	}
}
