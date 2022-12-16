package handlers

import (
	"html/template"
	"net/http"
)

type Homepage struct {
	tpl *template.Template
}

func NewHomepage(tpl *template.Template) Homepage {
	return Homepage{
		tpl: tpl,
	}
}

func (h Homepage) HomepageHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.tpl.ExecuteTemplate(w, "index.html", nil); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
