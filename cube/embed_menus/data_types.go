package embedmenus

import "github.com/bwmarrin/discordgo"

// MenuCache is used to represent all of the current menus.
var MenuCache = map[string]*EmbedMenu{}

// MenuInfo contains the information about the menu.
type MenuInfo struct {
	MenuID string
	Author string
	Info   []string
}

// MenuButton is the datatype containing information about the button.
type MenuButton struct {
	Emoji       string
	Name        string
	Description string
}

// MenuReaction represents the button and the function which it triggers.
type MenuReaction struct {
	Button   MenuButton
	Function func(ChannelID string, MessageID string, menu *EmbedMenu, client *discordgo.Session)
}

// MenuReactions are all of the reactions which the menu has.
type MenuReactions struct {
	ReactionSlice []MenuReaction
}

// EmbedMenu is the base menu.
type EmbedMenu struct {
	Reactions *MenuReactions
	parent    *EmbedMenu
	Embed     *discordgo.MessageEmbed
	MenuInfo  *MenuInfo
}
