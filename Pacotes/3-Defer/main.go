package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// defer

	req, err := http.Get("http://google.com")
	if err != nil {
		panic(err)
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}
	println(string(res))
	req.Body.Close()

	defer fmt.Println("Primeira linha")
	defer fmt.Println("Segunda linha")
	fmt.Println("Terceira linha")

}
