package stdlib

// Code generated by 'goexports compress/flate'. DO NOT EDIT.

import (
	"compress/flate"
	"reflect"
)

func init() {
	Value["compress/flate"] = map[string]reflect.Value{
		"BestCompression":    reflect.ValueOf(flate.BestCompression),
		"BestSpeed":          reflect.ValueOf(flate.BestSpeed),
		"DefaultCompression": reflect.ValueOf(flate.DefaultCompression),
		"HuffmanOnly":        reflect.ValueOf(flate.HuffmanOnly),
		"NewReader":          reflect.ValueOf(flate.NewReader),
		"NewReaderDict":      reflect.ValueOf(flate.NewReaderDict),
		"NewWriter":          reflect.ValueOf(flate.NewWriter),
		"NewWriterDict":      reflect.ValueOf(flate.NewWriterDict),
		"NoCompression":      reflect.ValueOf(flate.NoCompression),
	}
	Type["compress/flate"] = map[string]reflect.Type{
		"CorruptInputError": reflect.TypeOf((*flate.CorruptInputError)(nil)).Elem(),
		"InternalError":     reflect.TypeOf((*flate.InternalError)(nil)).Elem(),
		"ReadError":         reflect.TypeOf((*flate.ReadError)(nil)).Elem(),
		"Reader":            reflect.TypeOf((*flate.Reader)(nil)).Elem(),
		"Resetter":          reflect.TypeOf((*flate.Resetter)(nil)).Elem(),
		"WriteError":        reflect.TypeOf((*flate.WriteError)(nil)).Elem(),
		"Writer":            reflect.TypeOf((*flate.Writer)(nil)).Elem(),
	}
}
