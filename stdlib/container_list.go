package stdlib

// Code generated by 'goexports container/list'. DO NOT EDIT.

import (
	"container/list"
	"reflect"
)

func init() {
	Value["container/list"] = map[string]reflect.Value{
		"New": reflect.ValueOf(list.New),
	}
	Type["container/list"] = map[string]reflect.Type{
		"Element": reflect.TypeOf((*list.Element)(nil)).Elem(),
		"List":    reflect.TypeOf((*list.List)(nil)).Elem(),
	}
}
