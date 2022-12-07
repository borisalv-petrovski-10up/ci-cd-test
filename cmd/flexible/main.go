package main

import (
	"net/http"
	"time"

	"github.com/borisalv-petrovski-10up/ci-cd-test/services/app-ae-flexible/handlers"
)

func main() {
	http.HandleFunc("/a", handlers.UniqueNames)

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
