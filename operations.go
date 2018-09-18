package flatstorage

import (
	"fmt"
)

// Delete removes a single object from a collection
func (fs *FlatStorage) Delete(collection string, objectID string) error {
	return fmt.Errorf("Not implemented")
}

// DeleteAll removes an entire collection from the filesystem
func (fs *FlatStorage) DeleteAll(collection string) error {
	return fmt.Errorf("Not implemented")
}

// Exists checks if an object is present in a collection
func (fs *FlatStorage) Exists(collection string, objectID string) bool {
	return false
}

// CollectionExists checks if a collection exists
func (fs *FlatStorage) CollectionExists(collection string) bool {
	return false
}

// Read reads a single object from a collection into an interface instance
func (fs *FlatStorage) Read(collection string, objectID string, out interface{}) error {
	return fmt.Errorf("Not implemented")
}

// ReadAll reads all objects from a collection into an interface array of objectType
func (fs *FlatStorage) ReadAll(collection string, objectID string, objectType interface{}) ([]interface{}, error) {
	return nil, fmt.Errorf("Not implemented")
}

// Write writes a single object into a collection from an interface instance
func (fs *FlatStorage) Write(collection string, objectID string, object interface{}) error {
	return fmt.Errorf("Not implemented")
}
