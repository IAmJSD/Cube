package wallets

import (
	"github.com/jakemakesstuff/Cube/cube/redis"
	"strconv"
	"sync"
)

// GetAll is used to get all of the guilds wallets.
func GetAll(GuildID string) map[string]int {
	// Gets all of the wallet keys/locks.
	Wallets := redis.Client.SMembers("w:" + GuildID).Val()
	Keys := make([]string, len(Wallets))
	Locks := make([]*sync.Mutex, len(Wallets))
	for i, v := range Wallets {
		Key := "b:" + GuildID + ":" + v
		Keys[i] = Key
		Locks[i] = getBalanceLock(Key)
	}

	// Lock all of the wallets.
	for _, v := range Locks {
		v.Lock()
	}

	// Get all of the wallets from the database.
	Map := map[string]int{}
	result, err := redis.Client.MGet(Keys...).Result()
	if err != nil {
		return Map
	}
	for i, v := range result {
		str, ok := v.(string)
		if ok {
			val, _ := strconv.Atoi(str)
			Map[Wallets[i]] = val
		}
	}

	// Unlock all of the wallets.
	for _, v := range Locks {
		v.Unlock()
	}

	// Return the map.
	return Map
}
