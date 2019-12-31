package guildscount

import "github.com/jakemakesstuff/Cube/cube/redis"

// Remove is used to delete the guild from the count.
func Remove(GuildID string) {
	redis.Client.SRem("guilds", GuildID)
}
