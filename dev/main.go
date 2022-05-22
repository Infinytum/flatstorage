package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(reflectTypeKey(reflect.TypeOf(http.Request{})))
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
