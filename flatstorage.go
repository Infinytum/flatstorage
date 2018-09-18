package flatstorage

import (
	"path/filepath"
	"sync"

	"github.com/sirupsen/logrus"
)

// FlatStorage is the interface to manage a flatstorage
type FlatStorage struct {
	// Storage changes must be made synchronous.
	mutex *sync.Mutex

	// Collection of mutexes per dataset.
	mutexes map[string]*sync.Mutex

	// File-system location of the filestorage
	path string

	// Logger which this storage will write to
	logger *logrus.Logger
}

// resourceExists checks if a resource is existent.
func (fs *FlatStorage) resourceExists(collection string, resource string) bool {
	return pathExists(filepath.Join(fs.path, collection, resource))

}

// resourceExists checks if a collection is existent.
func (fs *FlatStorage) collectionExists(collection string) bool {
	return pathExists(filepath.Join(fs.path, collection))
}

// resourceExists checks if a collection is existent.
func (fs *FlatStorage) databaseExists(collection string) bool {
	return pathExists(fs.path)
}

// NewFlatStorage opens a flatstorage at specified path
func NewFlatStorage(path string) (*FlatStorage, error) {
	return &FlatStorage{
		path:    filepath.Clean(path),
		mutex:   &sync.Mutex{},
		mutexes: make(map[string]*sync.Mutex),
		logger:  logrus.New(),
	}, nil
}
