package bootstrap

import (
	"embed"
	"test/src/hybrid"
	"test/src/hybrid/values"
	"test/src/tools/env"
)

type Configuration struct {
	UI_EMBED     *embed.FS
	ASSETS_EMBED *embed.FS
	DEBUG        bool
}

type Application struct {
	Name          string
	Window        *hybrid.Window
	Configuration *Configuration
}

type ApplicationCallback func(application Application)

var Instance Application

type ApplicationConfiguration interface {
	Boot(application Application)
}

// Comment
func (ctx *Configuration) Bootstrap(config ApplicationConfiguration) {
	env.LoadDefault()

	configuration := hybrid.Configuration{
		Title:     env.Get("WINDOW_TITLE"),
		Host:      env.Get("UI_ADDRESS"),
		EmbedPath: env.Get("UI_PATH"),
		Embed:     ctx.UI_EMBED,
		Debug:     ctx.DEBUG,
	}

	window := hybrid.Initialization(configuration, func(window *hybrid.Window) {
		Instance = Application{
			Name:   env.Get("APP_NAME"),
			Window: window,
		}

		config.Boot(Instance)
	})

	window.Open()
	window.Destroy()
}

// Comment
func (ctx *Application) Bind(name string, object values.BindingArg) {
	ctx.Window.Binding(name, object)
}
