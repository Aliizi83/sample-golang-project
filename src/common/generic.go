package common

import "reflect"

func GetTypeName[T any]() string {
	t := reflect.TypeOf((*T)(nil)).Elem()

	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	return t.Name()
}
