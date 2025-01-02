package main

import (
	"fmt"
	"log"
	"net/http"

	"gg"
)

func main() {
	r := gg.New()

	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	log.Printf("Server is running on http://localhost:9999")
	r.Run(":9999")
}
