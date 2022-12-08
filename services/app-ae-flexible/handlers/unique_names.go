package handlers

import (
	"fmt"
	"net/http"

	"github.com/samber/lo"
)

func UniqueNames(w http.ResponseWriter, r *http.Request) {
	names := lo.Uniq[string]([]string{"Samuel", "John", "Samuel", "Falcoooooooooooooo"})
	_, err := fmt.Fprintf(w, "Hello, %v!", names)
	if err != nil {
		panic(err)
	}
}
