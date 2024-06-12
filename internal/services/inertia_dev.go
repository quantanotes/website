//go:build dev

package services

import (
	"encoding/json"
	"html/template"
	"os"

	"github.com/petaki/inertia-go"
)

var (
	Inertia  = inertia.New("localhost:3000", "static/index.html", "")
	manifest = make(map[string]string)
)

func init() {
	Inertia.ShareFunc("asset", asset)
}

func asset(path string) template.URL {
	loadManifest()
	// hacky but it is what it is
	if path == "www/app.js" {
		path = "www/dev.js"
	}
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

	file, err := os.Open("static/build/manifest.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	if err := dec.Decode(&manifest); err != nil {
		return nil, err
	}

	return manifest, nil
}
