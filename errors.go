package flatstorage

import (
	"fmt"
)

func resourceNotExistent(collection string, resource string) error {
	return fmt.Errorf("The requested resource %s does not exists in the collection %s", resource, collection)
}

func collectionNotExistent(collection string) error {
	return fmt.Errorf("The requested collection %s does not exist", collection)
}
