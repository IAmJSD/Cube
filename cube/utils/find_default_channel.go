package utils

import "github.com/bwmarrin/discordgo"

// FindDefaultChannel is used to find the default channel to post the first message.
func FindDefaultChannel(guild *discordgo.Guild) string {
	// Find the system channel ID.
	if guild.SystemChannelID != "" {
		return guild.SystemChannelID
	}

	// If the embed channel is not blank, this is a safe bet.
	if guild.WidgetChannelID != "" {
		for _, v := range guild.Channels {
			if v.ID == guild.WidgetChannelID && v.Type == discordgo.ChannelTypeGuildText {
				// This is a text channel.
				return v.ID
			}
		}
	}

	// Try and find the general chat.
	for _, v := range guild.Channels {
		if v.Name == "general" || v.Name == "general_chat" || v.Name == "mainchat" || v.Name == "chat" {
			// This is probably a main channel.
			return v.ID
		}
	}

	// Return a blank string.
	return ""
}
