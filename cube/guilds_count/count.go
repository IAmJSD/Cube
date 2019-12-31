package guildscount

import "github.com/jakemakesstuff/Cube/cube/redis"

// Count is used to count all of the guilds.
func Count() int64 {
	i, err := redis.Client.SCard("guilds").Result()
	if err != nil {
		panic(err)
	}
	return i
}
