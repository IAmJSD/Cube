package cacher

import (
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"github.com/jakemakesstuff/Cube/cube/redis"
)

// GetChannel is used to get a channel from the cache.
func GetChannel(ChannelID string, session *discordgo.Session) (*discordgo.Channel, error) {
	// Attempts to get from cache.
	b, err := redis.Client.Get("c:" + ChannelID).Bytes()
	if err == redis.Nil {
		// Attempts to get from Discord.
		c, err := session.Channel(ChannelID)
		if err != nil {
			return nil, err
		}
		j, err := json.Marshal(&c)
		if err != nil {
			return nil, err
		}
		redis.Client.Set("c:" + ChannelID, j, 0)
		return c, nil
	} else if err == nil {
		// Returns from cache.
		var Channel *discordgo.Channel
		err := json.Unmarshal(b, &Channel)
		if err != nil {
			return nil, err
		}
		return Channel, nil
	} else {
		// A unknown Redis error happened.
		return nil, err
	}
}

// DeleteChannels is used to delete channels from the cache.
func DeleteChannels(ChannelIDS ...string) {
	ModifiedIDS := make([]string, len(ChannelIDS))
	for i, v := range ChannelIDS {
		ModifiedIDS[i] = "c:" + v
	}
	redis.Client.Del(ModifiedIDS...)
}
