package wallets

import (
	"github.com/jakemakesstuff/Cube/cube/redis"
	"sync"
)

// ensureWallet is used to ensure that the user has a wallet.
func ensureWallet(GuildID string, UserID string) (string, *sync.Mutex) {
	// Gets the key/lock.
	Key := "b:" + GuildID + ":" + UserID
	Lock := getBalanceLock(Key)

	// Check if the key exists.
	Lock.Lock()
	Exists := redis.Client.Exists(Key).Val()
	if Exists == 0 {
		// This key doesn't exist - We should make it and add it to the guilds wallet set.
		redis.Client.SAdd("w:"+GuildID, UserID)
		redis.Client.Set(Key, "0", 0)
	}
	Lock.Unlock()

	// Returns the key/lock.
	return Key, Lock
}
