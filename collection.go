package flatstorage

import "sync"

// LockCollection locks down a specific collection to thread-safe behaviour
func (fs *FlatStorage) LockCollection(collection string) {
	mutex := fs.getOrCreateCollectionMutex(collection)
	mutex.Lock()
}

// UnlockCollection unlocks a locked collection to be accessible by threads again
func (fs *FlatStorage) UnlockCollection(collection string) {
	mutex := fs.getOrCreateCollectionMutex(collection)
	mutex.Unlock()
}

// getOrCreateCollectionMutex creates a new collection specific mutex
func (fs *FlatStorage) getOrCreateCollectionMutex(collection string) *sync.Mutex {
	// Mutex modifications must be thread-safe
	defer fs.mutex.Unlock()
	fs.mutex.Lock()

	// Create a new mutex if none exists
	if _, exists := fs.mutexes[collection]; !exists {
		fs.mutexes[collection] = &sync.Mutex{}
	}

	return fs.mutexes[collection]
}
