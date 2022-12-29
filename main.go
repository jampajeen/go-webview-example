package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/webview/webview"
)

var count = 0

func main() {
	debug := true
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("Test Golang webview")
	w.SetSize(800, 600, webview.HintNone)

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	html, err := ioutil.ReadFile(path + "/index.html")
	if err != nil {
		panic("index.html not found")
	}
	w.SetHtml(string(html))
	w.Bind("onCount", func(param string) {
		go showCount(w, param)
	})

	w.Bind("onSum", func(a, b int) int {
		return a + b
	})

	w.Run()
}

func showCount(p webview.WebView, param string) {
	count++
	text := fmt.Sprintf("%s = %d", param, count)

	p.Eval("document.getElementById('countLabel').innerHTML = " + text)

	p.Eval("notify('" + text + "')")
}
