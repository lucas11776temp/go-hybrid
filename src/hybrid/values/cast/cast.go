package cast

import (
	"reflect"

	"github.com/spf13/cast"
)

// Comment
func Cast(t reflect.Type, value interface{}) any {
	switch t.Kind() {

	case reflect.Bool:
		return cast.ToBool(value)

	case reflect.Int:
		return cast.ToInt(value)

	case reflect.Int8:
		return cast.ToInt8(value)

	case reflect.Int16:
		return cast.ToInt16(value)

	case reflect.Int32:
		return cast.ToInt32(value)

	case reflect.Int64:
		return cast.ToInt64(value)

	case reflect.Uint:
		return cast.ToUint(value)

	case reflect.Uint8:
		return cast.ToUint8(value)

	case reflect.Uint16:
		return cast.ToUint16(value)

	case reflect.Uint32:
		return cast.ToUint32(value)

	case reflect.Uint64:
		return cast.ToUint64(value)

	// case reflect.Uintptr:
	// 	return cast.ToUintE(value)

	case reflect.Float32:
		return cast.ToFloat32(value)

	case reflect.Float64:
		return cast.ToFloat64(value)

	// case reflect.Complex64:
	// 	return cast.ToUint(value)

	// case reflect.Complex128:
	// 	return cast.ToUint(value)

	case reflect.String:
		return cast.ToString(value)

	default:
		return value

	}
}

// Invalid Kind = iota
// Bool
// Int
// Int8
// Int16
// Int32
// Int64
// Uint
// Uint8
// Uint16
// Uint32
// Uint64
// Uintptr
// Float32
// Float64
// Complex64
// Complex128
// Array
// Chan
// Func
// Interface
// Map
// Pointer
// Slice
// String
// Struct
// UnsafePointer
