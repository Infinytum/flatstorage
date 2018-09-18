package flatstorage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

// Delete removes a single resource from a collection
func (fs *FlatStorage) Delete(collection string, resource string) error {
	if fs.Exists(collection, resource) {
		defer fs.UnlockCollection(collection)
		fs.LockCollection(collection)
		logrus.WithField("collection", collection).WithField("resource", resource).Debug("Removing resource from collection")
		return os.Remove(fs.resourcePath(collection, resource))
	}
	return nil
}

// DeleteAll removes an entire collection from the filesystem
func (fs *FlatStorage) DeleteAll(collection string) error {
	if fs.CollectionExists(collection) {
		defer fs.UnlockCollection(collection)
		fs.LockCollection(collection)
		logrus.WithField("collection", collection).Debug("Deleting entire collection")
		return os.RemoveAll(fs.collectionPath(collection))
	}
	return nil
}

// Exists checks if a resource is present in a collection
func (fs *FlatStorage) Exists(collection string, resource string) bool {
	return pathExists(fs.resourcePath(collection, resource))
}

// CollectionExists checks if a collection exists
func (fs *FlatStorage) CollectionExists(collection string) bool {
	return pathExists(fs.collectionPath(collection))
}

// Read reads a single resource from a collection into an interface instance
func (fs *FlatStorage) Read(collection string, resource string, out interface{}) error {
	if !fs.CollectionExists(collection) {
		logrus.WithField("collection", collection).Debug("Tried to read resource from non-existent collection")
		return collectionNotExistent(collection)
	}

	if !fs.Exists(collection, resource) {
		logrus.WithField("collection", collection).WithField("resource", resource).Debug("Tried to read resource from non-existent collection")
		return resourceNotExistent(collection, resource)
	}

	bytes, err := ioutil.ReadFile(fs.resourcePath(collection, resource))
	if err != nil {
		logrus.WithField("collection", collection).WithField("resource", resource).Error("Error while trying to read resource", err)
		return err
	}

	err = json.Unmarshal(bytes, &out)
	if err != nil {
		logrus.WithField("collection", collection).WithField("resource", resource).Error("Error while trying to unmarshal resource", err)
	}
	return err
}

// ReadAll reads all resources from a collection into an interface array of resourceType
func (fs *FlatStorage) ReadAll(collection string, resourceType interface{}) ([]interface{}, error) {
	resourceList := make([]interface{}, 0)
	if !fs.CollectionExists(collection) {
		return resourceList, collectionNotExistent(collection)
	}

	resources, err := ioutil.ReadDir(fs.collectionPath(collection))
	if err != nil {
		logrus.WithField("collection", collection).Error("Error while trying to list resources", err)
		return resourceList, err
	}

	for _, resourceFile := range resources {
		var clone = reflect.New(reflect.ValueOf(resourceType).Elem().Type()).Interface()
		name := strings.TrimSuffix(resourceFile.Name(), filepath.Ext(resourceFile.Name()))
		err := fs.Read(collection, name, &clone)
		if err != nil {
			return make([]interface{}, 0), err
		}
		resourceList = append(resourceList, clone)
	}

	return resourceList, nil
}

// Write writes a single object into a collection from an interface instance
func (fs *FlatStorage) Write(collection string, resource string, object interface{}) error {

	if collection == "" {
		logrus.Warn("Tried to save resource without collection identifier")
		return fmt.Errorf("No collection specified for writing resource")
	}

	if resource == "" {
		logrus.Warn("Tried to save resource without resource identifier")
		return fmt.Errorf("No resource specified for writing resource")
	}

	defer fs.UnlockCollection(collection)
	fs.LockCollection(collection)

	resPath := fs.resourcePath(collection, resource)
	resTempPath := resPath + ".fstemp"

	if !fs.CollectionExists(collection) {
		err := os.MkdirAll(fs.collectionPath(collection), 0755)
		if err != nil {
			logrus.WithField("collection", collection).Error("Could not create collection directory", err)
			return err
		}
	}

	// Outputs pretty, tab-indented json files
	bytes, err := json.MarshalIndent(object, "", "\t")
	if err != nil {
		logrus.WithField("collection", collection).WithField("resource", resource).Error("Error while marshaling resource", err)
		return err
	}

	if err := ioutil.WriteFile(resTempPath, bytes, 0644); err != nil {
		logrus.WithField("collection", collection).WithField("resource", resource).Error("Error while writing resource to disk", err)
		return err
	}

	// Ensure we dont write in the middle of a read access to the file
	return os.Rename(resTempPath, resPath)
}
