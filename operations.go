package flatstorage

import (
	"fmt"
	"path/filepath"
)

// Delete removes a single resource from a collection
func (fs *FlatStorage) Delete(collection string, resource string) error {
	if fs.Exists(collection, resource) {
		return fmt.Errorf("Not implemented")
	}
	return nil
}

// DeleteAll removes an entire collection from the filesystem
func (fs *FlatStorage) DeleteAll(collection string) error {
	if fs.CollectionExists(collection) {
		return fmt.Errorf("Not implemented")
	}
	return nil
}

// Exists checks if a resource is present in a collection
func (fs *FlatStorage) Exists(collection string, resource string) bool {
	return pathExists(filepath.Join(fs.path, collection, resource))
}

// CollectionExists checks if a collection exists
func (fs *FlatStorage) CollectionExists(collection string) bool {
	return pathExists(filepath.Join(fs.path, collection))
}

// Read reads a single resource from a collection into an interface instance
func (fs *FlatStorage) Read(collection string, resource string, out interface{}) error {
	return fmt.Errorf("Not implemented")
}

// ReadAll reads all resources from a collection into an interface array of resourceType
func (fs *FlatStorage) ReadAll(collection string, resourceType interface{}) ([]interface{}, error) {
	return nil, fmt.Errorf("Not implemented")
}

// Write writes a single object into a collection from an interface instance
func (fs *FlatStorage) Write(collection string, resource string, object interface{}) error {
	return fmt.Errorf("Not implemented")
}
