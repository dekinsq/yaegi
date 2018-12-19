package stdlib

// Code generated by 'goexports hash/crc64'. DO NOT EDIT.

import (
	"hash/crc64"
	"reflect"
)

func init() {
	Value["hash/crc64"] = map[string]reflect.Value{
		"Checksum":  reflect.ValueOf(crc64.Checksum),
		"ECMA":      reflect.ValueOf(uint(crc64.ECMA)),
		"ISO":       reflect.ValueOf(uint(crc64.ISO)),
		"MakeTable": reflect.ValueOf(crc64.MakeTable),
		"New":       reflect.ValueOf(crc64.New),
		"Size":      reflect.ValueOf(crc64.Size),
		"Update":    reflect.ValueOf(crc64.Update),
	}
	Type["hash/crc64"] = map[string]reflect.Type{
		"Table": reflect.TypeOf((*crc64.Table)(nil)).Elem(),
	}
}
