package wallets

import "sync"

// balanceLocks is used to lock the balance while transactions go on.
var balanceLocks = map[string]*sync.Mutex{}

// balanceLocksLock is used to lock BalanceLocks if a balance lock needs to be made.
var balanceLocksLock = sync.RWMutex{}

// getBalanceLock is used to get the balance lock.
func getBalanceLock(Key string) *sync.Mutex {
	balanceLocksLock.RLock()
	lock, ok := balanceLocks[Key]
	balanceLocksLock.RUnlock()
	if !ok {
		lock = &sync.Mutex{}
		balanceLocksLock.Lock()
		balanceLocks[Key] = lock
		balanceLocksLock.Unlock()
	}
	return lock
}
