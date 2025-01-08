package main

import (
	"embed"
	"math"
	"reflect"

	bootstrap "test/src"
)

//go:embed assets/*
var assets_embed embed.FS

//go:embed build/*
var ui_embed embed.FS

type Math struct {
}

func (ctx *Math) Addition(a int32, b int32) int32 {
	return a + b
}

func (ctx *Math) Multiple(a int32, b int32) int32 {
	return a * b
}

func (ctx *Math) Sin(x float64) float64 {
	return math.Sin(x)
}

func main() {
	configuration := bootstrap.Configuration{
		UI_EMBED:     &ui_embed,
		ASSETS_EMBED: &assets_embed,
		DEBUG:        true,
	}

	configuration.Bootstrap(func(application bootstrap.Application) {
		application.Bind("Math2", Math{})
	})
}

type StructMethods map[string]reflect.Value

// func GetStructMethods(s interface{}) StructMethods {
// 	methods := make(StructMethods)

// 	typeOf := reflect.TypeOf(s)

// 	for i := 0; i < typeOf.NumMethod(); i++ {
// 		method := typeOf.Method(i)

// 		fmt.Println(
// 			method.Name, reflect.ValueOf(s).Method(i),
// 		)

// 		methods[strings.ToUpper(method.Name)] = method.Func

// 		// fmt.Println(method.Func.Call([]reflect.Value{}))

// 		// ctx.bindings[name][method.Name] = reflect.ValueOf(&m).MethodByName(method.Name)
// 	}

// 	return methods
// }

// application.Window.Webview.Init(`window.API_ADDRESS="127.0.0.1:9999"`)
