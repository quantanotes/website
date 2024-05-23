//go:build !dev

package handlers

import (
	"io/fs"
	"net/http"
	_static "quanta/static"
)

func files(strip string, fs fs.FS) http.Handler {
	return http.StripPrefix(strip, http.FileServerFS(fs))
}

func static() http.Handler {
	return files("/static", _static.FS)
}

func assets() http.Handler {
	assets, _ := fs.Sub(_static.Assets, "build/assets")
	return files("/assets", assets)
}
