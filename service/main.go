package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"time"
)

var pwd = os.Getenv("PWD")

type DefaultHandler struct{}

func writeFileTo(writer http.ResponseWriter, filename string) {
	content, err := os.ReadFile(filename)

	if err == nil {
		writer.Write(content)
	} else {
		writer.WriteHeader(500)
	}
}

func (*DefaultHandler) ServeHTTP(writer http.ResponseWriter, req *http.Request) {
	urlObj := req.URL

	switch urlObj.Path {
	case "/js/iframe-sleep-2s.js":
		time.Sleep(2 * time.Second)
		writer.Header().Add("Content-Type", "text/javascript")
		writeFileTo(writer, path.Join(pwd, "/js/iframe-sleep-2s.js"))
	case "/js/iframe-sleep-3s.js":
		time.Sleep(3 * time.Second)
		writer.Header().Add("Content-Type", "text/javascript")
		writeFileTo(writer, path.Join(pwd, "/js/iframe-sleep-3s.js"))
	case "/js/sleep-2s.js":
		time.Sleep(2 * time.Second)
		writer.Header().Add("Content-Type", "text/javascript")
		writeFileTo(writer, path.Join(pwd, "/js/sleep-2s.js"))
	case "/js/sleep-3s.js":
		time.Sleep(3 * time.Second)
		writer.Header().Add("Content-Type", "text/javascript")
		writeFileTo(writer, path.Join(pwd, "/js/sleep-3s.js"))
	case "/iframe.html":
		writer.Header().Add("Content-Type", "text/html")
		writeFileTo(writer, path.Join(pwd, "/iframe.html"))
	case "/index.html":
		writer.Header().Add("Content-Type", "text/html")
		writeFileTo(writer, path.Join(pwd, "/index.html"))
	default:
		writer.WriteHeader(500)
	}

}

func main() {
	tip := "http server is running.\n" +
		"visit http://localhost:8077/index.html."

	fmt.Println(tip)

	handler := &DefaultHandler{}
	http.ListenAndServe(":8077", handler)
}
