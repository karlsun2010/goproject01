package main

import "net/http"

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

type aboutHandler struct{}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About!"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome!"))
}

func main() {
	a := helloHandler{}
	b := aboutHandler{}

	server := http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: nil,
	}

	//使用http.Handle
	http.Handle("/hello", &a)
	http.Handle("/about", &b)
	http.Handle("/welcome2", http.HandlerFunc(welcome))
	//使用http.HandleFunc
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home!"))
	})
	http.HandleFunc("/welcome", welcome)
	server.ListenAndServe()
}
