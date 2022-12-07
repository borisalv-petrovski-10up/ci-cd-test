package main

import (
	"net/http"

	"github.com/borisalv-petrovski-10up/ci-cd-test/services/app-ae-flexible/handlers"
)

func main() {
	http.HandleFunc("/a", handlers.UniqueNames)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
