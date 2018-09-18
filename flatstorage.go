package flatstorage

import (
	"os"
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
	path := filepath.Join(fs.collectionPath(collection), resource)

	if filepath.Ext(path) != ".json" {
		path += ".json"
	}
	return path
}

func (fs *FlatStorage) collectionPath(collection string) string {
	return filepath.Join(fs.path, collection)
}

// NewFlatStorage opens a flatstorage at specified path
func NewFlatStorage(path string) (*FlatStorage, error) {
	path = filepath.Clean(path)

	if !pathExists(path) {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			logrus.WithField("path", path).Error("Could not create database", err)
			return nil, err
		}
	}

	return &FlatStorage{
		path:    filepath.Clean(path),
		mutex:   &sync.Mutex{},
		mutexes: make(map[string]*sync.Mutex),
		logger:  logrus.New(),
	}, nil
}
