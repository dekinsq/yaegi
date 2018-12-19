package stdlib

// Code generated by 'goexports crypto/dsa'. DO NOT EDIT.

import (
	"crypto/dsa"
	"reflect"
)

func init() {
	Value["crypto/dsa"] = map[string]reflect.Value{
		"ErrInvalidPublicKey": reflect.ValueOf(dsa.ErrInvalidPublicKey),
		"GenerateKey":         reflect.ValueOf(dsa.GenerateKey),
		"GenerateParameters":  reflect.ValueOf(dsa.GenerateParameters),
		"L1024N160":           reflect.ValueOf(dsa.L1024N160),
		"L2048N224":           reflect.ValueOf(dsa.L2048N224),
		"L2048N256":           reflect.ValueOf(dsa.L2048N256),
		"L3072N256":           reflect.ValueOf(dsa.L3072N256),
		"Sign":                reflect.ValueOf(dsa.Sign),
		"Verify":              reflect.ValueOf(dsa.Verify),
	}
	Type["crypto/dsa"] = map[string]reflect.Type{
		"ParameterSizes": reflect.TypeOf((*dsa.ParameterSizes)(nil)).Elem(),
		"Parameters":     reflect.TypeOf((*dsa.Parameters)(nil)).Elem(),
		"PrivateKey":     reflect.TypeOf((*dsa.PrivateKey)(nil)).Elem(),
		"PublicKey":      reflect.TypeOf((*dsa.PublicKey)(nil)).Elem(),
	}
}
