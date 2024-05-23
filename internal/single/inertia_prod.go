//go:build !dev

package single

import (
	"encoding/json"
	"html/template"
	"quanta/internal/globals"
	"quanta/static"

	"github.com/petaki/inertia-go"
)

var (
	Inertia  = inertia.NewWithFS("localhost:3000", "index.html", "", static.FS)
	manifest = make(map[string]string)
)

func init() {
	if globals.SSRURL == "" {
		Inertia.EnableSsrWithDefault()
	} else {
		Inertia.EnableSsr(globals.SSRURL)
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
