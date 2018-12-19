package stdlib

// Code generated by 'goexports expvar'. DO NOT EDIT.

import (
	"expvar"
	"reflect"
)

func init() {
	Value["expvar"] = map[string]reflect.Value{
		"Do":        reflect.ValueOf(expvar.Do),
		"Get":       reflect.ValueOf(expvar.Get),
		"Handler":   reflect.ValueOf(expvar.Handler),
		"NewFloat":  reflect.ValueOf(expvar.NewFloat),
		"NewInt":    reflect.ValueOf(expvar.NewInt),
		"NewMap":    reflect.ValueOf(expvar.NewMap),
		"NewString": reflect.ValueOf(expvar.NewString),
		"Publish":   reflect.ValueOf(expvar.Publish),
	}
	Type["expvar"] = map[string]reflect.Type{
		"Float":    reflect.TypeOf((*expvar.Float)(nil)).Elem(),
		"Func":     reflect.TypeOf((*expvar.Func)(nil)).Elem(),
		"Int":      reflect.TypeOf((*expvar.Int)(nil)).Elem(),
		"KeyValue": reflect.TypeOf((*expvar.KeyValue)(nil)).Elem(),
		"Map":      reflect.TypeOf((*expvar.Map)(nil)).Elem(),
		"String":   reflect.TypeOf((*expvar.String)(nil)).Elem(),
		"Var":      reflect.TypeOf((*expvar.Var)(nil)).Elem(),
	}
}
