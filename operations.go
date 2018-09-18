package flatstorage

import (
	"fmt"
)

// Read reads a single object from a collection into an interface instance
func (fs *FlatStorage) Read(collection string, objectID string, out interface{}) error {
	return fmt.Errorf("Not implemented")
}

// ReadAll reads all objects from a collection into an interface array of objectType
func (fs *FlatStorage) ReadAll(collection string, objectID string, objectType interface{}) ([]interface{}, error) {
	return nil, fmt.Errorf("Not implemented")
}
