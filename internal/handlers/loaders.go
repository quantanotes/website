package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

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
