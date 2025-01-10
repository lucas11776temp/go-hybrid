package bootstrap

import (
	"embed"
	"test/src/hybrid"
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

func (ctx *Configuration) Bootstrap(setup ApplicationCallback) {
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

		setup(Instance)
	})

	window.Open()
	window.Destroy()
}

func (ctx *Application) Bind(name string, object hybrid.BindingArg) {
	ctx.Window.Binding(name, object)
}
