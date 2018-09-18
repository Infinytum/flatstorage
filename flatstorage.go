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

func (fs *FlatStorage) resourcePath(collection string, resource string) string {
	return filepath.Clean(fs.collectionPath(collection), resource)
}

func (fs *FlatStorage) collectionPath(collection string) string {
	return filepath.Clean(fs.path, collection)
}

// NewFlatStorage opens a flatstorage at specified path
func NewFlatStorage(path string) (*FlatStorage, error) {

	// ToDo: Create DB on init
	return &FlatStorage{
		path:    filepath.Clean(path),
		mutex:   &sync.Mutex{},
		mutexes: make(map[string]*sync.Mutex),
		logger:  logrus.New(),
	}, nil
}
