package handlers

import (
	"fmt"
	"net/http"

	"github.com/samber/lo"
)

func UniqueNames(w http.ResponseWriter, r *http.Request) {
	names := lo.Uniq[string]([]string{"Samuel", "John", "Samuel", "Falco"})
	fmt.Fprintf(w, "Hello, %v!", names)
}
