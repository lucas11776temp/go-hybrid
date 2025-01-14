package values

import (
	"reflect"
	"test/src/hybrid/values/cast"
)

type BindingArg interface{}

type ObjectMethods map[string]reflect.Method

// Comment
func GetValue(value reflect.Value) any {
	switch value.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return value.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return value.Uint()
	case reflect.Float32, reflect.Float64:
		return value.Float()
	case reflect.Bool:
		return value.Bool()
	case reflect.String:
		return value.String()
	default:
		return value
	}
}

// Comment
func CastValue(t reflect.Type, value interface{}) reflect.Value {
	return reflect.ValueOf(cast.Cast(t, value))
}

// Comment
func Arguments(method reflect.Method, data []any) []reflect.Value {
	arguments := []reflect.Value{}
	typeOf := method.Func.Type()
	dLen := len(data)

	for i := 1; i < typeOf.NumIn(); i++ {
		if i > dLen {
			arguments = append(arguments, CastValue(typeOf.In(i), nil))
			continue
		}

		arguments = append(arguments, CastValue(typeOf.In(i), data[i-1]))
	}

	return arguments
}

// Comment
func Methods(object BindingArg) ObjectMethods {
	methods := make(ObjectMethods)
	typeOf := reflect.TypeOf(object)

	for i := 0; i < typeOf.NumMethod(); i++ {
		if typeOf.Method(i).Type.Kind() == reflect.Func {
			methods[typeOf.Method(i).Name] = typeOf.Method(i)
		}
	}

	return methods
}
