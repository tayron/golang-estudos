package main

import "github.com/yzhanginwa/webview"

func main() {
	webView := webview.New(webview.Settings{
		Title:                  "My test web view app",
		URL:                    "http://google.com",
		Width:                  1000,
		Height:                 800,
		Resizable:              true,
		Debug:                  true,
		ExternalInvokeCallback: nil,
	})

	webView.SetFullscreen(true)
	webView.Run()
}
