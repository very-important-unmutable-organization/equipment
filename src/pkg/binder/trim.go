package binder

import (
	"errors"
	"reflect"
	"strings"
)

const (
	trimTag  = "trim"
	skipTrim = "false"
)

func TrimStrFields(i interface{}) error {
	valPtr := reflect.ValueOf(i)
	if valPtr.Kind() != reflect.Ptr || valPtr.IsNil() {
		return errors.New("type of object is not pointer")
	}

	val := valPtr.Elem()
	typ := val.Type()
	for j := 0; j < val.NumField(); j++ {
		fieldVal := val.Field(j)

		if fieldVal.Kind() == reflect.Ptr && !fieldVal.IsNil() {
			fieldVal = fieldVal.Elem()
		}

		if fieldVal.Kind() != reflect.String || !fieldVal.CanSet() {
			continue
		}

		tag := typ.Field(j).Tag.Get(trimTag)
		if tag == skipTrim {
			continue
		}

		str := fieldVal.Interface().(string)
		str = strings.TrimSpace(str)
		fieldVal.SetString(str)
	}
	return nil
}
