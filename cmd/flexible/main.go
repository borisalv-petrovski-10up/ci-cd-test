package main

import (
	"net/http"

	"github.com/borisalv-petrovski-10up/ci-cd-test/services/app-ae-flexible/handlers"
)

func main() {
	http.HandleFunc("/", handlers.UniqueNames)
	http.ListenAndServe(":8080", nil)
}
