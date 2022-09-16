package main

import (
	"log"
	"net/http"
	"strings"
)

const directory = "./public"

func main() {
	fs := http.FileServer(http.Dir(directory))
	log.Print("Serving " + directory + " on http://localhost:3000")
	http.ListenAndServe(":3000", http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Add("Cache-Control", "no-cache")
		if strings.HasSuffix(req.URL.Path, ".wasm") {
			resp.Header().Set("content-type", "application/wasm")
		}
		fs.ServeHTTP(resp, req)
	}))
}
