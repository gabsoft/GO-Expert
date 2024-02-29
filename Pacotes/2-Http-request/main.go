package main

import (
	"io"
	"net/http"
)

func main() {
	// Chamada http

	req, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
	req.Body.Close()

}
