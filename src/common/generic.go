package common

import "reflect"

func GetTypeName[T any]() string {
	t := reflect.TypeFor[T]()

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Name()
}
