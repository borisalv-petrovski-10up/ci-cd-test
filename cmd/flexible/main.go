package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/borisalv-petrovski-10up/ci-cd-test/services/app-ae-flexible/handlers"
)

func main() {
	tpl := template.Must(template.ParseGlob("./static/*.html"))
	homepageHandler := handlers.NewHomepage(tpl)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", homepageHandler.HomepageHandler)
	http.HandleFunc("/names", handlers.UniqueNames)

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
