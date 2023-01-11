package utils

import (
	"reflect"
)

/* To get the Name of the struct */
func GetStructName(strct interface{}) string {
	return reflect.TypeOf(strct).Name()
}
