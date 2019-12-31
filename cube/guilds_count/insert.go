package guildscount

import "github.com/jakemakesstuff/Cube/cube/redis"

// Insert is used to insert the guild into the count.
func Insert(GuildIDS ...interface{}) {
	redis.Client.SAdd("guilds", GuildIDS...)
}
