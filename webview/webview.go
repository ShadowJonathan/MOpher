package main

import (
	"github.com/zserge/webview"
	"log"
	"net/http"
	"strings"
)

var wv webview.WebView

func L(text string) {
	wv.Eval("add(\"" + strings.TrimSpace(text) + "\")")
}

func main() {
	callback := func(w webview.WebView, data string) {
		L("> " + data)
	}

	addr := "127.0.0.1:9999"
	go func() {
		m := http.NewServeMux()
		m.HandleFunc("/", http.FileServer(http.Dir(".")).ServeHTTP)
		log.Fatal(http.ListenAndServe(addr, m))
	}()

	wv = webview.New(webview.Settings{Title: "CMD", URL: "http://" + addr, Width: 375, Height: 310, Resizable: true, ExternalInvokeCallback: callback})
	wv.Eval("location.refresh()")
	wv.Run()
}
