//go:build dev

package handlers

import (
	"net/http"
)

func files(strip string, dir http.Dir) http.Handler {
	return http.StripPrefix(strip, http.FileServer(dir))
}

func static() http.Handler {
	return files("/static", http.Dir("./static"))
}

func assets() http.Handler {
	return files("/assets", http.Dir("./static/build/assets"))
}
