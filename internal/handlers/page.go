package handlers

import (
	"net/http"
	"quanta/internal/middleware"
	"quanta/internal/single"
)

type loader = func(http.ResponseWriter, *http.Request) (map[string]any, error)

func page(page string, loader loader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, r, page, loader)
	}
}

func restrictedPage(public string, private string, loader loader) http.HandlerFunc {
	f := func(w http.ResponseWriter, r *http.Request) {
		if auth, ok := r.Context().Value("authorised").(bool); ok && auth {
			renderPage(w, r, private, loader)
		} else {
			renderPage(w, r, public, loader)
		}
	}
	h := middleware.Restricted(http.HandlerFunc(f))
	return func(w http.ResponseWriter, r *http.Request) {
		h.ServeHTTP(w, r)
	}
}

func renderPage(w http.ResponseWriter, r *http.Request, page string, loader loader) {
	var (
		props map[string]any
		err   error
	)

	if loader != nil {
		if props, err = loader(w, r); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := single.Inertia.Render(w, r, page, props); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
