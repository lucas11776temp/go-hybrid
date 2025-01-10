package hybrid

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"

	webview "github.com/webview/webview_go"
)

//go:embed scripts/*
var scripts embed.FS

type Server struct {
	IP_Address string
	Port       int
	EntryFile  string
	Server     *http.Server
	EmbedPath  string
	Embed      *embed.FS
}

type HttpServerHandler struct {
	EmbedPath *string
	Embed     *embed.FS
}

//	type BindingArguments[T any] struct {
//		Object string `json:"object"`
//		Method string `json:"method"`
//		Data   []any      `json:"data"`
//	}
type BindingArguments struct {
	Object string `json:"object"`
	Method string `json:"method"`
	// Data   []reflect.Value `json:"data"`
	Data []any `json:"data"`
}

// type Bindings map[string]map[string]func(...[]any)

// type Bindings map[string]map[string]reflect.Value

type BindingArg interface{}

type ObjectBinding struct {
	Object   BindingArg
	Bindings map[string]reflect.Method
}

type Bindings map[string]ObjectBinding

// type Bindings map[string]interface{}

type Window struct {
	Server   Server
	Webview  webview.WebView
	bindings Bindings
}

type InitializationCallback func(window *Window)

type Configuration struct {
	Title     string
	Host      string
	EmbedPath string
	Embed     *embed.FS
	Debug     bool
}

func (ctx *HttpServerHandler) OpenFile(name string) ([]byte, error) {
	file, err := ctx.Embed.ReadFile(*ctx.EmbedPath + "/" + name)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func isHtml(filename string) bool {
	split := strings.Split(filename, ".")

	return strings.ToLower(split[len(split)-1]) == "html"
}

func (h HttpServerHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fileRequest := strings.Trim(req.URL.Path, "/")

	if fileRequest == "" {
		fileRequest = "index.html"
	}

	fileRequest = strings.Split(fileRequest, "?")[0]

	file, err := h.OpenFile(fileRequest)

	if err != nil {
		file = []byte("View Not Found")
	}

	if isHtml(fileRequest) {
		res.Header().Add("content-type", "text/html")
	} else {
		res.Header().Add("content-type", req.Header.Get("accept"))
	}

	res.Write(file)
}

func (ctx *Server) run() {
	ctx.Server = &http.Server{
		Addr: ctx.Address(),
		Handler: HttpServerHandler{
			EmbedPath: &ctx.EmbedPath,
			Embed:     ctx.Embed,
		},
	}

	go func() {
		err := ctx.Server.ListenAndServe()

		if err != nil {
			log.Fatal(err)
		}
	}()
}

func (ctx *Server) Address() string {
	return strings.Join([]string{ctx.IP_Address, "9999"}, ":")
}

func (ctx *Server) Host() string {
	return "http://" + ctx.Address()
}

func Initialization(init Configuration, setup InitializationCallback) Window {
	window := Window{
		Server: Server{
			IP_Address: init.Host,
			Port:       9999,
			EmbedPath:  init.EmbedPath,
			Embed:      init.Embed,
		},
		Webview:  webview.New(init.Debug),
		bindings: make(Bindings),
	}

	window.Server.run()

	setup(&window)

	window.Webview.SetTitle(init.Title)

	window.setUpBindingEventHandler()

	return window
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
	default:
		return value.String()
	}
}

func (ctx *Window) setUpBindingEventHandler() {
	ctx.Webview.Bind("__BINDING__", func(event string) any {
		// var evt BindingArguments[any]
		var evt BindingArguments

		err := json.Unmarshal([]byte(event), &evt)

		if err != nil {
			fmt.Println("Invalid binding arguments:", event)
			return nil
		}

		objMap, ok := ctx.bindings[evt.Object]

		if !ok {
			fmt.Println("Binding to object", evt.Object, "does not exist")
			return nil
		}

		m := objMap.Bindings[evt.Method]

		re := m.Func.Call([]reflect.Value{
			reflect.ValueOf(objMap.Object),
		})

		vs := []any{}

		for i := 0; i < len(re); i++ {
			vs = append(vs, GetValue(re[i]))
		}

		fmt.Println("Return Values", vs)

		if len(vs) > 1 {
			return vs
		}

		return vs[0]
	})
}

func (ctx *Window) Binding(name string, obj BindingArg) {
	ctx.bindings[name] = ObjectBinding{
		Object:   obj,
		Bindings: make(map[string]reflect.Method),
	}

	typeOf := reflect.TypeOf(obj)

	for i := 0; i < typeOf.NumMethod(); i++ {
		if typeOf.Method(i).Type.Kind() == reflect.Func {
			ctx.bindings[name].Bindings[typeOf.Method(i).Name] = typeOf.Method(i)
		}
	}
}

func (ctx *Window) Open() {
	ctx.Webview.Navigate(ctx.Server.Host())
	ctx.Webview.Run()
}

func (ctx *Window) Destroy() {
	ctx.Webview.Destroy()
}
