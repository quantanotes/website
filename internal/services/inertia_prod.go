//go:build !dev

package services

import (
	"encoding/json"
	"html/template"
	"os"
	"quanta/static"

	"github.com/petaki/inertia-go"
)

var (
	Inertia  = inertia.NewWithFS("localhost:3000", "index.html", "", static.FS)
	manifest = make(map[string]string)
)

func init() {
	ssrURL := os.Getenv("SSR_URL")
	if ssrURL == "" {
		Inertia.EnableSsrWithDefault()
	} else {
		Inertia.EnableSsr(ssrURL)
	}
	Inertia.ShareFunc("asset", asset)
	if err := loadManifest(); err != nil {
		panic(err)
	}
}

func asset(path string) template.URL {
	return template.URL(manifest[path])
}

func loadManifest() error {
	manifestFile, err := loadManifestFile()
	if err != nil {
		return err
	}

	for k, v := range manifestFile {
		manifest[k] = v.File
	}

	return nil
}

func loadManifestFile() (map[string]struct{ File string }, error) {
	var manifest map[string]struct{ File string }
	if err := json.Unmarshal([]byte(static.Manifest), &manifest); err != nil {
		return nil, err
	}
	return manifest, nil
}
