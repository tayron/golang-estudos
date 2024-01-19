package main

import "github.com/webview/webview"

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Hospeda App")
	//w.SetSize(1024, 768, webview.HintNone)
	w.Navigate("https://hospeda.app")
	w.Run()
}
