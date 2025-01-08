package log

import (
	webview "github.com/webview/webview_go"
)

func Fatal(err error) {
	window := webview.New(false)

	window.SetSize(400, 100, webview.HintFixed)
	window.SetTitle("Application Has Crushed")
	window.SetHtml(err.Error())
	window.Run()
	window.Destroy()
}
