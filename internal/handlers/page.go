package handlers

import (
	"net/http"
	"quanta/internal/globals"
	"quanta/internal/middleware"
	"quanta/internal/services"

	"github.com/go-chi/chi/v5"
)

type loader = func(http.ResponseWriter, *http.Request) (map[string]any, error)

func page(page string, loader loader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderPage(w, r, page, loader)
	}
}

func restrictedPage(public string, private string, loader loader) http.HandlerFunc {
	h := func(w http.ResponseWriter, r *http.Request) {
		if auth, ok := r.Context().Value(globals.AuthorisedContextKey).(bool); ok && auth {
			renderPage(w, r, private, loader)
		} else {
			renderPage(w, r, public, loader)
		}
	}
	return middleware.Restricted(http.HandlerFunc(h)).ServeHTTP
}

func getParam(param string) loader {
	return func(w http.ResponseWriter, r *http.Request) (map[string]any, error) {
		props := make(map[string]any, 1)
		props[param] = chi.URLParam(r, param)
		return props, nil
	}
}

func getQuery(queries ...string) loader {
	return func(w http.ResponseWriter, r *http.Request) (map[string]any, error) {
		props := make(map[string]any, len(queries))
		for _, q := range queries {
			props[q] = r.URL.Query().Get(q)
		}
		return props, nil
	}
}

func renderPage(w http.ResponseWriter, r *http.Request, page string, loader loader) {
	var (
		props map[string]any
		err   error
	)

	if loader != nil {
		if props, err = loader(w, r); err != nil {
			writeError(w, err)
			return
		}
	}

	if err := services.Inertia.Render(w, r, page, props); err != nil {
		writeError(w, err)
	}
}
