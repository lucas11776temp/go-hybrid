package hybrid

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"test/src/hybrid/values"

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

type BindingArguments struct {
	Object string `json:"object"`
	Method string `json:"method"`
	Data   []any  `json:"data"`
}

type ObjectBinding struct {
	Object   values.BindingArg
	Bindings values.ObjectMethods
}

type Bindings map[string]ObjectBinding

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

// Comment
func (ctx *HttpServerHandler) OpenFile(name string) ([]byte, error) {
	file, err := ctx.Embed.ReadFile(*ctx.EmbedPath + "/" + name)

	if err != nil {
		return nil, err
	}

	return file, nil
}

// Comment
func isHtml(filename string) bool {
	split := strings.Split(filename, ".")

	return strings.ToLower(split[len(split)-1]) == "html"
}

// Comment
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

// Comment
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

// Comment
func (ctx *Server) Address() string {
	return strings.Join([]string{ctx.IP_Address, "9999"}, ":")
}

// Comment
func (ctx *Server) Host() string {
	return "http://" + ctx.Address()
}

// Comment
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

// Comment
func (ctx *Window) setUpBindingEventHandler() {
	ctx.Webview.Bind("__BINDING__", func(payload string) any {
		var argument BindingArguments

		err := json.Unmarshal([]byte(payload), &argument)

		if err != nil {
			fmt.Println("Invalid binding arguments:", payload)
			return nil
		}

		objMap, ok := ctx.bindings[argument.Object]

		if !ok {
			fmt.Println("Binding to object", argument.Object, "does not exist")
			return nil
		}

		method := objMap.Bindings[argument.Method]

		args := []reflect.Value{reflect.ValueOf(objMap.Object)}

		args = append(args, values.Arguments(method, argument.Data)...)

		returnValues := method.Func.Call(args)

		vs := []any{}

		for i := 0; i < len(returnValues); i++ {
			vs = append(vs, values.GetValue(returnValues[i]))
		}

		if len(vs) > 1 {
			return vs
		}

		return vs[0]
	})
}

// Comment
func (ctx *Window) Binding(name string, obj values.BindingArg) {
	ctx.bindings[name] = ObjectBinding{
		Object:   obj,
		Bindings: values.Methods(obj),
	}
}

// Comment
func (ctx *Window) Open() {
	ctx.Webview.Navigate(ctx.Server.Host())
	ctx.Webview.Run()
}

// Comment
func (ctx *Window) Destroy() {
	ctx.Webview.Destroy()
}
