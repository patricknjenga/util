package validator

import (
	"reflect"
)

func New(tag string, structs ...any) map[string]map[string]any {
	r := map[string]map[string]any{}
	for _, v := range structs {
		var (
			f = map[string]any{}
			t = reflect.TypeOf(v)
		)
		for i := 0; i < t.NumField(); i++ {
			f[t.Field(i).Name] = t.Field(i).Tag.Get(tag)
		}
		r[t.Name()] = f
	}
	return r
}
