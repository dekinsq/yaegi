package stdlib

// Code generated by 'goexports mime'. DO NOT EDIT.

import (
	"mime"
	"reflect"
)

func init() {
	Value["mime"] = map[string]reflect.Value{
		"AddExtensionType":         reflect.ValueOf(mime.AddExtensionType),
		"BEncoding":                reflect.ValueOf(mime.BEncoding),
		"ErrInvalidMediaParameter": reflect.ValueOf(mime.ErrInvalidMediaParameter),
		"ExtensionsByType":         reflect.ValueOf(mime.ExtensionsByType),
		"FormatMediaType":          reflect.ValueOf(mime.FormatMediaType),
		"ParseMediaType":           reflect.ValueOf(mime.ParseMediaType),
		"QEncoding":                reflect.ValueOf(mime.QEncoding),
		"TypeByExtension":          reflect.ValueOf(mime.TypeByExtension),
	}
	Type["mime"] = map[string]reflect.Type{
		"WordDecoder": reflect.TypeOf((*mime.WordDecoder)(nil)).Elem(),
		"WordEncoder": reflect.TypeOf((*mime.WordEncoder)(nil)).Elem(),
	}
}
