package stdlib

// Code generated by 'goexports hash/crc32'. DO NOT EDIT.

import (
	"hash/crc32"
	"reflect"
)

func init() {
	Value["hash/crc32"] = map[string]reflect.Value{
		"Castagnoli":   reflect.ValueOf(crc32.Castagnoli),
		"Checksum":     reflect.ValueOf(crc32.Checksum),
		"ChecksumIEEE": reflect.ValueOf(crc32.ChecksumIEEE),
		"IEEE":         reflect.ValueOf(crc32.IEEE),
		"IEEETable":    reflect.ValueOf(crc32.IEEETable),
		"Koopman":      reflect.ValueOf(crc32.Koopman),
		"MakeTable":    reflect.ValueOf(crc32.MakeTable),
		"New":          reflect.ValueOf(crc32.New),
		"NewIEEE":      reflect.ValueOf(crc32.NewIEEE),
		"Size":         reflect.ValueOf(crc32.Size),
		"Update":       reflect.ValueOf(crc32.Update),
	}
	Type["hash/crc32"] = map[string]reflect.Type{
		"Table": reflect.TypeOf((*crc32.Table)(nil)).Elem(),
	}
}
