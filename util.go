package flatstorage

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func reflectTypeKey(t reflect.Type) string {
	nameType := t
	for nameType != nil && nameType.Kind() == reflect.Pointer {
		nameType = nameType.Elem()
	}

	pkg, name := "UNKNOWN_PACKAGE", t.String()
	if nameType != nil {
		pkg = nameType.PkgPath()
	}
	pkg = strings.ReplaceAll(pkg, "/", "-")
	return fmt.Sprintf("%s_%s", pkg, name)
}
