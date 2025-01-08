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
	Data any `json:"data"`
}

// type Bindings map[string]map[string]func(...[]any)

type Bindings map[string]map[string]reflect.Value

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

func (ctx *Window) setUpBindingEventHandler() {
	ctx.Webview.Bind("__BINDING__", func(event string) {
		// var evt BindingArguments[any]
		var evt BindingArguments

		err := json.Unmarshal([]byte(event), &evt)

		if err != nil {
			fmt.Println("Invalid binding arguments:", event)
			return
		}

		object, ok := ctx.bindings[strings.ToUpper(evt.Object)]

		if !ok {
			fmt.Println("Binding to object", evt.Object, "does not exist")
			return
		}

		fmt.Println(object)
		// fmt.Println(object[evt.Method].Call(evt.Data))

		// fmt.Println(object[evt.Method].Call([]reflect.Value{}))
	})
}

func (ctx *Window) Binding(name string, obj interface{}) {
	typeOf := reflect.TypeOf(&obj)

	fmt.Println(typeOf)

	// ctx.bindings[name] = obj

	ctx.bindings[name] = make(map[string]reflect.Value)

	fmt.Println(obj, typeOf.NumMethod(), reflect.ValueOf(&obj).MethodByName("Add"))

	for i := 0; i < typeOf.NumMethod(); i++ {
		method := typeOf.Method(i)

		fmt.Println(
			typeOf.Name(), ":", reflect.ValueOf(&obj).MethodByName(method.Name),
		)

		ctx.bindings[name][method.Name] = reflect.ValueOf(&obj).MethodByName(method.Name)
	}
}

func (ctx *Window) Open() {
	ctx.Webview.Navigate(ctx.Server.Host())
	ctx.Webview.Run()
}

func (ctx *Window) Destroy() {
	ctx.Webview.Destroy()
}
