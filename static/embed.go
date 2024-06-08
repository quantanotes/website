//go:build !dev

package static

import "embed"

var (
	//go:embed index.html images/*
	FS embed.FS

	//go:embed build/assets/*
	Assets embed.FS

	//go:embed build/manifest.json
	Manifest string
)
