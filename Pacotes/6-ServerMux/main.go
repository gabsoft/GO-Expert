package servermux

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/oi", HomeHandler)
	mux.Handle("/blog", blog{title: "My Blog"})

	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá, mundo!"))
}

type blog struct {
	title string
}

func (b blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(b.title))
}
